// author: wsfuyibing <websearch@163.com>
// date: 2021-02-22

package interfaces

import (
	"time"
)

type LineInterface interface {
	ColorLevel() string
	Content() string
	Duration() float64
	Level() string
	ParentSpanId() string
	Pid() int
	RequestInfo() (method string, url string)
	ServiceAddr() string
	ServiceName() string
	SpanId() string
	SpanVersion() string
	Time() time.Time
	Timeline() string
	TraceId() string
	Tracing() bool
}