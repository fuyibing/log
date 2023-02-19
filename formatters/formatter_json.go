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

	JsonLine struct {
		Content  string  `json:"content"`
		Level    string  `json:"level"`
		Start    int64   `json:"start"`
		Duration float64 `json:"duration,omitempty"`

		ParentSpanId string `json:"parent_span_id,omitempty"`
		SpanId       string `json:"span_id,omitempty"`
		SpanVersion  string `json:"span_version,omitempty"`
		TraceId      string `json:"trace_id,omitempty"`

		Property *JsonProperty `json:"property,omitempty"`
		Resource *JsonResource `json:"resource,omitempty"`
	}

	JsonProperty struct {
		HttpHeaders       map[string][]string `json:"http.headers,omitempty"`
		HttpProtocol      string              `json:"http.protocol,omitempty"`
		HttpRequestMethod string              `json:"http.request.method,omitempty"`
		HttpRequestUrl    string              `json:"http.request.url,omitempty"`
		HttpUserAgent     string              `json:"http.user.agent,omitempty"`
	}

	JsonResource struct {
		DeployAddress     string `json:"deploy.address,omitempty"`
		DeployEnvironment string `json:"deploy.environment,omitempty"`
		DeployPort        int    `json:"deploy.port,omitempty"`
		ProcessId         int    `json:"process.pid,omitempty"`
		ServiceName       string `json:"service.name,omitempty"`
		ServiceVersion    string `json:"service.version,omitempty"`
	}
)

func NewJsonFormatter() *JsonFormatter {
	return &JsonFormatter{}
}

// /////////////////////////////////////////////////////////////
// Interface methods
// /////////////////////////////////////////////////////////////

func (o *JsonFormatter) Body(line *base.Line) []byte {
	body, _ := json.Marshal((&JsonLine{}).apply(line))
	return body
}

func (o *JsonFormatter) String(line *base.Line) string {
	return string(o.Body(line))
}

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *JsonFormatter) init() *JsonFormatter { return o }

func (o *JsonLine) apply(line *base.Line) *JsonLine {
	// Basic fields.
	o.Content = line.Text
	o.Duration = line.Duration
	o.Start = line.Time.UnixMilli()
	o.Level = line.Level.String()

	// Resource fields.
	o.Resource = &JsonResource{
		DeployAddress:     conf.Config.GetServiceAddr(),
		DeployEnvironment: conf.Config.GetServiceEnvironment(),
		DeployPort:        conf.Config.GetServicePort(),
		ProcessId:         conf.Config.GetPid(),
		ServiceName:       conf.Config.GetServiceName(),
		ServiceVersion:    conf.Config.GetServiceVersion(),
	}

	// Property fields.
	o.Property = &JsonProperty{}

	// OpenTracing inherit.
	if tracing := line.Tracing(); tracing != nil {
		o.ParentSpanId = tracing.ParentSpanId
		o.SpanId = tracing.SpanId
		o.SpanVersion = tracing.GenVersion(line.TracingOffset())
		o.TraceId = tracing.TraceId

		// Inherit http fields.
		if tracing.Http {
			o.Property.HttpHeaders = tracing.HttpHeaders
			o.Property.HttpProtocol = tracing.HttpProtocol
			o.Property.HttpRequestMethod = tracing.HttpRequestMethod
			o.Property.HttpRequestUrl = tracing.HttpRequestUrl
			o.Property.HttpUserAgent = tracing.HttpUserAgent
		}
	}

	return o
}
