// author: wsfuyibing <websearch@163.com>
// date: 2021-02-23

package log

import (
	"context"
)

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
