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
// date: 2023-03-01

package conf

import (
	"github.com/fuyibing/log/v5/traces"
	"gopkg.in/yaml.v3"
	"os"
)

var (
	Config ConfigManager
)

type (
	// ConfigManager
	// 配置管理器.
	ConfigManager interface {
		GetServiceName() string
		GetServicePort() int
		GetServiceVersion() string

		ConfigBucket
		ConfigOpenTracing
		ConfigLogger
		FileLogger

		ConfigTracer
		JaegerTracer
	}

	// config
	// 配置字段.
	config struct {
		// 批量阈值.
		// 异步批量上报时, 每个批次最大值.
		BucketBatch int `yaml:"bucket-batch"`

		// 并行阈值.
		// 异步上报时, 最大允许并行协程数.
		BucketConcurrency int32 `yaml:"bucket-concurrency"`

		// 数据桶容量.
		// 内存队列(数据桶)最大允许积压数量.
		BucketCapacity int `yaml:"bucket-capacity"`

		// 数据桶频率.
		// 每隔多少时间(毫秒)检查数据桶是否有积压.
		BucketFrequency int `yaml:"bucket-frequency"`

		LoggerExporter string       `yaml:"logger-exporter"`
		LoggerLevel    traces.Level `yaml:"logger-level"`

		OpenTracingSampled string `yaml:"open-tracing-sampled"`
		OpenTracingSpanId  string `yaml:"open-tracing-span-id"`
		OpenTracingTraceId string `yaml:"open-tracing-trace-id"`

		ServiceName    string `yaml:"service-name"`
		ServicePort    int    `yaml:"service-port"`
		ServiceVersion string `yaml:"service-version"`

		TracerExporter string `yaml:"tracer-exporter"`
		TracerTopic    string `yaml:"tracer-topic"`

		FileLogger   *fileLogger   `yaml:"file-logger"`
		JaegerTracer *jaegerTracer `yaml:"jaeger-tracer"`
	}
)

func (o *config) GetServiceName() string    { return o.ServiceName }
func (o *config) GetServicePort() int       { return o.ServicePort }
func (o *config) GetServiceVersion() string { return o.ServiceVersion }

// /////////////////////////////////////////////////////////////////////////////
// Access and constructor
// /////////////////////////////////////////////////////////////////////////////

func (o *config) scan() {
	for _, path := range []string{"config/log.yaml", "../config/log.yaml"} {
		if buf, err := os.ReadFile(path); err == nil {
			if yaml.Unmarshal(buf, o) == nil {
				return
			}
		}
	}
}

func (o *config) init() *config {
	o.scan()
	o.initLoggerDefaults()
	o.initDefaultsOpenTracing()
	o.initTracerDefaults()
	o.initBucketDefaults()

	o.initFileLogger()
	o.initJaegerTracer()
	return o
}

func (o *config) initFileLogger() {
	if o.FileLogger == nil {
		o.FileLogger = &fileLogger{}
	}
	o.FileLogger.initDefaults()
}

func (o *config) initJaegerTracer() {
	if o.JaegerTracer == nil {
		o.JaegerTracer = &jaegerTracer{}
	}
	o.JaegerTracer.initDefaults()
}
