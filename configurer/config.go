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
// date: 2023-03-03

package configurer

import (
	"github.com/fuyibing/log/v5/common"
	"gopkg.in/yaml.v3"
	"os"
	"sync"
)

var (
	// Config
	// 全部配置.
	Config Configuration
)

type (
	// Configuration
	// 全部配置接口.
	Configuration interface {
		ConfigBucket
		ConfigLogger
		ConfigLoggerFile
		ConfigTracer
		ConfigTracerJaeger
		ConfigTracerZipkin
		ConfigFileTracer
		ConfigOpenTracing

		// Setter
		// 设置配置参数.
		Setter() *Setter
	}

	// Setter
	// 设置字段.
	Setter struct {
		config *config
	}

	// 配置结构.
	config struct {
		OpenTracingSampled string `yaml:"open-tracing-sampled"`
		OpenTracingSpanId  string `yaml:"open-tracing-span-id"`
		OpenTracingTraceId string `yaml:"open-tracing-trace-id"`

		// +-------------------------------------------------------------------+
		// | Bucket for ASYNC                                                  |
		// +-------------------------------------------------------------------+

		// 批处理量.
		// 每个批次最大 Logger/Span 数量.
		// 默认: 100
		BucketBatch int `yaml:"bucket-batch"`

		// 批处理并发.
		// 最大允许多少个协程同时上报.
		// 默认: 10
		BucketConcurrency int32 `yaml:"bucket-concurrency"`

		// 数据桶容量.
		// 当瞬间数量太多且来不急处理时, 先暂存在内存中, 此配置定义最多在内存中存储数量.
		// 默认: 30,000
		BucketCapacity int `yaml:"bucket-capacity"`

		// 批处理频率.
		// 每 200 毫秒自动上报积压的数据.
		// 默认: 200 (毫秒)
		BucketFrequency int `yaml:"bucket-frequency"`

		// +-------------------------------------------------------------------+
		// | Logger                                                            |
		// +-------------------------------------------------------------------+

		// 日志级别.
		// 默认: INFO.
		LoggerLevel common.Level `yaml:"logger-level"`

		// 日志上报适配.
		// 接受: term, file.
		// 默认: term
		LoggerExporter string `yaml:"logger-exporter"`

		// 文件日志.
		FileLogger *fileLogger `yaml:"file-logger"`

		// +-------------------------------------------------------------------+
		// | Tracer                                                            |
		// +-------------------------------------------------------------------+

		// 链路上报主题.
		// 说明: 应用于 Kafka主题, Jaeger服务名等, 与 TracerExporter 配合使用.
		TracerTopic string `yaml:"tracer-topic"`

		// 链路上报适配.
		// 接受: term, file, jaeger, zipkin
		// 默认: term
		TracerExporter string `yaml:"tracer-exporter"`

		// 输出到 File.
		FileTracer *fileTracer `yaml:"file-tracer"`

		// 上报到 Jaeger.
		JaegerTracer *jaegerTracer `yaml:"jaeger-tracer"`

		// 上报到 Zipkin.
		ZipkinTracer *zipkinTracer `yaml:"zipkin-tracer"`

		// +-------------------------------------------------------------------+
		// | Internal                                                          |
		// +-------------------------------------------------------------------+

		setter                                    *Setter
		debugOn, infoOn, warnOn, errorOn, fatalOn bool
	}
)

func (o *config) Setter() *Setter { return o.setter }

// /////////////////////////////////////////////////////////////////////////////
// Access and constructor
// /////////////////////////////////////////////////////////////////////////////

func (o *config) scan() {
	// 解析文件.
	// 包加载时扫描并读取YAML配置文件, 并基于配置赋值.
	for _, path := range []string{"config/log.yaml", "../config/log.yaml"} {
		if buf, err := os.ReadFile(path); err == nil {
			if yaml.Unmarshal(buf, o) == nil {
				return
			}
		}
	}
}

func (o *config) init() *config {
	// 初始化.
	o.setter = &Setter{config: o}
	o.scan()

	// 默认值.
	o.defaultBucket()
	o.defaultLogger()
	o.defaultTracer()
	o.defaultOpenTracing()

	// Logger{file}.
	o.initFileLogger()

	// Tracer{file|jaeger|zipkin}

	o.initFileTracer()
	o.initJaegerTracer()
	o.initZipkinTracer()
	return o
}

func (o *config) initFileLogger() {
	if o.FileLogger == nil {
		o.FileLogger = &fileLogger{}
	}
	o.FileLogger.initDefaults()
}

func (o *config) initFileTracer() {
	if o.FileTracer == nil {
		o.FileTracer = &fileTracer{}
	}
	o.FileTracer.initDefaults()
}

func (o *config) initJaegerTracer() {
	if o.JaegerTracer == nil {
		o.JaegerTracer = &jaegerTracer{}
	}
	o.JaegerTracer.initDefaults()
}

func (o *config) initZipkinTracer() {
	if o.ZipkinTracer == nil {
		o.ZipkinTracer = &zipkinTracer{}
	}
	o.ZipkinTracer.initDefaults()
}

func init() { new(sync.Once).Do(func() { Config = (&config{}).init() }) }
