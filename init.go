// author: wsfuyibing <websearch@163.com>
// date: 2021-03-14

package log

import (
	"sync"
)

// Defaults.
const (
	DefaultAdapter     = AdapterTerm
	DefaultLevel       = LevelDebug
	DefaultTimeFormat  = "2006-01-02 15:04:05.999999"
	DefaultTraceId     = "X-B3-Traceid"
	DefaultSpanId      = "X-B3-Spanid"
	DefaultSpanVersion = "X-B3-Version"
	OpenTracingKey     = "OpenTracingKey"
)

var (
	Config *configuration
	Client *client
)

func init() {
	new(sync.Once).Do(func() {
		Config = &configuration{}
		Config.initialize()
		Client = &client{config: Config}
	})
}
