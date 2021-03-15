// author: wsfuyibing <websearch@163.com>
// date: 2021-03-14

package log

import (
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v3"
)

// 配置参数.
type configuration struct {
	Adapter     Adapter `yaml:"-"`
	AdapterName string  `yaml:"adapter"`
	Level       Level   `yaml:"-"`
	LevelName   string  `yaml:"level"`
	TimeFormat  string  `yaml:"time"`
	SpanId      string  `yaml:"span-currId"`
	SpanVersion string  `yaml:"span-version"`
	TraceId     string  `yaml:"trace-currId"`
	handler     Handler `yaml:"-"`
	debugOn     bool
	infoOn      bool
	warnOn      bool
	errorOn     bool
}

func (o *configuration) DebugOn() bool { return o.debugOn }
func (o *configuration) InfoOn() bool  { return o.infoOn }
func (o *configuration) WarnOn() bool  { return o.warnOn }
func (o *configuration) ErrorOn() bool { return o.errorOn }
func (o *configuration) GetTrace() (traceId, spanId, spanVersion string) {
	return o.TraceId, o.SpanId, o.SpanVersion
}

// 设置回调.
func (o *configuration) SetHandler(handler Handler) *configuration {
	o.handler = handler
	return o
}

// 初始化.
func (o *configuration) initialize() {
	o.Adapter = DefaultAdapter
	o.Level = DefaultLevel
	// 从YAML读取配置.
	for _, file := range []string{"./tmp/log.yaml", "../tmp/log.yaml", "./config/log.yaml", "../config/log.yaml"} {
		body, e1 := ioutil.ReadFile(file)
		if e1 != nil {
			continue
		}
		if e2 := yaml.Unmarshal(body, o); e2 != nil {
			continue
		}
		break
	}
	// Trace identify.
	if o.SpanId == "" {
		o.SpanId = DefaultSpanId
	}
	if o.SpanVersion == "" {
		o.SpanVersion = DefaultSpanVersion
	}
	if o.TraceId == "" {
		o.TraceId = DefaultTraceId
	}
	// Use default time format.
	if o.TimeFormat == "" {
		o.TimeFormat = DefaultTimeFormat
	}
	// User Level.
	if o.LevelName != "" {
		s := strings.ToUpper(o.LevelName)
		for l, n := range Levels {
			if s == n {
				o.Level = l
				break
			}
		}
	}
	o.onStatus()
	// User Adapter.
	if o.AdapterName != "" {
		s := strings.ToUpper(o.AdapterName)
		for a, n := range Adapters {
			if s == n {
				o.Adapter = a
				break
			}
		}
	}
	// Handler.
	switch o.Adapter {
	case AdapterTerm:
		o.SetHandler(newAdapterTerm().handler)
	case AdapterFile:
		o.SetHandler(newAdapterFile().handler)
	case AdapterRedis:
		o.SetHandler(newAdapterRedis().handler)
	}
}

func (o *configuration) onStatus() {
	o.debugOn = o.Level >= LevelDebug
	o.infoOn = o.Level >= LevelInfo
	o.warnOn = o.Level >= LevelWarn
	o.errorOn = o.Level >= LevelError
}
