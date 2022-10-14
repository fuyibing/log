// author: wsfuyibing <websearch@163.com>
// date: 2022-10-13

package file

import (
	"sync"
	"time"

	"github.com/fuyibing/log/v3/adapters"
)

type (
	// Handler
	// 文件适配器.
	Handler interface {
		// Interrupt
		// 注册拦截.
		//
		// 当提交日志出现错误时, 接过拦截器转发降级, 本适配降级时转发给终端输出将
		// 日志打印到标准输出流.
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
		engines   map[string]Engine
		interrupt func(line *adapters.Line, err error)
		mu        sync.RWMutex
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
	en := o.getEngine(line.Time)
	err := en.Write(line.String())

	if err == nil {
		// println("line: ", line.GetId(), " -> ", line.GetAcquires())
		line.Release()
		return
	}

	if o.interrupt != nil {
		o.interrupt(line, err)
	}
}

// Start
// 启动服务.
func (o *handler) Start() {}

// Stop
// 退出服务.
func (o *handler) Stop() {}

// 读取引擎.
func (o *handler) getEngine(t time.Time) (en Engine) {
	o.mu.Lock()
	defer o.mu.Unlock()

	var (
		ok  bool
		key = t.Format(Config.Name)
	)

	if en, ok = o.engines[key]; ok {
		return
	}

	en = (&engine{}).init(t)
	o.engines[key] = en
	return

}

// 构造实例.
func (o *handler) init() *handler {
	o.engines = make(map[string]Engine)
	o.mu = sync.RWMutex{}
	return o
}
