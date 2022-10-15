// author: wsfuyibing <websearch@163.com>
// date: 2022-10-15

package kafka

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/fuyibing/log/v3/base"
	"github.com/fuyibing/log/v3/formatters"
)

type (
	handler struct {
		ch          chan *base.Line
		ct          *time.Ticker
		concurrency int32

		index    uint64
		mapper   map[uint64]*base.Line
		mu       *sync.RWMutex
		producer *kafka.Producer

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
func (o *handler) Log(line *base.Line, _ error) {
	// 1. 服务退出.
	if o.ch == nil {
		if o.engine != nil {
			o.engine.Log(line, fmt.Errorf("error on kafka adapter: stopped"))
		}
		return
	}

	// 2. 捕获异常.
	defer func() {
		if r := recover(); r != nil && o.engine != nil {
			o.engine.Log(line, fmt.Errorf("panic on kafka adapter: %v", r))
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
		var err error

		// 1. 监听结束.
		defer func() {
			// 关闭通道.
			if o.ch != nil {
				close(o.ch)
				o.ch = nil
			}

			// 关闭定时.
			if o.ct != nil {
				o.ct.Stop()
				o.ct = nil
			}

			// 关闭连接.
			if o.producer != nil {
				o.producer.Close()
			}

			// 捕获异常.
			if r := recover(); r != nil && o.engine != nil {
				o.engine.Log(nil, fmt.Errorf("panic on kafka adapter: %v", r))
			}
		}()

		// 2. 创建连接.
		if o.producer, err = kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers":         strings.Join(Config.Brokers, ","),
			"socket.timeout.ms":         10,
			"message.timeout.ms":        10,
			"go.delivery.report.fields": "key,value,headers"},
		); err != nil {
			if o.engine != nil {
				o.engine.Log(nil, fmt.Errorf("error on kafka adapter: %s", err.Error()))
			}
			return
		}
		if err = ctx.Err(); err != nil {
			if o.engine != nil {
				o.engine.Log(nil, fmt.Errorf("error on kafka adapter: %s", err.Error()))
			}
			return
		}

		// 3. 监听信息.
		o.ch = make(chan *base.Line)
		o.ct = time.NewTicker(time.Duration(Config.BatchFrequency) * time.Millisecond)
		for {
			select {
			case evt := <-o.producer.Events():
				go o.onEvent(evt)
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
	o.mapper = make(map[uint64]*base.Line)
	o.mu = new(sync.RWMutex)
	return o
}

// 发布事件.
func (o *handler) onEvent(e kafka.Event) {
}

// 取出日志.
func (o *handler) onPop() {
	var (
		concurrency = atomic.AddInt32(&o.concurrency, 1)
		count       = 0
		err         error
		list        []*base.Line
	)

	// 1. 并发限流.
	if concurrency > Config.Concurrency {
		atomic.AddInt32(&o.concurrency, -1)
		return
	}

	// 2. 监听结束.
	defer func() {
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

	// 4. 写入数据.
	for _, line := range list {
		err = o.send(line)

		// 写入成功.
		if err == nil {
			line.Release()
			continue
		}

		// 写入失败.
		if o.engine != nil {
			o.engine.Log(line, err)
		}
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

// 发布日志.
func (o *handler) send(line *base.Line) error {
	return o.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &Config.Topic, Partition: kafka.PartitionAny},
		Value:          []byte(formatters.Formatter.AsJson(line, nil)),
	}, nil)
}
