// author: wsfuyibing <websearch@163.com>
// date: 2022-10-15

package redis

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/fuyibing/log/v3/base"
	"github.com/fuyibing/log/v3/formatters"
	"github.com/gomodule/redigo/redis"
)

type (
	handler struct {
		ch          chan *base.Line
		ct          *time.Ticker
		concurrency int32

		connections *redis.Pool
		index       uint64
		listKey     string
		mapper      map[uint64]*base.Line
		mu          *sync.RWMutex

		engine base.AdapterEngine
	}
)

// New
// 创建执行器.
func New() base.AdapterEngine {
	return (&handler{}).init()
}

// Log
// 发送日志.
func (o *handler) Log(line *base.Line) {
	if line == nil {
		return
	}

	// 1. 服务降级.
	if o.ch == nil {
		if o.engine != nil {
			o.engine.Log(line.WithError(fmt.Errorf("error on redis adapter: stopped or not start")))
		}
		return
	}

	// 2. 捕获异常.
	defer func() {
		if r := recover(); r != nil && o.engine != nil {
			o.engine.Log(line.WithError(fmt.Errorf("panic on redis adapter: %v", r)))
		}
	}()

	// 3. 发布日志.
	o.ch <- line
}

// Parent
// 绑定上级/降级.
func (o *handler) Parent(engine base.AdapterEngine) base.AdapterEngine {
	o.engine = engine
	return o
}

// Start
// 启动服务.
func (o *handler) Start(ctx context.Context) {
	// 1. 启动上级.
	if o.engine != nil {
		o.engine.Start(ctx)
	}

	// 2. 监听信号.
	go func() {
		o.ch = make(chan *base.Line)
		o.ct = time.NewTicker(time.Duration(Config.BatchFrequency) * time.Millisecond)

		// 结束监听.
		defer func() {
			// 关闭信息.
			close(o.ch)
			o.ch = nil

			// 关闭定时.
			o.ct.Stop()
			o.ct = nil

			// 捕获异常.
			if r := recover(); r != nil && o.engine != nil {
				o.engine.Log(base.NewInternalLine(fmt.Sprintf("panic on redis channel: %v", r)))
			}
		}()

		// 监听过程.
		for {
			select {
			case line := <-o.ch:
				go o.onPush(line)
			case <-o.ct.C:
				go o.onPop()
			case <-ctx.Done():
				return
			}
		}
	}()
}

// 构造实例.
func (o *handler) init() *handler {
	o.listKey = fmt.Sprintf("%s:%s", Config.KeyPrefix, Config.KeyList)
	o.mapper = make(map[uint64]*base.Line)
	o.mu = new(sync.RWMutex)
	o.initConnection()

	defaultPoolNode = strings.ReplaceAll(fmt.Sprintf("%s_%d_%d", base.LogHost, rand.Intn(9999), rand.Intn(9999)), ".", "_")
	return o
}

// 构建连接.
func (o *handler) initConnection() {
	// 1. 建连接池.
	o.connections = &redis.Pool{
		MaxActive: Config.MaxActive,
		MaxIdle:   Config.MaxIdle,
		Wait:      *Config.Wait,
	}

	// 2. 连接选项.
	o.connections.Dial = func() (redis.Conn, error) {
		opts := make([]redis.DialOption, 0)
		opts = append(opts, redis.DialPassword(Config.Password), redis.DialDatabase(Config.Database))

		if Config.Timeout > 0 {
			opts = append(opts, redis.DialConnectTimeout(time.Duration(Config.Timeout)*time.Second))
		}
		if Config.ReadTimeout > 0 {
			opts = append(opts, redis.DialReadTimeout(time.Duration(Config.ReadTimeout)*time.Second))
		}
		if Config.WriteTimeout > 0 {
			opts = append(opts, redis.DialWriteTimeout(time.Duration(Config.WriteTimeout)*time.Second))
		}

		return redis.Dial(Config.Network, Config.Address, opts...)
	}
}

// 取出消息.
func (o *handler) onPop() {
	var (
		conn        redis.Conn
		concurrency = atomic.AddInt32(&o.concurrency, 1)
		count       = 0
		err         error
		key         string
		keys        []interface{}
		list        []*base.Line
	)

	// 1. 并发限流.
	if concurrency > Config.Concurrency {
		atomic.AddInt32(&o.concurrency, -1)
		return
	}

	// 2. 监听结束.
	defer func() {
		// 关闭连接.
		if conn != nil {
			_ = conn.Close()
		}

		// 恢复统计.
		recover()
		atomic.AddInt32(&o.concurrency, -1)

		// 继续处理.
		if count > 0 {
			o.onPop()
		}
	}()

	// 3. 取出消息.
	if list, count = func() (l []*base.Line, n int) {
		o.mu.Lock()
		defer o.mu.Unlock()
		l = make([]*base.Line, 0)
		for k, v := range o.mapper {
			delete(o.mapper, k)
			l = append(l, v)
			if n++; n >= Config.BatchLimit {
				break
			}
		}
		return
	}(); count == 0 {
		return
	}

	// 4. 发布消息.
	conn = o.connections.Get()
	keys = []interface{}{o.listKey}

	// 4.1 遍历消费.
	for _, line := range list {
		key, err = o.sendLine(conn, line)

		// 发布成功.
		if err == nil {
			keys = append(keys, key)
			line.Release()
			continue
		}

		// 发布出错.
		if o.engine != nil {
			o.engine.Log(line.WithError(err))
		}
	}

	// 4.2 加入列表.
	if len(keys) > 1 {
		o.sendList(conn, keys)
	}
}

// 加入缓冲.
func (o *handler) onPush(line *base.Line) {
	var (
		i = atomic.AddUint64(&o.index, 1)
		n = 0
	)

	// 加入缓冲.
	o.mu.Lock()
	o.mapper[i] = line
	n = len(o.mapper)
	o.mu.Unlock()

	// 立即取出.
	if n >= Config.BatchLimit {
		o.onPop()
	}
}

// 发送日志.
func (o *handler) sendLine(conn redis.Conn, line *base.Line) (key string, err error) {
	key = fmt.Sprintf("%s:%s:%d", Config.KeyPrefix, defaultPoolNode, line.GetIndex())
	text := formatters.Formatter.AsJson(line)
	err = conn.Send("SET", key, text, "EX", Config.KeyLifetime)
	return
}

// 追加列表.
func (o *handler) sendList(conn redis.Conn, keys []interface{}) {
	_ = conn.Send("RPUSH", keys...)
}
