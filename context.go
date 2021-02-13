// author: wsfuyibing <websearch@163.com>
// date: 2021-02-10

package log

import (
	"context"

	"github.com/kataras/iris/v12"
)

// Return context.Context with default tracing.
func NewContext() context.Context {
	return context.WithValue(context.TODO(), OpenTracingContext, NewTracing().UseDefault())
}

// Return context.Context with iris.Context.
func NewContextWithIris(ctx iris.Context) context.Context {
	if v := ctx.Values().Get(OpenTracingContext); v != nil {
		if c, ok := v.(context.Context); ok {
			return c
		}
	}
	return NewContext()
}

// Reset context for iris.
func ResetIrisContext(ctx iris.Context) {
	if v := ctx.Values().Get(OpenTracingContext); v != nil {
		return
	}
	ctx.Values().Set(OpenTracingContext, NewTracing().UseRequest(ctx.Request()))
}
