// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package log

import (
	"context"
	"github.com/fuyibing/log/v8/conf"
)

// Level=Debug

func Debug(text string) {
	if Config.DebugOn() {
		Client.PushIntoBucket(nil, conf.Debug, nil, text)
	}
}

func Debugf(text string, args ...interface{}) {
	if Config.DebugOn() {
		Client.PushIntoBucket(nil, conf.Debug, nil, text, args...)
	}
}

func Debugfc(ctx context.Context, text string, args ...interface{}) {
	if Config.DebugOn() {
		Client.PushIntoBucket(ctx, conf.Debug, nil, text, args...)
	}
}

// Level=Info

func Info(text string) {
	if Config.InfoOn() {
		Client.PushIntoBucket(nil, conf.Info, nil, text)
	}
}

func Infof(text string, args ...interface{}) {
	if Config.InfoOn() {
		Client.PushIntoBucket(nil, conf.Info, nil, text, args...)
	}
}

func Infofc(ctx context.Context, text string, args ...interface{}) {
	if Config.InfoOn() {
		Client.PushIntoBucket(ctx, conf.Info, nil, text, args...)
	}
}

// Level=Warn

func Warn(text string) {
	if Config.WarnOn() {
		Client.PushIntoBucket(nil, conf.Warn, nil, text)
	}
}

func Warnf(text string, args ...interface{}) {
	if Config.WarnOn() {
		Client.PushIntoBucket(nil, conf.Warn, nil, text, args...)
	}
}

func Warnfc(ctx context.Context, text string, args ...interface{}) {
	if Config.WarnOn() {
		Client.PushIntoBucket(ctx, conf.Warn, nil, text, args...)
	}
}

// Level=Error

func Error(text string) {
	if Config.ErrorOn() {
		Client.PushIntoBucket(nil, conf.Error, nil, text)
	}
}

func Errorf(text string, args ...interface{}) {
	if Config.ErrorOn() {
		Client.PushIntoBucket(nil, conf.Error, nil, text, args...)
	}
}

func Errorfc(ctx context.Context, text string, args ...interface{}) {
	if Config.ErrorOn() {
		Client.PushIntoBucket(ctx, conf.Error, nil, text, args...)
	}
}

// Level=Fatal

func Fatal(text string) {
	if Config.FatalOn() {
		Client.PushIntoBucket(nil, conf.Fatal, nil, text)
	}
}

func Fatalf(text string, args ...interface{}) {
	if Config.FatalOn() {
		Client.PushIntoBucket(nil, conf.Fatal, nil, text, args...)
	}
}

func Fatalfc(ctx context.Context, text string, args ...interface{}) {
	if Config.FatalOn() {
		Client.PushIntoBucket(ctx, conf.Fatal, nil, text, args...)
	}
}

// Adapter: Fatal, Compatible with Panic

func Panic(text string) {
	if Config.FatalOn() {
		Client.PushIntoBucket(nil, conf.Fatal, nil, text)
	}
}

func Panicf(text string, args ...interface{}) {
	if Config.FatalOn() {
		Client.PushIntoBucket(nil, conf.Fatal, nil, text, args...)
	}
}

func Panicfc(ctx context.Context, text string, args ...interface{}) {
	if Config.FatalOn() {
		Client.PushIntoBucket(ctx, conf.Fatal, nil, text, args...)
	}
}
