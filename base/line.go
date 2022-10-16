// author: wsfuyibing <websearch@163.com>
// date: 2022-10-15

package base

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var (
	lineId, lineIndex uint64
	linePool          *sync.Pool
)

// Line
// 日志行.
type Line struct {
	id, acquires, index uint64

	// 日志值.
	Time    time.Time // 日期
	Level   Level     // 级别
	Content string    // 内容

	// 扩展值.

	Duration float64 // 执行时长

	// 链路值.

	Trace                                    bool
	SpanId, ParentSpanId, TraceId            string
	SpanOffset                               int32
	SpanPrefix                               string
	RequestAction, RequestMethod, RequestUrl string
}

// NewLine
// 创建实例.
func NewLine(level Level, text string, args []interface{}) *Line {
	o := linePool.Get().(*Line).before()
	o.Time = time.Now()
	o.Level = level
	o.Content = fmt.Sprintf(text, args...)
	return o
}

func NewInternalLine(text string, args ...interface{}) *Line {
	return NewLine(Error, text, args)
}

// GetIdentify
// 返回实例ID.
func (o *Line) GetIdentify() (id, acquires uint64) {
	return o.id, o.acquires
}

// GetIndex
// 返回日志索引量.
func (o *Line) GetIndex() uint64 {
	return o.index
}

// Release
// 释放入池.
func (o *Line) Release() {
	o.after()
	linePool.Put(o)
	// println("id = ", o.id, " & index = ", o.index, " & goroutines = ", runtime.NumGoroutine())
}

// WithContext
// 追加链路信息.
func (o *Line) WithContext(ctx context.Context) *Line {
	if x, ok := ctx.Value(TracingKey).(*Tracing); ok {
		o.Trace = true
		o.TraceId = x.GetTraceId()
		o.ParentSpanId = x.GetParentSpanId()
		o.SpanId = x.GetSpanId()
		o.SpanOffset = x.GetOffsetIncr()
		o.SpanPrefix = x.GetPrefix()
	}
	return o
}

// WithError
// 追加错误.
func (o *Line) WithError(err error) *Line {
	o.Content += fmt.Sprintf(" << interrupt: %s", err.Error())
	return o
}

// WithStack
// 追加堆栈信息.
func (o *Line) WithStack() *Line {
	for i := 3; ; i++ {
		if _, f, l, g := runtime.Caller(i); g {
			o.Content += fmt.Sprintf("\n%s:%d", strings.TrimSpace(f), l)
			continue
		}
		break
	}
	return o
}

// /////////////////////////////////////////////////////////////
// 池实例隐性操作
// /////////////////////////////////////////////////////////////

// 后置执行.
func (o *Line) after() *Line {
	return o
}

// 前置执行.
func (o *Line) before() *Line {
	atomic.AddUint64(&o.acquires, 1)
	o.index = atomic.AddUint64(&lineIndex, 1)
	return o
}

// 构造实例.
func (o *Line) init() *Line {
	o.id = atomic.AddUint64(&lineId, 1)
	return o
}
