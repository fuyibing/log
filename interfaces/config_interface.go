// author: wsfuyibing <websearch@163.com>
// date: 2021-02-22

package interfaces

// Config interface.
type ConfigInterface interface {
	AppAddr() string
	AppName() string
	DebugOn() bool
	ErrorOn() bool
	GetHandler() Handler
	GetLevel(level Level) string
	GetPid() int
	GetTimeFormat() string
	GetTrace() (traceId string, spanId string, spanVersion string)
	InfoOn() bool
	LoadYaml(string) error
	SetAdapter(string) ConfigInterface
	SetHandler(Handler) ConfigInterface
	SetLevel(string) ConfigInterface
	SetTimeFormat(string) ConfigInterface
	WarnOn() bool
}
