// author: wsfuyibing <websearch@163.com>
// date: 2022-10-17

package log

import (
	"context"
	"github.com/fuyibing/log/v3/base"
)

// Debug
// 调试日志.
func Debug(text string) {
	if Client.adapterOn && Config.debugOn {
		Client.log(nil, false, base.Debug, text)
	}
}

// Debugf
// 调试日志.
func Debugf(text string, args ...interface{}) {
	if Client.adapterOn && Config.debugOn {
		Client.log(nil, false, base.Debug, text, args...)
	}
}

// Debugfc
// 调试日志.
func Debugfc(ctx context.Context, text string, args ...interface{}) {
	if Client.adapterOn && Config.debugOn {
		Client.log(ctx, false, base.Debug, text, args...)
	}
}

// Info
// 业务日志.
func Info(text string) {
	if Client.adapterOn && Config.infoOn {
		Client.log(nil, false, base.Info, text)
	}
}

// Infof
// 业务日志.
func Infof(text string, args ...interface{}) {
	if Client.adapterOn && Config.infoOn {
		Client.log(nil, false, base.Info, text, args...)
	}
}

// Infofc
// 业务日志.
func Infofc(ctx context.Context, text string, args ...interface{}) {
	if Client.adapterOn && Config.infoOn {
		Client.log(ctx, false, base.Info, text, args...)
	}
}

// Warn
// 警告日志.
func Warn(text string) {
	if Client.adapterOn && Config.warnOn {
		Client.log(nil, false, base.Warn, text)
	}
}

// Warnf
// 警告日志.
func Warnf(text string, args ...interface{}) {
	if Client.adapterOn && Config.warnOn {
		Client.log(nil, false, base.Warn, text, args...)
	}
}

// Warnfc
// 警告日志.
func Warnfc(ctx context.Context, text string, args ...interface{}) {
	if Client.adapterOn && Config.warnOn {
		Client.log(ctx, false, base.Warn, text, args...)
	}
}

// Error
// 错误日志.
func Error(text string) {
	if Client.adapterOn && Config.errorOn {
		Client.log(nil, false, base.Error, text)
	}
}

// Errorf
// 错误日志.
func Errorf(text string, args ...interface{}) {
	if Client.adapterOn && Config.errorOn {
		Client.log(nil, false, base.Error, text, args...)
	}
}

// Errorfc
// 错误日志.
func Errorfc(ctx context.Context, text string, args ...interface{}) {
	if Client.adapterOn && Config.errorOn {
		Client.log(ctx, false, base.Error, text, args...)
	}
}

// Panic
// 异常日志.
func Panic(text string) {
	if Client.adapterOn && Config.errorOn {
		Client.log(nil, true, base.Error, text)
	}
}

// Panicf
// 异常日志.
func Panicf(text string, args ...interface{}) {
	if Client.adapterOn && Config.errorOn {
		Client.log(nil, true, base.Error, text, args...)
	}
}

// Panicfc
// 异常日志.
func Panicfc(ctx context.Context, text string, args ...interface{}) {
	if Client.adapterOn && Config.errorOn {
		Client.log(ctx, true, base.Error, text, args...)
	}
}
