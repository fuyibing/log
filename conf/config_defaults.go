// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package conf

const (
	DefaultAdapter    = "term"
	DefaultLevel      = Info
	DefaultTimeFormat = "2006-01-02 15:04:05.999"

	DefaultBatchConcurrency = 10
	DefaultBatchFrequency   = 150
	DefaultBatchLimit       = 100

	DefaultParentSpanId = "X-B3-Parentspanid"
	DefaultSpanId       = "X-B3-Spanid"
	DefaultTraceId      = "X-B3-Traceid"
	DefaultTraceVersion = "X-B3-Version"

	Agent          = "LogV8"
	OpenTracingKey = "OpenTracingKey"
)

const (
	DefaultFileBasePath      = "./"
	DefaultFileExtName       = "log"
	DefaultFileFileName      = "2006-01-02"
	DefaultFileSeparatorPath = "2006-01"
)
