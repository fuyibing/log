// author: wsfuyibing <websearch@163.com>
// date: 2022-10-17

package log

import (
	"context"

	"github.com/fuyibing/log/v3/base"
)

// Client
// 日志客户端.
var Client *ClientManager

// ClientManager
// 客户端管理结构体.
type ClientManager struct {
	adapter   base.AdapterEngine // 适配器/引擎
	adapterOn bool               // 适配器/状态
	cancel    context.CancelFunc // 适配器/可取消回调
	ctx       context.Context    // 适配器/运行上下文
}

// Debug
// 调试日志.
func (o *ClientManager) Debug(text string) {
	if o.adapterOn && Config.debugOn {
		o.log(nil, false, base.Debug, text)
	}
}

// Debugf
// 调试日志.
func (o *ClientManager) Debugf(text string, args ...interface{}) {
	if o.adapterOn && Config.debugOn {
		o.log(nil, false, base.Debug, text, args...)
	}
}

// Debugfc
// 调试日志.
func (o *ClientManager) Debugfc(ctx context.Context, text string, args ...interface{}) {
	if o.adapterOn && Config.debugOn {
		o.log(ctx, false, base.Debug, text, args...)
	}
}

// Info
// 业务日志.
func (o *ClientManager) Info(text string) {
	if o.adapterOn && Config.infoOn {
		o.log(nil, false, base.Info, text)
	}
}

// Infof
// 业务日志.
func (o *ClientManager) Infof(text string, args ...interface{}) {
	if o.adapterOn && Config.infoOn {
		o.log(nil, false, base.Info, text, args...)
	}
}

// Infofc
// 业务日志.
func (o *ClientManager) Infofc(ctx context.Context, text string, args ...interface{}) {
	if o.adapterOn && Config.infoOn {
		o.log(ctx, false, base.Info, text, args...)
	}
}

// Warn
// 警告日志.
func (o *ClientManager) Warn(text string) {
	if o.adapterOn && Config.warnOn {
		o.log(nil, false, base.Warn, text)
	}
}

// Warnf
// 警告日志.
func (o *ClientManager) Warnf(text string, args ...interface{}) {
	if o.adapterOn && Config.warnOn {
		o.log(nil, false, base.Warn, text, args...)
	}
}

// Warnfc
// 警告日志.
func (o *ClientManager) Warnfc(ctx context.Context, text string, args ...interface{}) {
	if o.adapterOn && Config.warnOn {
		o.log(ctx, false, base.Warn, text, args...)
	}
}

// Error
// 错误日志.
func (o *ClientManager) Error(text string) {
	if o.adapterOn && Config.errorOn {
		o.log(nil, false, base.Error, text)
	}
}

// Errorf
// 错误日志.
func (o *ClientManager) Errorf(text string, args ...interface{}) {
	if o.adapterOn && Config.errorOn {
		o.log(nil, false, base.Error, text, args...)
	}
}

// Errorfc
// 错误日志.
func (o *ClientManager) Errorfc(ctx context.Context, text string, args ...interface{}) {
	if o.adapterOn && Config.errorOn {
		o.log(ctx, false, base.Error, text, args...)
	}
}

// Panic
// 异常日志.
func (o *ClientManager) Panic(text string) {
	if o.adapterOn && Config.errorOn {
		o.log(nil, true, base.Error, text)
	}
}

// Panicf
// 异常日志.
func (o *ClientManager) Panicf(text string, args ...interface{}) {
	if o.adapterOn && Config.errorOn {
		o.log(nil, true, base.Error, text, args...)
	}
}

// Panicfc
// 异常日志.
func (o *ClientManager) Panicfc(ctx context.Context, text string, args ...interface{}) {
	if o.adapterOn && Config.errorOn {
		o.log(ctx, true, base.Error, text, args...)
	}
}

// Start
// 启动服务.
func (o *ClientManager) Start() {
	if o.adapterOn && Config.Level != base.Off {
		o.ctx, o.cancel = context.WithCancel(context.Background())
		o.adapter.Start(o.ctx)
	}
}

// Stop
// 退出服务.
func (o *ClientManager) Stop() {
	if o.ctx != nil && o.ctx.Err() == nil {
		o.cancel()
		o.adapter.Wait()
	}
}

// 构造实例.
func (o *ClientManager) init() *ClientManager {
	return o
}

// 发送日志.
func (o *ClientManager) log(ctx context.Context, stack bool, level base.Level, text string, args ...interface{}) {
	// 1. 单行日志.
	line := base.NewLine(level, text, args)

	// 2. 绑上下文.
	if ctx != nil {
		line.WithContext(ctx)
	}

	// 3. 追加堆栈.
	if stack {
		line.WithStack()
	}

	// 4. 发送日志.
	o.adapter.Log(line)
}
