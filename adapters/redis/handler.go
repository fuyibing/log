// author: wsfuyibing <websearch@163.com>
// date: 2022-10-13

package redis

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/fuyibing/log/v3/adapters"
	"github.com/gomodule/redigo/redis"
)

type (
	// Handler
	// Redis适配器.
	Handler interface {
		// Interrupt
		// 注册拦截.
		//
		// 当提交日志出现错误时, 接过拦截器转发降级, 本适配降级时转发给文件存储将
		// 日志写入到本地文件中.
		Interrupt(fatal func(line *adapters.Line, err error)) Handler

		// Run
		// 提交日志.
		Run(line *adapters.Line, err error)

		// Start
		// 启动服务.
		Start()

		// Stop
		// 退出服务.
		Stop()
	}

	handler struct {
		cancel context.CancelFunc
		ctx    context.Context

		bufferIndex  uint64
		bufferMapper map[uint64]*adapters.Line            // 缓存列表
		ch           chan *adapters.Line                  // 消息通道
		concurrency  int32                                // 并发数
		connections  *redis.Pool                          // Redis连接池.
		interrupt    func(line *adapters.Line, err error) // 拦截回调
		listKey      string                               //
		mu           sync.RWMutex                         // 读写锁
	}
)

func New() Handler {
	return (&handler{}).init()
}

// Interrupt
// 注册拦截.
func (o *handler) Interrupt(interrupt func(line *adapters.Line, err error)) Handler {
	o.interrupt = interrupt
	return o
}

// Run
// 提交日志.
func (o *handler) Run(line *adapters.Line, _ error) {
	// 1. 拦截异常.
	defer func() {
		if v := recover(); v != nil {
			if o.interrupt != nil {
				o.interrupt(line, fmt.Errorf("panic on redis adapter: %v", v))
			}
		}
	}()

	// 2. 服务退出.
	if o.ch == nil {
		panic("service stopped")
		return
	}

	// 3. 提交数据.
	o.ch <- line
}

// Start
// 启动服务.
func (o *handler) Start() {
	go func() {
		// 1. 创建通道.
		o.mu.Lock()
		o.ctx, o.cancel = context.WithCancel(context.Background())
		o.ch = make(chan *adapters.Line)
		o.mu.Unlock()

		tick := time.NewTicker(time.Duration(Config.Ticker) * time.Millisecond)

		// 2. 监听结束.
		defer func() {
			recover()
			close(o.ch)
			tick.Stop()

			o.mu.Lock()
			o.ch = nil
			o.ctx = nil
			o.cancel = nil
			o.mu.Unlock()
		}()

		// 3. 监听信号.
		for {
			select {
			case line := <-o.ch:
				go o.push(line)
			case <-tick.C:
				go o.pop()
			case <-o.ctx.Done():
				return
			}
		}
	}()
}

// Stop
// 退出服务.
func (o *handler) Stop() {
	o.mu.RLock()
	defer o.mu.RUnlock()
	if o.ctx != nil && o.ctx.Err() == nil {
		o.cancel()
	}
}

// 构造实例.
func (o *handler) init() *handler {
	o.bufferMapper = make(map[uint64]*adapters.Line, 0)
	o.listKey = fmt.Sprintf("%s:%s", Config.KeyPrefix, Config.KeyList)
	o.mu = sync.RWMutex{}
	o.initConnection()
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
func (o *handler) pop() {
	// 1. 并发限流.
	if concurrency := atomic.AddInt32(&o.concurrency, 1); concurrency > Config.Concurrency {
		atomic.AddInt32(&o.concurrency, -1)
		return
	}

	// 2. 取出消息.
	//    若缓冲区无待提交的日志时退出.
	var (
		list  []*adapters.Line
		count int
	)
	defer func() {
		atomic.AddInt32(&o.concurrency, -1)

		// 继续提交.
		// 直到缓冲区处理完成.
		if count > 0 {
			o.pop()
		}
	}()
	if list, count = func() (l []*adapters.Line, n int) {
		o.mu.Lock()
		defer o.mu.Unlock()
		l = make([]*adapters.Line, 0)
		for k, v := range o.bufferMapper {
			delete(o.bufferMapper, k)
			l = append(l, v)
			n++
			if n >= Config.Limit {
				break
			}
		}
		return
	}(); count == 0 {
		return
	}

	// 3. 准备提交.
	var (
		conn = o.connections.Get()
		err  error
		key  string
		keys = []interface{}{o.listKey}
	)

	// 3.1 结束提交.
	defer func() {
		// 关闭连接.
		_ = conn.Close()

		// 运行异常.
		if v := recover(); v != nil && o.interrupt != nil {
			o.interrupt(nil, fmt.Errorf("panic on redis adapter: %v", v))
		}
	}()

	// 3.2 遍历日志.
	for _, line := range list {
		if key, err = o.sendLine(conn, line); err != nil {
			if o.interrupt != nil {
				o.interrupt(line, fmt.Errorf("error on redis adapter: %s", err.Error()))
			}
			continue
		}
		keys = append(keys, key)
	}

	// 3.2 写入列表.
	//     将 KEY 列表加入到 logger:list 中.
	if len(keys) > 1 {
		if err = o.sendList(conn, keys); err != nil && o.interrupt != nil {
			o.interrupt(nil, fmt.Errorf("error on redis adapter: %s", err.Error()))
		}
	}
}

// 塞入消息.
func (o *handler) push(line *adapters.Line) {
	var (
		i = atomic.AddUint64(&o.bufferIndex, 1)
		n = 0
	)

	// 1. 塞入缓冲.
	o.mu.Lock()
	o.bufferMapper[i] = line
	n = len(o.bufferMapper)
	o.mu.Unlock()

	// 2. 立即取出.
	if n >= Config.Limit {
		o.pop()
	}
}

// 写入数据.
//
// Redis: SET logger:172_16_0_1_3721_9981_1001 {"key":"value"} EX 3600
func (o *handler) sendLine(conn redis.Conn, line *adapters.Line) (key string, err error) {
	data := NewData(line)

	key = data.Key()
	if err = conn.Send("SET", key, data.String(), "EX", Config.KeyLifetime); err == nil {
		// println("line: ", line.GetId(), " -> ", line.GetAcquires())
		line.Release()
	}
	return
}

// 写入列表.
//
// Redis: RPUSH logger:list KEY1 KEY2 KEY3
func (o *handler) sendList(conn redis.Conn, keys []interface{}) error {
	return conn.Send("RPUSH", keys...)
}
