// author: wsfuyibing <websearch@163.com>
// date: 2021-02-22

package interfaces

// Config interface.
type ConfigInterface interface {
	DebugOn() bool
	InfoOn() bool
	WarnOn() bool
	ErrorOn() bool
	GetHandler() Handler
	GetTimeFormat() string
	GetTrace() (traceId string, spanId string, spanVersion string)
	LoadYaml(string) error
	SetAdapter(string) ConfigInterface
	SetHandler(Handler) ConfigInterface
	SetTimeFormat(string) ConfigInterface
	SetLevel(string) ConfigInterface
}
