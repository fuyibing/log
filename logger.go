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

package log

import (
	"context"
	"github.com/fuyibing/log/v5/base"
	"github.com/fuyibing/log/v5/conf"
	"github.com/fuyibing/log/v5/exporters"
	"github.com/fuyibing/util/v8/process"
	"sync"
	"time"
)

var (
	// Logger 日志管理.
	Logger LoggerManager
)

type (
	// LoggerManager 日志管理器.
	LoggerManager interface {
		// Start 启动日志.
		Start(ctx context.Context)

		// Stop 退出日志.
		Stop()
	}

	logger struct {
		name          string
		processor     process.Processor
		processorLog  process.Processor
		processorSpan process.Processor
	}
)

// Start 启动日志.
func (o *logger) Start(ctx context.Context) {
	// 并行启动.
	go func() {
		if err := o.processor.Start(ctx); err != nil {
			base.InternalError("<%s> start: %v", o.name, err)
		}
	}()

	// 等待完备.
	time.Sleep(time.Millisecond * 3)
	base.InternalInfo("<%s> started", o.name)
}

// Stop 退出日志.
func (o *logger) Stop() {
	// 并行退出.
	o.processor.Stop()

	// 等待完成.
	for {
		if func() bool {
			if o.processorLog != nil {
				return o.processorLog.Stopped()
			}
			return true
		}() && func() bool {
			if o.processorSpan != nil {
				return o.processorSpan.Stopped()
			}
			return true
		}() {
			base.InternalInfo("<%s> stopped", o.name)
			return
		}

		time.Sleep(time.Millisecond * 10)
	}
}

// /////////////////////////////////////////////////////////////////////////////
// Access and constructor
// /////////////////////////////////////////////////////////////////////////////

func (o *logger) init() *logger {
	o.name = "logger"
	o.processor = process.New(o.name).
		Before(o.onBeforeLogExporter, o.onBeforeSpanExporter, o.onBeforeUpdate).
		Callback(o.onCall).
		Panic(o.onPanic)
	return o
}

func (o *logger) onBeforeLogExporter(_ context.Context) (ignored bool) {
	var exporter base.LoggerExporter

	defer func() {
		if exporter != nil {
			base.InternalInfo("<%s> register logger: level=%s, exporter=%s, configured=%s",
				o.name,
				conf.Config.GetLoggerLevel(),
				exporter.Processor().Name(),
				conf.Config.GetLoggerExporter(),
			)
		}
	}()

	// 已经配置.
	if exporter = exporters.Exporter.GetLogger(); exporter != nil {
		o.processorLog = exporter.Processor()
		o.processor.Add(o.processorLog)
		return
	}

	// 默认配置.
	if callable := exporters.BuiltinLog(conf.Config.GetLoggerExporter()).Callable(); callable != nil {
		if exporter = callable(); exporter != nil {
			o.processorLog = exporter.Processor()
			o.processor.Add(o.processorLog)

			exporters.Exporter.SetLogger(exporter)
			return
		}
	}

	return
}

func (o *logger) onBeforeSpanExporter(_ context.Context) (ignored bool) {
	var exporter base.TracerExporter

	defer func() {
		if exporter != nil {
			base.InternalInfo("<%s> register tracer: topic=%s, exporter=%s, configured=%s",
				o.name,
				conf.Config.GetTracerTopic(),
				exporter.Processor().Name(),
				conf.Config.GetTracerExporter(),
			)
		}
	}()

	// 已经配置.
	if exporter = exporters.Exporter.GetTracer(); exporter != nil {
		o.processorSpan = exporter.Processor()
		o.processor.Add(o.processorSpan)
		return
	}

	// 默认配置.
	if callable := exporters.BuiltinSpan(conf.Config.GetTracerExporter()).Callable(); callable != nil {
		if exporter = callable(); exporter != nil {
			o.processorSpan = exporter.Processor()
			o.processor.Add(o.processorSpan)

			exporters.Exporter.SetTracer(exporter)
			return
		}
	}

	return
}

func (o *logger) onBeforeUpdate(_ context.Context) (ignored bool) {
	base.Resource.
		Add(base.ResourceServiceName, conf.Config.GetServiceName()).
		Add(base.ResourceServicePort, conf.Config.GetServicePort()).
		Add(base.ResourceServiceVersion, conf.Config.GetServiceVersion())
	return
}

func (o *logger) onCall(ctx context.Context) (ignored bool) {
	select {
	case <-ctx.Done():
		return
	}
}

func (o *logger) onPanic(_ context.Context, v interface{}) {
	base.InternalError("<%s> %v", o.name, v)
}

func init() { new(sync.Once).Do(func() { Logger = (&logger{}).init() }) }
