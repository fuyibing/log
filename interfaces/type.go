// author: wsfuyibing <websearch@163.com>
// date: 2021-02-22

package interfaces

type Adapter int

const (
	AdapterOff Adapter = iota
	AdapterTerm
	AdapterFile
	AdapterRedis
)

var AdapterTexts = map[Adapter]string{
	AdapterTerm:  "TERM",
	AdapterFile:  "FILE",
	AdapterRedis: "REDIS",
}

type Level int

const (
	LevelOff Level = iota
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
)

var LevelTexts = map[Level]string{
	LevelOff:   "OFF",
	LevelError: "ERROR",
	LevelWarn:  "WARN",
	LevelInfo:  "INFO",
	LevelDebug: "DEBUG",
}

const (
	DefaultAdapter     = AdapterTerm
	DefaultLevel       = LevelDebug
	DefaultTimeFormat  = "2006-01-02 15:04:05.999999"
	DefaultTraceId     = "X-B3-Traceid"
	DefaultSpanId      = "X-B3-Spanid"
	DefaultSpanVersion = "X-B3-Version"
	OpenTracingKey     = "OpenTracingKey"
)
