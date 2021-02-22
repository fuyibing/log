// author: wsfuyibing <websearch@163.com>
// date: 2021-02-22

package log

import (
	"context"
	"fmt"
	"os"

	"github.com/fuyibing/log/interfaces"
)

// 日志客户端.
type client struct{}

// 创建日志实例.
func newClient() interfaces.ClientInterface {
	o := &client{}
	return o
}

func (o *client) Debug(text string) {
	if Config.DebugOn() {
		o.log(nil, interfaces.LevelDebug, text)
	}
}

func (o *client) Debugf(text string, args ...interface{}) {
	if Config.DebugOn() {
		o.log(nil, interfaces.LevelDebug, text, args...)
	}
}

func (o *client) Debugfc(ctx context.Context, text string, args ...interface{}) {
	if Config.DebugOn() {
		o.log(ctx, interfaces.LevelDebug, text, args...)
	}
}

func (o *client) Info(text string) {
	if Config.InfoOn() {
		o.log(nil, interfaces.LevelInfo, text)
	}
}

func (o *client) Infof(text string, args ...interface{}) {
	if Config.InfoOn() {
		o.log(nil, interfaces.LevelInfo, text, args...)
	}
}

func (o *client) Infofc(ctx context.Context, text string, args ...interface{}) {
	if Config.InfoOn() {
		o.log(ctx, interfaces.LevelInfo, text, args...)
	}
}

func (o *client) Warn(text string) {
	if Config.WarnOn() {
		o.log(nil, interfaces.LevelWarn, text)
	}
}

func (o *client) Warnf(text string, args ...interface{}) {
	if Config.WarnOn() {
		o.log(nil, interfaces.LevelWarn, text, args...)
	}
}

func (o *client) Warnfc(ctx context.Context, text string, args ...interface{}) {
	if Config.WarnOn() {
		o.log(ctx, interfaces.LevelWarn, text, args...)
	}
}

func (o *client) Error(text string) {
	if Config.ErrorOn() {
		o.log(nil, interfaces.LevelError, text)
	}
}

func (o *client) Errorf(text string, args ...interface{}) {
	if Config.ErrorOn() {
		o.log(nil, interfaces.LevelError, text, args...)
	}
}

func (o *client) Errorfc(ctx context.Context, text string, args ...interface{}) {
	if Config.ErrorOn() {
		o.log(ctx, interfaces.LevelError, text, args...)
	}
}

func (o *client) log(ctx context.Context, level interfaces.Level, text string, args ...interface{}) {
	if handler := Config.GetHandler(); handler != nil {
		handler(NewLine(ctx, level, text, args))
	} else {
		_, _ = fmt.Fprintf(os.Stderr, "log adapter: log handler not defined\n")
	}
}
