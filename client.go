// author: wsfuyibing <websearch@163.com>
// date: 2022-10-15

package log

import (
	"context"

	"github.com/fuyibing/log/v3/base"
)

var Client ClientManager

type (
	// ClientManager
	// 客户端管理器.
	ClientManager interface {
		Debug(text string)
		Debugf(text string, args ...interface{})
		Debugfc(ctx context.Context, text string, args ...interface{})

		Info(text string)
		Infof(text string, args ...interface{})
		Infofc(ctx context.Context, text string, args ...interface{})

		Warn(text string)
		Warnf(text string, args ...interface{})
		Warnfc(ctx context.Context, text string, args ...interface{})

		Error(text string)
		Errorf(text string, args ...interface{})
		Errorfc(ctx context.Context, text string, args ...interface{})

		Panic(text string)
		Panicf(text string, args ...interface{})
		Panicfc(ctx context.Context, text string, args ...interface{})

		SetAdapter(engine base.AdapterEngine)
		Start()
		Stop()
	}

	// 客户端.
	client struct {
		cancel context.CancelFunc
		ctx    context.Context

		engine   base.AdapterEngine
		engineOn bool
	}
)

func (o *client) Debug(text string) {
	if o.engineOn && Config.debugOn {
		o.log(nil, false, base.Debug, text)
	}
}

func (o *client) Debugf(text string, args ...interface{}) {
	if o.engineOn && Config.debugOn {
		o.log(nil, false, base.Debug, text, args...)
	}
}

func (o *client) Debugfc(ctx context.Context, text string, args ...interface{}) {
	if o.engineOn && Config.debugOn {
		o.log(ctx, false, base.Debug, text, args...)
	}
}

func (o *client) Info(text string) {
	if o.engineOn && Config.infoOn {
		o.log(nil, false, base.Info, text)
	}
}

func (o *client) Infof(text string, args ...interface{}) {
	if o.engineOn && Config.infoOn {
		o.log(nil, false, base.Info, text, args...)
	}
}

func (o *client) Infofc(ctx context.Context, text string, args ...interface{}) {
	if o.engineOn && Config.infoOn {
		o.log(ctx, false, base.Info, text, args...)
	}
}

func (o *client) Warn(text string) {
	if o.engineOn && Config.warnOn {
		o.log(nil, false, base.Warn, text)
	}
}

func (o *client) Warnf(text string, args ...interface{}) {
	if o.engineOn && Config.warnOn {
		o.log(nil, false, base.Warn, text, args...)
	}
}

func (o *client) Warnfc(ctx context.Context, text string, args ...interface{}) {
	if o.engineOn && Config.warnOn {
		o.log(ctx, false, base.Warn, text, args...)
	}
}

func (o *client) Error(text string) {
	if o.engineOn && Config.errorOn {
		o.log(nil, false, base.Error, text)
	}
}

func (o *client) Errorf(text string, args ...interface{}) {
	if o.engineOn && Config.errorOn {
		o.log(nil, false, base.Error, text, args...)
	}
}

func (o *client) Errorfc(ctx context.Context, text string, args ...interface{}) {
	if o.engineOn && Config.errorOn {
		o.log(ctx, false, base.Error, text, args...)
	}
}

func (o *client) Panic(text string) {
	if o.engineOn && Config.errorOn {
		o.log(nil, true, base.Error, text)
	}
}

func (o *client) Panicf(text string, args ...interface{}) {
	if o.engineOn && Config.errorOn {
		o.log(nil, true, base.Error, text, args...)
	}
}

func (o *client) Panicfc(ctx context.Context, text string, args ...interface{}) {
	if o.engineOn && Config.errorOn {
		o.log(ctx, true, base.Error, text, args...)
	}
}

func (o *client) SetAdapter(engine base.AdapterEngine) {
	o.engine = engine
	o.engineOn = engine != nil
}

func (o *client) Start() {
	if o.engineOn {
		o.ctx, o.cancel = context.WithCancel(context.Background())
		o.engine.Start(o.ctx)
	}
}

func (o *client) Stop() {
	if o.ctx != nil && o.ctx.Err() == nil {
		o.cancel()
	}
}

// 构造实例.
func (o *client) init() *client {
	return o
}

// 发送日志.
func (o *client) log(ctx context.Context, stack bool, level base.Level, text string, args ...interface{}) {
	line := base.NewLine(level, text, args)

	if ctx != nil {
		line.WithContext(ctx)
	}

	if stack {
		line.WithStack()
	}

	o.engine.Log(line, nil)
}
