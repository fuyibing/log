// author: wsfuyibing <websearch@163.com>
// date: 2022-10-15

package kafka

import (
	"context"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/fuyibing/log/v3/base"
	"github.com/fuyibing/log/v3/formatters"
)

type (
	handler struct {
		ck       *kafka.ConfigMap
		engine   base.AdapterEngine
		producer *kafka.Producer

		cached map[string]int64
		mapper map[string]*base.Line
		mu     *sync.RWMutex
		node   string
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
	// 1. 降级处理.
	//    未启动 Kafka 生产者或正在启动中, 此过程若定义了降级适配器, 则
	//    转发到降级, 反之丢弃.
	if o.producer == nil {
		// 1.1 降级处理.
		if o.engine != nil {
			o.engine.Log(line.WithError(fmt.Errorf("error on kafka adapter: stopped or not start")))
			return
		}

		// 1.2 丢弃日志.
		line.Release()
		return
	}

	// 2. 准备发送.
	var (
		key   = fmt.Sprintf("%s_%d", o.node, line.GetIndex())
		value = formatters.Formatter.AsJson(line)
	)

	// 2.1 加入缓存.
	o.addCache(key, line)

	// 2.2 捕获异常.
	defer func() {
		if r := recover(); r == nil {
			return
		}
		o.onEngine(key)
		o.delCache(key)
	}()

	// 2.3 发送消息.
	o.producer.ProduceChannel() <- &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &Config.Topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          []byte(value),
	}
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
	// 2. 监听信息.
	go func() {
		o.node = strings.ReplaceAll(fmt.Sprintf("%s_%d_%d", base.LogHost, rand.Intn(9999), rand.Intn(9999)), ".", "_")

		for {
			if ctx.Err() != nil {
				break
			}
			o.listen(ctx)
		}
	}()
}

// /////////////////////////////////////////////////////////////
// 缓存操作
// /////////////////////////////////////////////////////////////

func (o *handler) addCache(key string, line *base.Line) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.cached[key] = line.Time.Unix()
	o.mapper[key] = line
}

func (o *handler) delCache(key string) {
	o.mu.Lock()
	defer o.mu.Unlock()
	delete(o.cached, key)
	delete(o.mapper, key)
}

func (o *handler) getCache(key string) *base.Line {
	o.mu.RLock()
	defer o.mu.RUnlock()
	if v, ok := o.mapper[key]; ok {
		return v
	}
	return nil
}

// /////////////////////////////////////////////////////////////
// 事件操作
// /////////////////////////////////////////////////////////////

func (o *handler) onClean() {
	var (
		kc = 0
		ks = make([]string, 0)
		ku = time.Now().Unix() - Config.Clean
	)

	// 遍历缓存.
	o.mu.RLock()
	for k, u := range o.cached {
		if u <= ku {
			kc++
			ks = append(ks, k)
		}
	}
	o.mu.RUnlock()

	// 遍历列表.
	for _, k := range ks {
		o.onEngine(k)
		o.delCache(k)
	}
}

func (o *handler) onEngine(key string) {
	if v := func() *base.Line {
		o.mu.RLock()
		defer o.mu.RUnlock()
		if v, ok := o.mapper[key]; ok {
			return v
		}
		return nil
	}(); v != nil {
		// 丢弃消息.
		if o.engine == nil {
			v.Release()
			return
		}

		// 降级处理.
		o.engine.Log(v)
	}
}

func (o *handler) onEvent(e kafka.Event) (err error) {
	switch t := e.(type) {
	case kafka.Error, *kafka.Error:
		// 1. 错误消息.
		//    Kafka 内部错误, 如网络断开等.
		err = fmt.Errorf("%v", t.String())

	case *kafka.Message:
		// 2. 发送成功.
		go func() {
			key := string(t.Key)
			if v := o.getCache(key); v != nil {
				v.Release()
			}
			o.delCache(key)
		}()

	default:
		// 3. 丢弃消息.
		{
			x := reflect.TypeOf(e)
			println("event type = ", x.Name(), " & text = ", t.String())
		}
	}
	return
}

// /////////////////////////////////////////////////////////////
// 实例操作
// /////////////////////////////////////////////////////////////

// 构造实例.
func (o *handler) init() *handler {
	o.cached = make(map[string]int64)
	o.mapper = make(map[string]*base.Line)
	o.mu = new(sync.RWMutex)

	o.ck = &kafka.ConfigMap{
		"bootstrap.servers":                  strings.Join(Config.Brokers, ","),
		"socket.connection.setup.timeout.ms": 1000 * Config.Timeout,
		"queue.buffering.max.messages":       Config.Batch,
		"go.batch.producer":                  true,
		"go.delivery.report.fields":          "key",
	}
	return o
}

// 监听信号.
func (o *handler) listen(ctx context.Context) {
	var (
		err error
		ct  *time.Ticker
	)

	// 1. 监听结束.
	defer func() {
		// 1.1 捕获异常.
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}

		// 1.2 关闭定时.
		if ct != nil {
			ct.Stop()
			ct = nil
		}

		// 1.3 降级转发.
		if err != nil && o.engine != nil {
			err = fmt.Errorf("error on kafka adapter: %v", err)
			o.engine.Log(base.NewInternalLine(err.Error()))
		}
	}()

	// 2. 建立连接.
	if o.producer, err = kafka.NewProducer(o.ck); err == nil {
		if err = ctx.Err(); err != nil {
			return
		}
	}

	// 3. 监听通道.
	ct = time.NewTicker(time.Duration(Config.Clean) * time.Second)

	for {
		select {
		case e := <-o.producer.Events():
			// 3.1 事件消息.
			if err = o.onEvent(e); err != nil {
				return
			}

		case <-ct.C:
			// 3.2 清理消息.
			go o.onClean()

		case <-ctx.Done():
			// 3.3 退出消息.
			return
		}
	}
}
