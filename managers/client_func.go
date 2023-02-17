// author: wsfuyibing <websearch@163.com>
// date: 2023-02-17

package managers

import (
	"context"
	"github.com/fuyibing/log/v8/conf"
)

func (o *ClientManager) Debug(text string) {
	o.send(nil, conf.Debug, text)
}

func (o *ClientManager) Debugf(text string, args ...interface{}) {
	o.send(nil, conf.Debug, text, args...)
}

func (o *ClientManager) Debugfc(ctx context.Context, text string, args ...interface{}) {
	o.send(ctx, conf.Debug, text, args...)
}

func (o *ClientManager) Info(text string) {
	o.send(nil, conf.Info, text)
}

func (o *ClientManager) Infof(text string, args ...interface{}) {
	o.send(nil, conf.Info, text, args...)
}

func (o *ClientManager) Infofc(ctx context.Context, text string, args ...interface{}) {
	o.send(ctx, conf.Info, text, args...)
}

func (o *ClientManager) Warn(text string) {
	o.send(nil, conf.Warn, text)
}

func (o *ClientManager) Warnf(text string, args ...interface{}) {
	o.send(nil, conf.Warn, text, args...)
}

func (o *ClientManager) Warnfc(ctx context.Context, text string, args ...interface{}) {
	o.send(ctx, conf.Warn, text, args...)
}

func (o *ClientManager) Error(text string) {
	o.send(nil, conf.Error, text)
}

func (o *ClientManager) Errorf(text string, args ...interface{}) {
	o.send(nil, conf.Error, text, args...)
}

func (o *ClientManager) Errorfc(ctx context.Context, text string, args ...interface{}) {
	o.send(ctx, conf.Error, text, args...)
}

func (o *ClientManager) Fatal(text string) {
	o.send(nil, conf.Fatal, text)
}

func (o *ClientManager) Fatalf(text string, args ...interface{}) {
	o.send(nil, conf.Fatal, text, args...)
}

func (o *ClientManager) Fatalfc(ctx context.Context, text string, args ...interface{}) {
	o.send(ctx, conf.Fatal, text, args...)
}
