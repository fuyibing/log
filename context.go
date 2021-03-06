// author: wsfuyibing <websearch@163.com>
// date: 2021-02-23

package log

import (
	"context"

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
