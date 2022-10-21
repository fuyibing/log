// author: wsfuyibing <websearch@163.com>
// date: 2022-10-17

package file

import (
	"context"
	"fmt"
	"github.com/fuyibing/log/v3/formatters"
	"github.com/fuyibing/util/v2/process"
	"sync"
	"sync/atomic"
	"time"

	"github.com/fuyibing/log/v3/base"
)

type Engine struct {
	engine       base.AdapterEngine
	engineCancel context.CancelFunc
	processor    process.Processor

	mu                  sync.RWMutex
	receives, completes uint64
	writers             map[string]Writer
}

func New() base.AdapterEngine {
	return (&Engine{}).init()
}

// /////////////////////////////////////////////////////////////
// Interface methods.
// /////////////////////////////////////////////////////////////

func (o *Engine) Log(line *base.Line) {
	atomic.AddUint64(&o.receives, 1)
	o.onSend(line)
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
				o.debugger("start engine: %v", err)
			}
			if o.engine != nil {
				o.engine.Log(base.NewInternalLine("error on file adapter: %v", err))
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

// 调试信息.
func (o *Engine) debugger(text string, args ...interface{}) {
	println(fmt.Sprintf("[adapter=file][%s] %s", time.Now().Format("15:04:05.999999"), fmt.Sprintf(text, args...)))
}

func (o *Engine) getWriter(line *base.Line) Writer {
	o.mu.Lock()
	defer o.mu.Unlock()

	key := fmt.Sprintf("%s_%s", line.Time.Format(Config.Folder), line.Time.Format(Config.Name))

	if w, ok := o.writers[key]; ok {
		return w
	}

	w := (&writer{}).init(line.Time)
	o.writers[key] = w
	return w
}

func (o *Engine) init() *Engine {
	o.mu = sync.RWMutex{}
	o.writers = make(map[string]Writer)
	o.processor = process.New("file-log-adapter").Panic(o.onPanic).
		After(o.onAfter, o.onAfterWait).
		Before(o.onBefore).
		Callback(o.onListenBefore, o.onListen, o.onListenAfter)
	return o
}

// /////////////////////////////////////////////////////////////
// Adapter events.
// /////////////////////////////////////////////////////////////

// 发送日志.
func (o *Engine) onSend(line *base.Line) {
	var err error

	// 写入完成.
	defer func() {
		atomic.AddUint64(&o.completes, 1)

		// 捕获异常.
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}

		// 写入成功.
		if err == nil || o.engine == nil {
			if Config.Debugger {
				o.debugger("engine wrote: %d", line.GetIndex())
			}
			line.Release()
			return
		}

		// 转发上级.
		o.engine.Log(line.RetryTimesReset().WithError(fmt.Errorf("error on file adapter: %v", err)))
	}()

	// 写入文件.
	err = o.getWriter(line).Write(formatters.Formatter.AsFile(line))
}

// /////////////////////////////////////////////////////////////
// Processor events.
// /////////////////////////////////////////////////////////////

// 执行器后置.
// 退出前确保执行中/待执行的任务处理完成.
func (o *Engine) onAfter(_ context.Context) (ignored bool) {
	for {
		if atomic.LoadUint64(&o.receives) == atomic.LoadUint64(&o.completes) {
			if Config.Debugger {
				o.debugger("engine stopped")
			}
			break
		}
		time.Sleep(time.Millisecond * 50)
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
		o.debugger("start engine")
	}
	return
}

// 执行器监听.
func (o *Engine) onListen(ctx context.Context) (ignored bool) {
	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}

// 执行器监听/后置.
func (o *Engine) onListenAfter(_ context.Context) (ignored bool) {
	return
}

// 执行器监听/前置.
func (o *Engine) onListenBefore(_ context.Context) (ignored bool) {
	return
}

// 执行器异常.
func (o *Engine) onPanic(_ context.Context, v interface{}) {
	if o.engine != nil {
		o.engine.Log(
			base.NewInternalLine(
				fmt.Sprintf("panic on file adapter: %v",
					v,
				),
			),
		)
	}
}
