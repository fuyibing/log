// author: wsfuyibing <websearch@163.com>
// date: 2022-10-13

package log

import (
	"context"

	"github.com/fuyibing/log/v3/adapters"
)

var Client *client

type (
	client struct {
		handler   adapters.Handler
		handlerOn bool
	}
)

// /////////////////////////////////////////////////////////////
// Protects
// /////////////////////////////////////////////////////////////

func (o *client) Debug(text string) {
	if o.handlerOn && Config.debugOn {
		o.send(nil, false, adapters.Debug, text)
	}
}

func (o *client) Debugf(text string, args ...interface{}) {
	if o.handlerOn && Config.debugOn {
		o.send(nil, false, adapters.Debug, text, args...)
	}
}

func (o *client) Debugfc(ctx context.Context, text string, args ...interface{}) {
	if o.handlerOn && Config.debugOn {
		o.send(ctx, false, adapters.Debug, text, args...)
	}
}

func (o *client) Info(text string) {
	if o.handlerOn && Config.infoOn {
		o.send(nil, false, adapters.Info, text)
	}
}

func (o *client) Infof(text string, args ...interface{}) {
	if o.handlerOn && Config.infoOn {
		o.send(nil, false, adapters.Info, text, args...)
	}
}

func (o *client) Infofc(ctx context.Context, text string, args ...interface{}) {
	if o.handlerOn && Config.infoOn {
		o.send(ctx, false, adapters.Info, text, args...)
	}
}

func (o *client) Warn(text string) {
	if o.handlerOn && Config.warnOn {
		o.send(nil, false, adapters.Warn, text)
	}
}

func (o *client) Warnf(text string, args ...interface{}) {
	if o.handlerOn && Config.warnOn {
		o.send(nil, false, adapters.Warn, text, args...)
	}
}

func (o *client) Warnfc(ctx context.Context, text string, args ...interface{}) {
	if o.handlerOn && Config.warnOn {
		o.send(ctx, false, adapters.Warn, text, args...)
	}
}

func (o *client) Error(text string) {
	if o.handlerOn && Config.errorOn {
		o.send(nil, false, adapters.Error, text)
	}
}

func (o *client) Errorf(text string, args ...interface{}) {
	if o.handlerOn && Config.errorOn {
		o.send(nil, false, adapters.Error, text, args...)
	}
}

func (o *client) Errorfc(ctx context.Context, text string, args ...interface{}) {
	if o.handlerOn && Config.errorOn {
		o.send(ctx, false, adapters.Error, text, args...)
	}
}

func (o *client) Panic(text string) {
	if o.handlerOn && Config.errorOn {
		o.send(nil, true, adapters.Error, text)
	}
}

func (o *client) Panicf(text string, args ...interface{}) {
	if o.handlerOn && Config.errorOn {
		o.send(nil, true, adapters.Error, text, args...)
	}
}

func (o *client) Panicfc(ctx context.Context, text string, args ...interface{}) {
	if o.handlerOn && Config.errorOn {
		o.send(ctx, true, adapters.Error, text, args...)
	}
}

// /////////////////////////////////////////////////////////////
// Protects
// /////////////////////////////////////////////////////////////

// 构造实例.
func (o *client) init() *client {
	return o
}

// 写入日志.
func (o *client) send(ctx context.Context, stack bool, level adapters.Level, text string, args ...interface{}) {
	line := adapters.NewLine(level, text, args...)
	line.WithContext(ctx).WithStack(stack)
	o.handler(line, nil)
}

// 设置回调.
func (o *client) setHandler(handler adapters.Handler, fn ...func()) {
	for _, f := range fn {
		f()
	}
	o.handler = handler
	o.handlerOn = handler != nil
}
