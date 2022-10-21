// author: wsfuyibing <websearch@163.com>
// date: 2022-10-17

package log

import (
	"context"

	"github.com/fuyibing/log/v3/adapters/file"
	"github.com/fuyibing/log/v3/adapters/kafka"
	"github.com/fuyibing/log/v3/adapters/redis"
	"github.com/fuyibing/log/v3/adapters/term"
	"github.com/fuyibing/log/v3/base"
)

// Client
// 日志客户端.
var Client *client

// 客户端结构体.
type client struct {
	adapter   base.AdapterEngine // 适配器/引擎
	adapterOn bool               // 适配器/状态
	cancel    context.CancelFunc // 适配器/可取消回调
	ctx       context.Context    // 适配器/运行上下文
}

// Debug
// 调试日志.
func (o *client) Debug(text string) {
	if o.adapterOn && Config.debugOn {
		o.log(nil, false, base.Debug, text)
	}
}

// Debugf
// 调试日志.
func (o *client) Debugf(text string, args ...interface{}) {
	if o.adapterOn && Config.debugOn {
		o.log(nil, false, base.Debug, text, args...)
	}
}

// Debugfc
// 调试日志.
func (o *client) Debugfc(ctx context.Context, text string, args ...interface{}) {
	if o.adapterOn && Config.debugOn {
		o.log(ctx, false, base.Debug, text, args...)
	}
}

// Info
// 业务日志.
func (o *client) Info(text string) {
	if o.adapterOn && Config.infoOn {
		o.log(nil, false, base.Info, text)
	}
}

// Infof
// 业务日志.
func (o *client) Infof(text string, args ...interface{}) {
	if o.adapterOn && Config.infoOn {
		o.log(nil, false, base.Info, text, args...)
	}
}

// Infofc
// 业务日志.
func (o *client) Infofc(ctx context.Context, text string, args ...interface{}) {
	if o.adapterOn && Config.infoOn {
		o.log(ctx, false, base.Info, text, args...)
	}
}

// Warn
// 警告日志.
func (o *client) Warn(text string) {
	if o.adapterOn && Config.warnOn {
		o.log(nil, false, base.Warn, text)
	}
}

// Warnf
// 警告日志.
func (o *client) Warnf(text string, args ...interface{}) {
	if o.adapterOn && Config.warnOn {
		o.log(nil, false, base.Warn, text, args...)
	}
}

// Warnfc
// 警告日志.
func (o *client) Warnfc(ctx context.Context, text string, args ...interface{}) {
	if o.adapterOn && Config.warnOn {
		o.log(ctx, false, base.Warn, text, args...)
	}
}

// Error
// 错误日志.
func (o *client) Error(text string) {
	if o.adapterOn && Config.errorOn {
		o.log(nil, false, base.Error, text)
	}
}

// Errorf
// 错误日志.
func (o *client) Errorf(text string, args ...interface{}) {
	if o.adapterOn && Config.errorOn {
		o.log(nil, false, base.Error, text, args...)
	}
}

// Errorfc
// 错误日志.
func (o *client) Errorfc(ctx context.Context, text string, args ...interface{}) {
	if o.adapterOn && Config.errorOn {
		o.log(ctx, false, base.Error, text, args...)
	}
}

// Panic
// 异常日志.
func (o *client) Panic(text string) {
	if o.adapterOn && Config.errorOn {
		o.log(nil, true, base.Error, text)
	}
}

// Panicf
// 异常日志.
func (o *client) Panicf(text string, args ...interface{}) {
	if o.adapterOn && Config.errorOn {
		o.log(nil, true, base.Error, text, args...)
	}
}

// Panicfc
// 异常日志.
func (o *client) Panicfc(ctx context.Context, text string, args ...interface{}) {
	if o.adapterOn && Config.errorOn {
		o.log(ctx, true, base.Error, text, args...)
	}
}

// Start
// 启动服务.
func (o *client) Start() {
	if o.adapterOn && Config.Level != base.Off {
		o.ctx, o.cancel = context.WithCancel(context.Background())
		o.adapter.Start(o.ctx)
	}
}

// Stop
// 退出服务.
func (o *client) Stop() {
	if o.ctx != nil && o.ctx.Err() == nil {
		o.cancel()
		o.adapter.Wait()
	}
}

// 构造实例.
func (o *client) init() *client {
	o.initAdapters()
	return o
}

func (o *client) initAdapters() {
	var a base.AdapterEngine
	for _, c := range Config.Adapter {
		switch c.Adapter() {
		case base.Kafka:
			a = kafka.New().Parent(a)
		case base.Redis:
			a = redis.New().Parent(a)
		case base.File:
			a = file.New().Parent(a)
		case base.Term:
			a = term.New().Parent(a)
		}
	}
	if a != nil {
		o.adapter = a
		o.adapterOn = true
	}
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
	o.adapter.Log(line)
}
