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
	// Interface
	// 暴露配置接口.
	Interface interface {
		// With
		// 绑定配置选项.
		With(opts ...Option)

		InterfaceLogger
		InterfaceOpentracing
		InterfaceService
		InterfaceTracer
	}

	// Option
	// 配置选项.
	Option func(c *configuration)

	// 配置字段.
	configuration struct {
		// Logger

		LoggerExporter string     `yaml:"logger-exporter"`
		LoggerLevel    base.Level `yaml:"logger-level"`

		// Opentracing

		OpenTracingSample  string `yaml:"open-tracing-sample"`
		OpenTracingSpanId  string `yaml:"open-tracing-span-id"`
		OpenTracingTraceId string `yaml:"open-tracing-trace-id"`

		// Service

		ServiceName    string `yaml:"service-name"`
		ServicePort    int    `yaml:"service-port"`
		ServiceVersion string `yaml:"service-version"`

		// Tracer

		TracerExporter string                     `yaml:"tracer-exporter"`
		TracerTopic    string                     `yaml:"tracer-topic"`
		JaegerTracer   *jaegerTracerConfiguration `yaml:"jaeger-tracer"`

		// State

		debugOn, infoOn, warnOn, errorOn, fatalOn bool
	}
)

// With
// 绑定配置选项.
func (o *configuration) With(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

// init
// 配置构造.
func (o *configuration) init() *configuration {
	o.scan()
	o.initDefaults()
	o.initChild()
	o.updateLevelState()
	return o
}

// initChild
// 子配置构造.
func (o *configuration) initChild() *configuration {
	if o.JaegerTracer == nil {
		o.JaegerTracer = &jaegerTracerConfiguration{}
	}

	o.JaegerTracer.initDefaults()
	return o
}

// initDefaults
// 赋默认值.
func (o *configuration) initDefaults() {
	// 调用请求链.

	if o.OpenTracingSample == "" {
		o.OpenTracingSample = base.OpenTracingSample
	}
	if o.OpenTracingSpanId == "" {
		o.OpenTracingSpanId = base.OpenTracingSpanId
	}
	if o.OpenTracingTraceId == "" {
		o.OpenTracingTraceId = base.OpenTracingTraceId
	}

	// 默认日志适配.
	if o.LoggerExporter == "" {
		o.LoggerExporter = base.DefaultLoggerExporter
	}

	// 日志名转大写.
	if s := o.LoggerLevel.String(); s != "" {
		o.LoggerLevel = base.Level(strings.ToUpper(s))
	}

	// 默认日志级别.
	if o.LoggerLevel.Int() == 0 {
		o.LoggerLevel = base.DefaultLoggerLevel
	}
}

// scan
// 扫描并读取配置文件, 将配置参数映射到对应的字段上.
func (o *configuration) scan() {
	for _, path := range []string{"config/log.yaml", "../config/log.yaml"} {
		if buf, err := os.ReadFile(path); err == nil {
			if yaml.Unmarshal(buf, o) == nil {
				return
			}
		}
	}
}

// updateLevelState
// 更新日志级别状态.
func (o *configuration) updateLevelState() {
	li := o.LoggerLevel.Int()
	lo := li > base.Off.Int()

	o.debugOn = lo && li >= base.Debug.Int()
	o.infoOn = lo && li >= base.Info.Int()
	o.warnOn = lo && li >= base.Warn.Int()
	o.errorOn = lo && li >= base.Error.Int()
	o.fatalOn = lo && li >= base.Fatal.Int()
}
