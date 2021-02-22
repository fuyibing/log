// author: wsfuyibing <websearch@163.com>
// date: 2021-02-22

package interfaces

import (
	"context"
)

// 客户端接口
type ClientInterface interface {
	Debug(text string)
	Debugf(text string, args ...interface{})
	Debugfc(ctx context.Context, text string, args ...interface{})
	Info(text string)
	Infof(text string, args ...interface{})
	Infofc(ctx context.Context, text string, args ...interface{})
	Warn(text string)
	Warnf(text string, args ...interface{})
	Warnfc(ctx context.Context, text string, args ...interface{})
	Error(text string)
	Errorf(text string, args ...interface{})
	Errorfc(ctx context.Context, text string, args ...interface{})
}
