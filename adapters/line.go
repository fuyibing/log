// author: wsfuyibing <websearch@163.com>
// date: 2022-10-13

package adapters

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

type Line struct {
	id, acquires uint64
	index        uint64

	// 实时值

	Duration float64
	Content  string
	Level    Level
	Time     time.Time

	// 请求链.
	Trace         bool
	TraceId       string
	ParentSpanId  string
	SpanId        string
	SpanOffset    int32
	SpanPrefix    string
	Action        string
	RequestMethod string
	RequestUrl    string

	// 固定值

	Host       string
	Name       string
	Pid, Port  int
	Software   string
	TimeFormat string
	Version    string
}

// NewLine
// 创建实例.
func NewLine(level Level, text string, args ...interface{}) *Line {
	o := linePool.Get().(*Line)
	o.Content = fmt.Sprintf(text, args...)
	o.Level = level
	o.Time = time.Now()
	o.index = atomic.AddUint64(&lineIndex, 1)
	atomic.AddUint64(&o.acquires, 1)
	return o
}

func (o *Line) GetAcquires() uint64 {
	return o.acquires
}

func (o *Line) GetId() uint64 {
	return o.id
}

func (o *Line) GetIndex() uint64 {
	return o.index
}

func (o *Line) Release() {
	o.after()
	linePool.Put(o)
}

// String
// 格式化文本.
func (o *Line) String() (str string) {
	// 1. 基础信息.
	str = fmt.Sprintf("[%s][%s:%d][%s][%s][pid=%d]",
		o.Time.Format(o.TimeFormat),
		o.Host, o.Port,
		o.Name, o.Level.String(),
		o.Pid,
	)

	// 2. 链路信息.
	if o.Trace {
		str += fmt.Sprintf("[trace-id=%s][parent-span-id=%s][span-id=%s][span-version=%s.%d]",
			o.TraceId, o.ParentSpanId,
			o.SpanId, o.SpanPrefix, o.SpanOffset,
		)
		if o.Duration > 0 {
			str += fmt.Sprintf("[duration=%.06f]", o.Duration)
		}
		if o.RequestMethod != "" {
			str += fmt.Sprintf("[request-method=%s]", o.RequestMethod)
		}
		if o.RequestUrl != "" {
			str += fmt.Sprintf("[request-url=%s]", o.RequestUrl)
		}
	}

	// 3. 基础内容.
	str += fmt.Sprintf(" %s", o.Content)
	return
}

// WithContext
// 指定上下文.
func (o *Line) WithContext(ctx context.Context) *Line {
	if x, ok := ctx.Value(TracingKey).(*Tracing); ok {
		o.Trace = true
		o.TraceId = x.GetTraceId()
		o.ParentSpanId = x.GetSpanId()
		o.SpanId = x.GetSpanId()
		o.SpanOffset = x.GetOffsetIncr()
		o.SpanPrefix = x.GetPrefix()
	}
	return o
}

// WithStack
// 绑定堆栈信息.
func (o *Line) WithStack(stack bool) *Line {
	if stack {
		for i := 1; ; i++ {
			if _, f, l, g := runtime.Caller(i); g {
				o.Content += fmt.Sprintf("\n%s:%d", strings.TrimSpace(f), l)
				continue
			}
			break
		}
	}
	return o
}

func (o *Line) after() *Line {
	o.Content = ""

	// 重置请求链.

	o.Trace = false
	o.TraceId = ""
	o.ParentSpanId = ""
	o.SpanId = ""
	o.SpanOffset = 0
	o.SpanPrefix = ""

	return o
}

func (o *Line) init() *Line {
	o.id = atomic.AddUint64(&lineId, 1)
	o.Host = Host
	o.Name = Name
	o.Pid = Pid
	o.Port = Port
	o.Software = Software
	o.Version = Version
	o.TimeFormat = TimeFormat
	return o
}
