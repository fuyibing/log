// author: wsfuyibing <websearch@163.com>
// date: 2023-02-20

package log

import (
	"context"
	"github.com/fuyibing/log/v8/conf"
)

type (
	// Map
	// configure struct fields into log.
	Map map[string]interface{}
)

func (p Map) Debugf(text string, args ...interface{}) {
	if Config.DebugOn() {
		Client.PushIntoBucket(nil, conf.Debug, p, text, args...)
	}
}

func (p Map) Debugfc(ctx context.Context, text string, args ...interface{}) {
	if Config.DebugOn() {
		Client.PushIntoBucket(ctx, conf.Debug, p, text, args...)
	}
}

func (p Map) Infof(text string, args ...interface{}) {
	if Config.InfoOn() {
		Client.PushIntoBucket(nil, conf.Info, p, text, args...)
	}
}

func (p Map) Infofc(ctx context.Context, text string, args ...interface{}) {
	if Config.InfoOn() {
		Client.PushIntoBucket(ctx, conf.Info, p, text, args...)
	}
}

func (p Map) Warnf(text string, args ...interface{}) {
	if Config.WarnOn() {
		Client.PushIntoBucket(nil, conf.Warn, p, text, args...)
	}
}

func (p Map) Warnfc(ctx context.Context, text string, args ...interface{}) {
	if Config.WarnOn() {
		Client.PushIntoBucket(ctx, conf.Warn, p, text, args...)
	}
}

func (p Map) Errorf(text string, args ...interface{}) {
	if Config.ErrorOn() {
		Client.PushIntoBucket(nil, conf.Error, p, text, args...)
	}
}

func (p Map) Errorfc(ctx context.Context, text string, args ...interface{}) {
	if Config.ErrorOn() {
		Client.PushIntoBucket(ctx, conf.Error, p, text, args...)
	}
}

func (p Map) Fatalf(text string, args ...interface{}) {
	if Config.FatalOn() {
		Client.PushIntoBucket(nil, conf.Fatal, p, text, args...)
	}
}

func (p Map) Fatalfc(ctx context.Context, text string, args ...interface{}) {
	if Config.FatalOn() {
		Client.PushIntoBucket(ctx, conf.Fatal, p, text, args...)
	}
}
