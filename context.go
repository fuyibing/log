// author: wsfuyibing <websearch@163.com>
// date: 2021-02-23

package log

import (
	"context"
	"fmt"

	"github.com/kataras/iris/v12"

	"github.com/fuyibing/log/v2/interfaces"
)

// 绑定Tracing.
// 在请求的入口进行绑定, 请求过程即可复用. 整个业务过程中使用
// 绑定后的Context, 可以保障同一个请求下的日志含相同的Span、
// Trace待标识.
func IrisBind(ctx iris.Context) {
	if ctx == nil {
		return
	}
	ctx.Values().Set(interfaces.OpenTracingKey, NewTracing().UseRequest(ctx.Request()))
}

// 创建上下文.
func NewContext() context.Context {
	return context.WithValue(context.TODO(), interfaces.OpenTracingKey, NewTracing().UseDefault())
}

// 子级上下文.
func ChildContext(ctx interface{}, text string, args ...interface{}) context.Context {
	// 1. return NewContext() if param ctx nil.
	if ctx == nil {
		return NewContext()
	}
	// 2. parse bound tracing.
	t := ParseTracing(ctx)
	if t == nil {
		return NewContext()
	}
	// 3. generate child tracing.
	prefix := ""
	if text != "" {
		prefix = fmt.Sprintf(text, args...) + " "
	}
	Infofc(ctx, prefix+"build child span.")
	ctn := context.WithValue(context.TODO(), interfaces.OpenTracingKey, NewTracing().Use(t.GetTraceId(), t.GenPreviewVersion()))
	Infofc(ctn, prefix+"child span builded.")
	return ctn
}
