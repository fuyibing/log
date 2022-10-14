// author: wsfuyibing <websearch@163.com>
// date: 2022-10-13

package log

import (
	"context"

	"github.com/fuyibing/log/v3/adapters"
)

// NewContext
// 创建链路上下文.
func NewContext() context.Context {
	return context.WithValue(
		context.Background(),
		adapters.TracingKey, adapters.NewTracing().WithStart(),
	)
}

// ChildContext
// 创建子链路上下文.
func ChildContext(ctx context.Context, text string, args ...interface{}) context.Context {
	if x, ok := ctx.Value(adapters.TracingKey).(*adapters.Tracing); ok {
		ctn := context.WithValue(ctx, adapters.TracingKey, adapters.NewTracing().WithParent(x))
		if text == "" {
			Client.Infofc(ctx, "begin child")
		} else {
			Client.Infofc(ctx, text, args...)
		}
		return ctn
	}
	return NewContext()
}
