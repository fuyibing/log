// author: wsfuyibing <websearch@163.com>
// date: 2022-10-15

package log

import (
	"context"

	"github.com/fuyibing/log/v3/base"
)

// NewContext
// 创建链路上下文.
func NewContext() context.Context {
	return context.WithValue(context.Background(),
		base.TracingKey,
		base.NewTracing().WithStart(),
	)
}

// ChildContext
// 创建子链路上下文.
func ChildContext(ctx context.Context, text string, args ...interface{}) context.Context {
	if x, ok := ctx.Value(base.TracingKey).(*base.Tracing); ok {
		ctn := context.WithValue(ctx, base.TracingKey, base.NewTracing().WithParent(x))
		if text == "" {
			Client.Infofc(ctx, "child begin")
		} else {
			Client.Infofc(ctx, text, args...)
		}
		return ctn
	}
	return NewContext()
}
