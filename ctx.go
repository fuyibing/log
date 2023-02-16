// author: wsfuyibing <websearch@163.com>
// date: 2023-02-16

package log

import (
	"context"
)

func NewContext() context.Context {
	return nil
}

func NewContextInfo(text string, args ...interface{}) context.Context {
	return nil
}

func NewChild(ctx context.Context) context.Context {
	return nil
}

func NewChildInfo(ctx context.Context, text string, args ...interface{}) context.Context {
	return nil
}
