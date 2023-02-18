// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package formatters

import (
	"encoding/json"
	"github.com/fuyibing/log/v8/base"
	"github.com/fuyibing/log/v8/conf"
)

type (
	JsonFormatter struct{}

	// JsonLine
	// for kafka.
	JsonLine struct {
		Datetime    int64   `json:"datetime"`
		Duration    float64 `json:"duration"`
		Level       string  `json:"level"`
		Pid         int     `json:"pid"`
		ServiceHost string  `json:"service_host"`
		ServiceName string  `json:"service_name"`
		ServicePort int     `json:"service_port"`

		ParentSpanId string `json:"parent_span_id"`
		SpanId       string `json:"span_id"`
		SpanVersion  string `json:"span_version"`
		TraceId      string `json:"trace_id"`

		RequestMethod string `json:"request_method"`
		RequestUrl    string `json:"request_url"`

		Content string `json:"content"`
	}
)

func NewJsonFormatter() *JsonFormatter {
	return &JsonFormatter{}
}

// /////////////////////////////////////////////////////////////
// Interface methods
// /////////////////////////////////////////////////////////////

func (o *JsonFormatter) Body(line *base.Line) []byte {
	body, _ := json.Marshal((&JsonLine{}).init(line))
	return body
}

func (o *JsonFormatter) String(line *base.Line) string {
	return string(o.Body(line))
}

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *JsonFormatter) init() *JsonFormatter { return o }

func (o *JsonLine) init(line *base.Line) *JsonLine {
	// Basic fields
	o.Content = line.Text
	o.Datetime = line.Time.UnixMilli()
	o.Level = line.Level.String()
	o.Pid = conf.Config.GetPid()
	o.ServiceHost = conf.Config.GetServiceHost()
	o.ServiceName = conf.Config.GetServiceName()
	o.ServicePort = conf.Config.GetServicePort()

	// Open tracing fields.
	if t := line.Tracing(); t != nil {
		o.ParentSpanId = t.ParentSpanId
		o.SpanId = t.SpanId
		o.SpanVersion = t.GenVersion(line.TracingOffset())
		o.TraceId = t.TraceId

		// HTTP Request fields.
		o.RequestMethod = t.RequestMethod
		o.RequestUrl = t.RequestUrl
	}

	return o
}
