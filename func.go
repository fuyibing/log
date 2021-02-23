// author: wsfuyibing <websearch@163.com>
// date: 2021-02-23

package log

import (
	"context"
)

// 添加Debug日志.
func Debug(text string) {
	if Config.DebugOn() {
		Client.Debug(text)
	}
}

// 添加Debug日志, 支持格式化.
func Debugf(text string, args ...interface{}) {
	if Config.DebugOn() {
		Client.Debugf(text, args...)
	}
}

// 添加Debug日志, 支持格式化和请求链.
func Debugfc(ctx context.Context, text string, args ...interface{}) {
	if Config.DebugOn() {
		Client.Debugfc(ctx, text, args...)
	}
}

// 添加Info日志.
func Info(text string) {
	if Config.InfoOn() {
		Client.Info(text)
	}
}

// 添加Info日志, 支持格式化.
func Infof(text string, args ...interface{}) {
	if Config.InfoOn() {
		Client.Infof(text, args...)
	}
}

// 添加Info日志, 支持格式化和请求链.
func Infofc(ctx context.Context, text string, args ...interface{}) {
	if Config.InfoOn() {
		Client.Infofc(ctx, text, args...)
	}
}

// 添加Warn日志.
func Warn(text string) {
	if Config.WarnOn() {
		Client.Warn(text)
	}
}

// 添加Warn日志, 支持格式化.
func Warnf(text string, args ...interface{}) {
	if Config.WarnOn() {
		Client.Warnf(text, args...)
	}
}

// 添加Warn日志, 支持格式化和请求链.
func Warnfc(ctx context.Context, text string, args ...interface{}) {
	if Config.WarnOn() {
		Client.Warnfc(ctx, text, args...)
	}
}

// 添加Error日志.
func Error(text string) {
	if Config.ErrorOn() {
		Client.Error(text)
	}
}

// 添加Error日志, 支持格式化.
func Errorf(text string, args ...interface{}) {
	if Config.ErrorOn() {
		Client.Errorf(text, args...)
	}
}

// 添加Error日志, 支持格式化和请求链.
func Errorfc(ctx context.Context, text string, args ...interface{}) {
	if Config.ErrorOn() {
		Client.Errorfc(ctx, text, args...)
	}
}
