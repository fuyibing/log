// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package log

import (
	"context"
)

// Level=Debug

func Debug(text string) {
	if Config.DebugOn() {
		Client.Debug(text)
	}
}

func Debugf(text string, args ...interface{}) {
	if Config.DebugOn() {
		Client.Debugf(text, args...)
	}
}

func Debugfc(ctx context.Context, text string, args ...interface{}) {
	if Config.DebugOn() {
		Client.Debugfc(ctx, text, args...)
	}
}

// Level=Info

func Info(text string) {
	if Config.InfoOn() {
		Client.Info(text)
	}
}

func Infof(text string, args ...interface{}) {
	if Config.InfoOn() {
		Client.Infof(text, args...)
	}
}

func Infofc(ctx context.Context, text string, args ...interface{}) {
	if Config.InfoOn() {
		Client.Infofc(ctx, text, args...)
	}
}

// Level=Warn

func Warn(text string) {
	if Config.WarnOn() {
		Client.Warn(text)
	}
}

func Warnf(text string, args ...interface{}) {
	if Config.WarnOn() {
		Client.Warnf(text, args...)
	}
}

func Warnfc(ctx context.Context, text string, args ...interface{}) {
	if Config.WarnOn() {
		Client.Warnfc(ctx, text, args...)
	}
}

// Level=Error

func Error(text string) {
	if Config.ErrorOn() {
		Client.Error(text)
	}
}

func Errorf(text string, args ...interface{}) {
	if Config.ErrorOn() {
		Client.Errorf(text, args...)
	}
}

func Errorfc(ctx context.Context, text string, args ...interface{}) {
	if Config.ErrorOn() {
		Client.Errorfc(ctx, text, args...)
	}
}

// Level=Fatal

func Fatal(text string) {
	if Config.FatalOn() {
		Client.Fatal(text)
	}
}

func Fatalf(text string, args ...interface{}) {
	if Config.FatalOn() {
		Client.Fatalf(text, args...)
	}
}

func Fatalfc(ctx context.Context, text string, args ...interface{}) {
	if Config.FatalOn() {
		Client.Fatalfc(ctx, text, args...)
	}
}

// Adapter: Fatal, Compatible with Panic

func Panic(text string) {
	if Config.FatalOn() {
		Client.Fatal(text)
	}
}

func Panicf(text string, args ...interface{}) {
	if Config.FatalOn() {
		Client.Fatalf(text, args...)
	}
}

func Panicfc(ctx context.Context, text string, args ...interface{}) {
	if Config.FatalOn() {
		Client.Fatalfc(ctx, text, args...)
	}
}
