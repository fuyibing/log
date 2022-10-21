// author: wsfuyibing <websearch@163.com>
// date: 2022-10-17

package redis

import (
	"context"
	"fmt"
	"github.com/fuyibing/log/v3/formatters"
	"github.com/fuyibing/util/v2/process"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
	"sync"
	"sync/atomic"
	"time"

	"github.com/fuyibing/log/v3/base"
)

var (
	errStopped = fmt.Errorf("redis engine stopped")
)

type Engine struct {
	engine       base.AdapterEngine // 上级引擎
	engineCancel context.CancelFunc // 上级引擎可取消回调

	lineBuffers     map[uint64]*base.Line // 暂存区
	lineChan        chan *base.Line       // 暂存通道
	lineConcurrency int32                 // 并发限流
	lineDelay       []*base.Line          // 延时区
	lineTicker      *time.Ticker          // 定时器
	mu              sync.RWMutex          // 读写锁
	node            string                // 节点名
	redisList       string                // Redis List

	processor process.Processor // 执行器
	producer  *redis.Pool       // 生产者
}

func New() base.AdapterEngine {
	return (&Engine{}).init()
}

// /////////////////////////////////////////////////////////////
// Interface methods.
// /////////////////////////////////////////////////////////////

// Log
// 日志入口.
//
// 1. 发送暂存信号
// 2. 延时发送
func (o *Engine) Log(line *base.Line) {
	go func() {
		delay := false

		defer func() {
			if r := recover(); r != nil {
				delay = true
			}
			if delay {
				o.addDelay(line)
			}
		}()

		if o.lineChan != nil {
			o.lineChan <- line
			return
		}

		delay = true
	}()
}

// Parent
// 绑定上级引擎.
func (o *Engine) Parent(engine base.AdapterEngine) base.AdapterEngine {
	o.engine = engine
	return o
}

// Start
// 启动引擎.
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
				o.engine.Log(base.NewInternalLine("error on redis adapter: %v", err))
			}
		}
	}()
}

// Wait
// 等待完成.
func (o *Engine) Wait() {
	if o.processor.Stopped() {
		return
	}
	time.Sleep(time.Millisecond * 10)
	o.Wait()
}

// /////////////////////////////////////////////////////////////
// Adapter.
// /////////////////////////////////////////////////////////////

func (o *Engine) debugger(text string, args ...interface{}) {
	println(fmt.Sprintf("[adapter=redis][%s] %s", time.Now().Format("15:04:05.999999"), fmt.Sprintf(text, args...)))
}

func (o *Engine) init() *Engine {
	o.lineDelay = make([]*base.Line, 0)
	o.lineBuffers = make(map[uint64]*base.Line)
	o.mu = sync.RWMutex{}
	o.node = uuid.New().String()
	o.redisList = fmt.Sprintf("%s:%s", Config.KeyPrefix, Config.KeyList)
	o.processor = process.New("redis-log-adapter").Panic(o.onPanic).
		After(o.onAfter, o.onAfterWait).
		Before(o.onBefore).
		Callback(o.onListen)

	o.initConnection()
	return o
}

func (o *Engine) initConnection() {
	// 1. 建连接池.
	o.producer = &redis.Pool{
		MaxIdle:         Config.MaxIdle,
		MaxActive:       Config.MaxActive,
		MaxConnLifetime: time.Duration(Config.Lifetime) * time.Second,
		Wait:            *Config.Wait,
	}

	// 2. 连接选项.
	o.producer.Dial = func() (redis.Conn, error) {
		// 2.1 账号与库.
		options := []redis.DialOption{
			redis.DialPassword(Config.Password),
			redis.DialDatabase(Config.Database),
		}
		// 2.2 连接超时.
		if Config.Timeout > 0 {
			options = append(options,
				redis.DialConnectTimeout(time.Duration(Config.Timeout)*time.Second),
			)
		}
		// 2.3 读取超时.
		if Config.ReadTimeout > 0 {
			options = append(options,
				redis.DialReadTimeout(time.Duration(Config.ReadTimeout)*time.Second),
			)
		}
		// 2.3 写入超时.
		if Config.WriteTimeout > 0 {
			options = append(options,
				redis.DialWriteTimeout(time.Duration(Config.WriteTimeout)*time.Second),
			)
		}
		// 2.4 连接过程.
		return redis.Dial(Config.Network, Config.Address, options...)
	}
}

// /////////////////////////////////////////////////////////////
// Adapter actions.
// /////////////////////////////////////////////////////////////

// 加入延时.
func (o *Engine) addDelay(line *base.Line) {
	o.mu.Lock()
	o.lineDelay = append(o.lineDelay, line.RetryTimesIncrement())
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

// 发送日志.
func (o *Engine) doSend(lines ...*base.Line) {
	var (
		conn  redis.Conn
		count = 0
		err   error
		key   string
		keys  []interface{}
	)

	// 1. 发送结束.
	defer func() {
		if r := recover(); r != nil && Config.Debugger {
			o.debugger("redis panic: %v", r)
		}
		if conn != nil {
			_ = conn.Close()
		}
	}()

	// 2. 准备发送.
	conn = o.producer.Get()
	keys = []interface{}{o.redisList}

	// 3. 遍历日志.
	for _, line := range lines {
		key = fmt.Sprintf("%s:%s:%d", Config.KeyPrefix, o.node, line.GetIndex())

		// 3.1 写入过程.
		if err = conn.Send("SET", key, formatters.Formatter.AsJson(line), "EX", Config.KeyLifetime); err == nil {
			// 3.1.1 写入成功.
			count++
			keys = append(keys, key)
			o.doSendSucceed(line)
		} else {
			// 3.1.2 写入出错.
			if Config.Debugger {
				o.debugger("redis error on set: %v", err)
			}
			o.doSendError(line, err)
		}
	}

	// 2.2 推送列表.
	if count > 0 {
		if err = conn.Send("RPUSH", keys...); err != nil && Config.Debugger {
			o.debugger("redis error on push: %v", err)
		}
	}
}

// 发送出错.
func (o *Engine) doSendError(line *base.Line, err error) {
	if line.RetryTimes() < Config.RetryTimes {
		o.addDelay(line)
	} else {
		o.doParent(line, err)
	}
}

// 发送成功.
func (o *Engine) doSendSucceed(line *base.Line) {
	line.Release()
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
	o.doSend(list...)
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

	// 1. 转发上级/延时.
	for _, line := range o.lineDelay {
		o.doParent(line, errStopped)
	}

	// 2. 转发上级/暂存.
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

// 执行器异常.
func (o *Engine) onPanic(_ context.Context, v interface{}) {
	if o.engine != nil {
		o.engine.Log(
			base.NewInternalLine(
				fmt.Sprintf("panic on redis adapter: %v",
					v,
				),
			),
		)
	}
}
