// author: wsfuyibing <websearch@163.com>
// date: 2021-02-22

package log

import (
	"fmt"
	"os"

	"github.com/fuyibing/log/v2/interfaces"
)

// 日志客户端.
type client struct{}

// 创建日志实例.
func newClient() interfaces.ClientInterface {
	o := &client{}
	return o
}

// 添加Debug日志.
func (o *client) Debug(text string) {
	if Config.DebugOn() {
		o.log(nil, interfaces.LevelDebug, text)
	}
}

// 添加Debug日志, 支持格式化.
func (o *client) Debugf(text string, args ...interface{}) {
	if Config.DebugOn() {
		o.log(nil, interfaces.LevelDebug, text, args...)
	}
}

// 添加Debug日志, 支持格式化和请求链.
func (o *client) Debugfc(ctx interface{}, text string, args ...interface{}) {
	if Config.DebugOn() {
		o.log(ctx, interfaces.LevelDebug, text, args...)
	}
}

// 添加Info日志.
func (o *client) Info(text string) {
	if Config.InfoOn() {
		o.log(nil, interfaces.LevelInfo, text)
	}
}

// 添加Info日志, 支持格式化.
func (o *client) Infof(text string, args ...interface{}) {
	if Config.InfoOn() {
		o.log(nil, interfaces.LevelInfo, text, args...)
	}
}

// 添加Info日志, 支持格式化和请求链.
func (o *client) Infofc(ctx interface{}, text string, args ...interface{}) {
	if Config.InfoOn() {
		o.log(ctx, interfaces.LevelInfo, text, args...)
	}
}

// 添加Warn日志.
func (o *client) Warn(text string) {
	if Config.WarnOn() {
		o.log(nil, interfaces.LevelWarn, text)
	}
}

// 添加Warn日志, 支持格式化.
func (o *client) Warnf(text string, args ...interface{}) {
	if Config.WarnOn() {
		o.log(nil, interfaces.LevelWarn, text, args...)
	}
}

// 添加Warn日志, 支持格式化和请求链.
func (o *client) Warnfc(ctx interface{}, text string, args ...interface{}) {
	if Config.WarnOn() {
		o.log(ctx, interfaces.LevelWarn, text, args...)
	}
}

// 添加Error日志.
func (o *client) Error(text string) {
	if Config.ErrorOn() {
		o.log(nil, interfaces.LevelError, text)
	}
}

// 添加Error日志, 支持格式化.
func (o *client) Errorf(text string, args ...interface{}) {
	if Config.ErrorOn() {
		o.log(nil, interfaces.LevelError, text, args...)
	}
}

// 添加Error日志, 支持格式化和请求链.
func (o *client) Errorfc(ctx interface{}, text string, args ...interface{}) {
	if Config.ErrorOn() {
		o.log(ctx, interfaces.LevelError, text, args...)
	}
}

// 日志处理逻辑.
func (o *client) log(ctx interface{}, level interfaces.Level, text string, args ...interface{}) {
	if handler := Config.GetHandler(); handler != nil {
		handler(NewLine(ctx, level, text, args))
	} else {
		_, _ = fmt.Fprintf(os.Stderr, "log adapter: log handler not defined\n")
	}
}
