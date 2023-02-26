// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// author: wsfuyibing <websearch@163.com>
// date: 2023-02-25

package conf

import (
	"github.com/fuyibing/log/v5/base"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type (
	Configuration interface {
		GetJaegerTracer() JaegerTracerConfiguration

		GetLoggerExporter() string
		GetLoggerLevel() base.Level

		GetOpenTracingSample() string
		GetOpenTracingSpanId() string
		GetOpenTracingTraceId() string

		GetServiceName() string
		GetServicePort() int
		GetServiceVersion() string

		GetTracerExporter() string
		GetTracerTopic() string

		DebugOn() bool
		ErrorOn() bool
		FatalOn() bool
		InfoOn() bool
		WarnOn() bool

		With(opts ...Option)
	}

	configuration struct {
		LoggerExporter string     `yaml:"logger-exporter"`
		LoggerLevel    base.Level `yaml:"logger-level"`

		OpenTracingSample  string `yaml:"open-tracing-sample"`
		OpenTracingSpanId  string `yaml:"open-tracing-span-id"`
		OpenTracingTraceId string `yaml:"open-tracing-trace-id"`

		ServiceName    string `yaml:"service-name"`
		ServicePort    int    `yaml:"service-port"`
		ServiceVersion string `yaml:"service-version"`

		TracerExporter string `yaml:"tracer-exporter"`
		TracerTopic    string `yaml:"tracer-topic"`

		JaegerTracer *jaegerTracerConfiguration `yaml:"jaeger-tracer"`

		debugOn, infoOn, warnOn, errorOn, fatalOn bool
	}
)

// /////////////////////////////////////////////////////////////////////////////
// Configuration: logger fields
// /////////////////////////////////////////////////////////////////////////////

func (o *configuration) GetLoggerExporter() string  { return o.LoggerExporter }
func (o *configuration) GetLoggerLevel() base.Level { return o.LoggerLevel }

// /////////////////////////////////////////////////////////////////////////////
// Configuration: get open tracing definitions
// /////////////////////////////////////////////////////////////////////////////

func (o *configuration) GetOpenTracingSample() string  { return o.OpenTracingSample }
func (o *configuration) GetOpenTracingSpanId() string  { return o.OpenTracingSpanId }
func (o *configuration) GetOpenTracingTraceId() string { return o.OpenTracingTraceId }

func (o *configuration) GetServiceName() string    { return o.ServiceName }
func (o *configuration) GetServicePort() int       { return o.ServicePort }
func (o *configuration) GetServiceVersion() string { return o.ServiceVersion }

// /////////////////////////////////////////////////////////////////////////////
// Configuration: tracer
// /////////////////////////////////////////////////////////////////////////////

func (o *configuration) GetTracerExporter() string { return o.TracerExporter }
func (o *configuration) GetTracerTopic() string    { return o.TracerTopic }

// /////////////////////////////////////////////////////////////////////////////
// Configuration: log state
// /////////////////////////////////////////////////////////////////////////////

func (o *configuration) DebugOn() bool { return o.debugOn }
func (o *configuration) ErrorOn() bool { return o.errorOn }
func (o *configuration) FatalOn() bool { return o.fatalOn }
func (o *configuration) InfoOn() bool  { return o.infoOn }
func (o *configuration) WarnOn() bool  { return o.warnOn }

func (o *configuration) With(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

// /////////////////////////////////////////////////////////////////////////////
// Configuration: children
// /////////////////////////////////////////////////////////////////////////////

func (o *configuration) GetJaegerTracer() JaegerTracerConfiguration { return o.JaegerTracer }

// /////////////////////////////////////////////////////////////////////////////
// Configuration: access and constructor
// /////////////////////////////////////////////////////////////////////////////

func (o *configuration) init() *configuration {
	o.scan()

	o.initDefaults()
	o.initChild()

	o.updateState()
	return o
}

func (o *configuration) initChild() *configuration {
	// Tracer: jaeger.
	if o.JaegerTracer == nil {
		o.JaegerTracer = &jaegerTracerConfiguration{}
	}

	// Tracer: init jaeger defaults
	o.JaegerTracer.initDefaults()
	return o
}

func (o *configuration) initDefaults() {
	// OpenTracing: Sample
	if o.OpenTracingSample == "" {
		o.OpenTracingSample = base.OpenTracingSample
	}

	// OpenTracing: SpanId
	if o.OpenTracingSpanId == "" {
		o.OpenTracingSpanId = base.OpenTracingSpanId
	}

	// OpenTracing: TraceId
	if o.OpenTracingTraceId == "" {
		o.OpenTracingTraceId = base.OpenTracingTraceId
	}

	// LogLevel: is lower string verify.
	if s := o.LoggerLevel.String(); s != "" {
		o.LoggerLevel = base.Level(strings.ToUpper(s))
	}

	// LogLevel: use default level.
	if o.LoggerLevel.Int() == 0 {
		o.LoggerLevel = base.Info
	}
}

// scan read yaml content from config/log.yaml, then assign fields value on
// configuration.
func (o *configuration) scan() {
	for _, path := range []string{"config/log.yaml", "../config/log.yaml"} {
		if buf, err := os.ReadFile(path); err == nil {
			if yaml.Unmarshal(buf, o) == nil {
				return
			}
		}
	}
}

func (o *configuration) updateState() {
	li := o.LoggerLevel.Int()
	ls := li > base.Off.Int()

	o.debugOn = ls && li >= base.Debug.Int()
	o.infoOn = ls && li >= base.Info.Int()
	o.warnOn = ls && li >= base.Warn.Int()
	o.errorOn = ls && li >= base.Error.Int()
	o.fatalOn = ls && li >= base.Fatal.Int()
}
