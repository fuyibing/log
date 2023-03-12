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
	Config Configuration
)

type (
	Configuration interface {
		ConfigBucket
		ConfigOpenTracing

		ConfigLogger
		ConfigLoggerFile

		ConfigTracer
		ConfigTracerJaeger
		ConfigTracerZipkin
		ConfigTracerFile

		Setter() *Setter
	}

	Setter struct {
		config *config
	}

	config struct {
		OpenTracingSampled string `yaml:"open-tracing-sampled"`
		OpenTracingSpanId  string `yaml:"open-tracing-span-id"`
		OpenTracingTraceId string `yaml:"open-tracing-trace-id"`

		// +-------------------------------------------------------------------+
		// | Bucket for ASYNC                                                  |
		// +-------------------------------------------------------------------+

		// Batch count.
		// Default: 100
		BucketBatch int `yaml:"bucket-batch"`

		// Batch concurrency.
		// Default: 10 (goroutines)
		BucketConcurrency int32 `yaml:"bucket-concurrency"`

		// Bucket capacity.
		// Default: 30,000
		BucketCapacity int `yaml:"bucket-capacity"`

		// Upload span per 200 ms if not triggered.
		// Default: 200 (Millisecond)
		BucketFrequency int `yaml:"bucket-frequency"`

		// +-------------------------------------------------------------------+
		// | Logger                                                            |
		// +-------------------------------------------------------------------+

		// Logger level.
		// Default: INFO.
		LoggerLevel common.Level `yaml:"logger-level"`

		// Logger name.
		// Accept: term, file.
		// Default: term
		LoggerExporter string `yaml:"logger-exporter"`

		// Save custom log to local files.
		FileLogger *fileLogger `yaml:"file-logger"`

		// +-------------------------------------------------------------------+
		// | Tracer                                                            |
		// +-------------------------------------------------------------------+

		// Tracer topic.
		// Tracer span storages.
		TracerTopic string `yaml:"tracer-topic"`

		// Tracer name.
		// Accept: term, file, jaeger, zipkin
		// Default: term
		TracerExporter string `yaml:"tracer-exporter"`

		// Save span to local files.
		FileTracer *fileTracer `yaml:"file-tracer"`

		// Upload span to Jaeger.
		JaegerTracer *jaegerTracer `yaml:"jaeger-tracer"`

		// Upload span to Zipkin.
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
	// Read config file then assign to configuration fields.
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
	o.setter = &Setter{config: o}

	// Init default fields.

	o.defaultBucket()
	o.defaultOpenTracing()

	// Logger definitions.

	o.defaultLogger()
	o.initFileLogger()

	// Tracer{file|jaeger|zipkin}

	o.defaultTracer()
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
