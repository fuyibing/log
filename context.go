// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package log

import (
	"context"
	"github.com/fuyibing/log/v8/trace"
	"net/http"
)

func NewChild(ctx context.Context) context.Context {
	return trace.NewChild(ctx)
}

func NewChildInfo(ctx context.Context, text string, args ...interface{}) context.Context {
	c := trace.NewChild(ctx)
	Infofc(c, text, args...)
	return c
}

func NewContext() context.Context {
	return trace.NewContext()
}

func NewContextInfo(text string, args ...interface{}) context.Context {
	ctx := trace.NewContext()
	Infofc(ctx, text, args...)
	return ctx
}

func NewRequest(request *http.Request) context.Context {
	return trace.NewRequest(request)
}

func NewRequestInfo(request *http.Request, text string, args ...interface{}) context.Context {
	c := trace.NewRequest(request)
	Infofc(c, text, args...)
	return c
}
