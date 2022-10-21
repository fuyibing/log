// author: wsfuyibing <websearch@163.com>
// date: 2022-10-17

package kafka

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/fuyibing/log/v3/base"
	"github.com/fuyibing/log/v3/formatters"
	"github.com/fuyibing/util/v2/process"
	"sync"
	"sync/atomic"
	"time"
)

type Engine struct {
	engine       base.AdapterEngine
	engineCancel context.CancelFunc

	lineBuffers     map[uint64]*base.Line
	lineChan        chan *base.Line
	lineConcurrency int32
	lineDelay       []*base.Line
	lineTicker      *time.Ticker
	mu              sync.RWMutex

	producer  sarama.SyncProducer
	processor process.Processor
}

var (
	errStarting = fmt.Errorf("kafka engine starting")
	errStopped  = fmt.Errorf("kafka engine stopped")
)

func New() base.AdapterEngine {
	return (&Engine{}).init()
}

// /////////////////////////////////////////////////////////////
// Interface methods.
// /////////////////////////////////////////////////////////////

func (o *Engine) Log(line *base.Line) {
	go func() {
		var err error

		// 1. 发送结果.
		defer func() {
			// 1.1 捕获异常.
			if r := recover(); r != nil {
				err = fmt.Errorf("%v", r)
			}
			if err == nil {
				return
			}

			// 1.2 延时发送.
			if line.RetryTimes() < Config.RetryTimes {
				o.addDelay(line.RetryTimesIncrement())
				return
			}

			// 1.3 转发上级.
			//     发送(多次)失败后, 转发给上游引擎(文件)处理.
			if o.engine != nil {
				o.engine.Log(line.WithError(err))
				return
			}

			// 1.4 释放消息.
			line.Release()
		}()

		// 2. 发送完成.
		if o.lineChan != nil {
			o.lineChan <- line
			return
		}

		// 3. 延时发送.
		err = errStarting
	}()
}

func (o *Engine) Parent(engine base.AdapterEngine) base.AdapterEngine {
	o.engine = engine
	return o
}

func (o *Engine) Start(ctx context.Context) {
	// 1. 启动上级.
	if o.engine != nil {
		var c context.Context
		c, o.engineCancel = context.WithCancel(context.Background())
		o.engine.Start(c)
	}

	// 2. 启动本级.
	go func() {
		if err := o.processor.Start(ctx); err != nil {
			if Config.Debugger {
				o.debugger("engine start: %v", err)
			}
			if o.engine != nil {
				o.engine.Log(base.NewInternalLine("error on kafka adapter: %v", err))
			}
		}
	}()
}

func (o *Engine) Wait() {
	if o.processor.Stopped() {
		return
	}
	time.Sleep(time.Millisecond * 50)
	o.Wait()
}

// /////////////////////////////////////////////////////////////
// Adapter operations.
// /////////////////////////////////////////////////////////////

func (o *Engine) debugger(text string, args ...interface{}) {
	println(fmt.Sprintf("[adapter=kafka][%s] %s", time.Now().Format("15:04:05.999999"), fmt.Sprintf(text, args...)))
}

func (o *Engine) init() *Engine {
	o.lineDelay = make([]*base.Line, 0)
	o.lineBuffers = make(map[uint64]*base.Line)
	o.mu = sync.RWMutex{}

	o.processor = process.New("kafka-log-adapter").Panic(o.onPanic).
		After(o.onAfter, o.onAfterWait).
		Before(o.onBefore).
		Callback(o.onListenBefore, o.onListen, o.onListenAfter)
	return o
}

// /////////////////////////////////////////////////////////////
// Adapter actions.
// /////////////////////////////////////////////////////////////

// 加入延时.
func (o *Engine) addDelay(line *base.Line) {
	o.mu.Lock()
	o.lineDelay = append(o.lineDelay, line)
	o.mu.Unlock()
}

// 转发上级.
// 若未定义上级引擎, 则丢弃日志.
func (o *Engine) doParent(line *base.Line, err error) {
	// 上级未定义.
	if o.engine == nil {
		line.Release()
		return
	}
	// 转发给上级.
	o.engine.Log(line.RetryTimesReset().WithError(err))
}

// 发送消息.
func (o *Engine) doPublish(count int, lines ...*base.Line) {
	var (
		err  error
		list = make([]*sarama.ProducerMessage, 0)
	)

	// 1. 发送结束.
	defer func() {
		// 1.1 捕获异常.
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}

		// 1.2 记录结果.
		if err != nil {
			if Config.Debugger {
				o.debugger("kafka publish: count=%d, error: %v", count, err)
			}
		} else {
			if Config.Debugger {
				o.debugger("kafka published: %d messages", count)
			}
		}
	}()

	// 2. 消息列表.
	for _, line := range lines {
		list = append(list, &sarama.ProducerMessage{
			Topic: Config.KafkaTopic,
			Value: sarama.StringEncoder(formatters.Formatter.AsJson(line)),
		})
		line.Release()
	}
	err = o.producer.SendMessages(list)
}

// /////////////////////////////////////////////////////////////
// Adapter events.
// /////////////////////////////////////////////////////////////

// 取出日志.
func (o *Engine) onPop() {
	// 1. 计算并发.
	concurrency := atomic.AddInt32(&o.lineConcurrency, 1)

	// 2. 并发限流.
	if concurrency > Config.Concurrency {
		atomic.AddInt32(&o.lineConcurrency, -1)
		return
	}

	// 3. 处理过程.
	var (
		count = 0
		list  []*base.Line
	)

	// 3.1 处理完成.
	defer func() {
		atomic.AddInt32(&o.lineConcurrency, -1)

		// 继续处理.
		if count > 0 {
			o.onPop()
		}
	}()

	// 3.2 取出消息.
	if list, count = func() (l []*base.Line, n int) {
		o.mu.Lock()
		defer o.mu.Unlock()
		l = make([]*base.Line, 0)
		for i, line := range o.lineBuffers {
			delete(o.lineBuffers, i)
			l = append(l, line)
			if n++; n >= Config.Batch {
				break
			}
		}
		return
	}(); count == 0 {
		return
	}

	// 3.3 发送消息.
	o.doPublish(count, list...)
}

// 塞入缓存.
func (o *Engine) onPush(line *base.Line) {
	o.mu.Lock()
	o.lineBuffers[line.GetIndex()] = line
	o.mu.Unlock()
	o.onPop()
}

// 处理重试.
func (o *Engine) onRetry() {
	// 1. 延时区为空.
	if count := func() int {
		o.mu.RLock()
		defer o.mu.RUnlock()
		return len(o.lineDelay)
	}(); count == 0 {
		return
	}

	// 2. 重新发送.
	for _, line := range func() (l []*base.Line) {
		o.mu.Lock()
		defer o.mu.Unlock()
		l = o.lineDelay
		o.lineDelay = make([]*base.Line, 0)
		return
	}() {
		o.Log(line)
	}
}

// /////////////////////////////////////////////////////////////
// Processor events.
// /////////////////////////////////////////////////////////////

// 执行器后置.
// 退出前确保执行中/待执行的任务处理完成.
func (o *Engine) onAfter(_ context.Context) (ignored bool) {
	o.mu.Lock()
	defer o.mu.Unlock()

	// 1. 转发延时.
	for _, line := range o.lineDelay {
		o.doParent(line, errStopped)
	}

	// 2. 转发暂存.
	for _, line := range o.lineBuffers {
		o.doParent(line, errStopped)
	}

	// 3. 清空暂存.
	o.lineBuffers = make(map[uint64]*base.Line)
	o.lineDelay = make([]*base.Line, 0)

	// 4. 等待完成.
	for {
		if atomic.LoadInt32(&o.lineConcurrency) == 0 {
			break
		}
		time.Sleep(time.Millisecond * 50)
	}

	// 5. 完成退出.
	if Config.Debugger {
		o.debugger("engine stopped")
	}
	return
}

// 执行器后置/阻塞.
// 本级完全退出前, 通知上级退出并等待全部处理完成.
func (o *Engine) onAfterWait(_ context.Context) (ignored bool) {
	if o.engineCancel != nil {
		o.engineCancel()
		o.engine.Wait()
	}
	return
}

// 启动前置.
func (o *Engine) onBefore(_ context.Context) (ignored bool) {
	if Config.Debugger {
		o.debugger("engine start")
	}
	return
}

// 执行器监听.
func (o *Engine) onListen(ctx context.Context) (ignored bool) {
	if Config.Debugger {
		o.debugger("listener begin")
	}

	// 1. 准备监听.
	o.lineChan = make(chan *base.Line)
	o.lineTicker = time.NewTicker(time.Duration(Config.RetrySeconds) * time.Second)

	// 2. 完成监听.
	defer func() {
		// 2.1 关闭通道.
		close(o.lineChan)
		o.lineChan = nil

		// 2.2 关闭定时.
		o.lineTicker.Stop()
		o.lineTicker = nil

		if Config.Debugger {
			o.debugger("listener ended")
		}
	}()

	// 3. 监听过程.
	for {
		select {
		case line := <-o.lineChan:
			// 3.1 收到消息.
			go o.onPush(line)

		case <-o.lineTicker.C:
			// 3.2 定时重试.
			go o.onRetry()

		case <-ctx.Done():
			// 3.3 退出监听.
			return
		}
	}
}

// 执行器监听/后置.
func (o *Engine) onListenAfter(_ context.Context) (ignored bool) {
	// 1. Kafka 关闭.
	if err := o.producer.Close(); err != nil {
		if Config.Debugger {
			o.debugger("kafka close: %v", err)
		}
	} else if Config.Debugger {
		o.debugger("kafka closed")
	}

	// 2. Kafka 重置.
	o.producer = nil
	return
}

// 执行器监听/前置.
func (o *Engine) onListenBefore(_ context.Context) (ignored bool) {
	var (
		cfg *sarama.Config
		cli sarama.Client
		err error
	)

	// 1. 构建配置.
	cfg = sarama.NewConfig()
	cfg.ClientID = base.LogName

	// 1.1 网络超时.
	cfg.Net.MaxOpenRequests = Config.KafkaConnectRequests
	cfg.Net.KeepAlive = time.Duration(Config.KafkaConnectKeepAlive) * time.Second
	cfg.Net.DialTimeout = time.Duration(Config.KafkaConnectTimeout) * time.Second
	cfg.Net.ReadTimeout = time.Duration(Config.KafkaReadTimeout) * time.Second
	cfg.Net.WriteTimeout = time.Duration(Config.KafkaWriteTimeout) * time.Second

	// 1.2 生产者
	cfg.Producer.MaxMessageBytes = Config.KafkaMaxMessageSize
	cfg.Producer.RequiredAcks = sarama.NoResponse

	// 1.3 生产者返回.
	cfg.Producer.Return.Errors = true
	cfg.Producer.Return.Successes = true

	// 1.4 生产者刷盘.
	cfg.Producer.Flush.Messages = Config.KafkaFlushMessages
	cfg.Producer.Flush.Frequency = time.Duration(Config.KafkaFlushSeconds) * time.Second

	// 2. 建立连接.
	if cli, err = sarama.NewClient(Config.KafkaBrokers, cfg); err == nil {
		o.producer, err = sarama.NewSyncProducerFromClient(cli)
	}
	if err != nil {
		if err != nil {
			o.debugger("kafka connect: host=%v, error=%v", Config.KafkaBrokers, err)
		}
		o.processor.Restart()
	} else if Config.Debugger {
		o.debugger("kafka connected: %v", Config.KafkaBrokers)
	}
	return err != nil
}

// 执行器异常.
func (o *Engine) onPanic(_ context.Context, v interface{}) {
	if o.engine != nil {
		o.engine.Log(
			base.NewInternalLine(
				fmt.Sprintf("panic on kafka adapter: %v",
					v,
				),
			),
		)
	}
}
