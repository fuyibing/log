// author: wsfuyibing <websearch@163.com>
// date: 2022-10-13

package kafka

import (
	"context"
	"sync"

	"github.com/fuyibing/log/v3/adapters"
)

type (
	// Handler
	// Kafka适配器.
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
		mu     sync.RWMutex

		interrupt func(line *adapters.Line, err error)
	}
)

func New() Handler {
	return (&handler{}).init()
}

func (o *handler) Interrupt(interrupt func(line *adapters.Line, err error)) Handler {
	o.interrupt = interrupt
	return o
}

func (o *handler) Run(line *adapters.Line, err error) {
	if o.interrupt != nil {
		o.interrupt(line, err)
	}
}

// Start
// 启动服务.
func (o *handler) Start() {}

// Stop
// 退出服务.
func (o *handler) Stop() {
	o.mu.RLock()
	defer o.mu.RUnlock()
	if o.ctx != nil && o.ctx.Err() == nil {
		o.cancel()
	}
}

func (o *handler) init() *handler {
	return o
}
