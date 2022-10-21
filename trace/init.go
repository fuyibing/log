// author: wsfuyibing <websearch@163.com>
// date: 2022-10-17

package trace

const (
	defaultSpanPrefix = "0"
)

var (
	TracingKey          = "FlogOpenTracingKey"
	TracingTraceId      = "X-B3-Traceid"
	TracingParentSpanId = "X-B3-Parentspanid"
	TracingSpanId       = "X-B3-Spanid"
	TracingSpanVersion  = "X-B3-Version"
)
