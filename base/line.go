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

	"github.com/fuyibing/log/v3/trace"
)

var (
	lineId, lineIndex uint64
	linePool          *sync.Pool
)

// Line
// 日志行.
type Line struct {
	id, acquires, index uint64
	retryTimes          int32

	// 日志值.
	Content  string    // 内容
	Duration float64   // 执行时长
	Level    Level     // 级别
	Time     time.Time // 日期

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

// NewInternalLine
// 内部调用.
func NewInternalLine(text string, args ...interface{}) *Line {
	return NewLine(Warn, text, args)
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

// RetryTimes
// 读取重试次数.
func (o *Line) RetryTimes() int32 {
	return atomic.LoadInt32(&o.retryTimes)
}

// RetryTimesIncrement
// 重试次数递加.
func (o *Line) RetryTimesIncrement() *Line {
	atomic.AddInt32(&o.retryTimes, 1)
	return o
}

// RetryTimesReset
// 重置重试次数.
func (o *Line) RetryTimesReset() *Line {
	atomic.StoreInt32(&o.retryTimes, 0)
	return o
}

// Release
// 释放入池.
func (o *Line) Release() {
	o.after()
	linePool.Put(o)
}

// WithContext
// 追加链路信息.
func (o *Line) WithContext(ctx context.Context) *Line {
	if ctx != nil {
		if x, ok := ctx.Value(trace.OpenTracingKey).(*trace.Tracing); ok {
			o.Trace = true
			o.TraceId = x.TraceId
			o.ParentSpanId = x.ParentSpanId
			o.SpanId = x.SpanId
			o.SpanPrefix = x.SpanPrefix
			o.SpanOffset = x.Increment() - 1

			o.RequestMethod = x.RequestMethod
			o.RequestUrl = x.RequestUrl
		}
	}
	return o
}

// WithError
// 追加错误.
func (o *Line) WithError(err error) *Line {
	if err != nil {
		o.Content += fmt.Sprintf(" << interrupt: %s", err.Error())
	}
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
	o.Content = ""
	o.Duration = 0
	o.Trace = false
	o.SpanId = ""
	o.ParentSpanId = ""
	o.TraceId = ""
	o.SpanOffset = 0
	o.SpanPrefix = ""
	o.RequestAction = ""
	o.RequestMethod = ""
	o.RequestUrl = ""
	return o
}

// 前置执行.
func (o *Line) before() *Line {
	atomic.AddUint64(&o.acquires, 1)
	o.index = atomic.AddUint64(&lineIndex, 1)
	atomic.StoreInt32(&o.retryTimes, 0)
	return o
}

// 构造实例.
func (o *Line) init() *Line {
	o.id = atomic.AddUint64(&lineId, 1)
	return o
}
