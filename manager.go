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
// date: 2023-03-05

package log

import (
	"context"
	"github.com/fuyibing/log/v5/common"
	"github.com/fuyibing/log/v5/configurer"
	"github.com/fuyibing/log/v5/loggers"
	"github.com/fuyibing/log/v5/tracers"
	"github.com/fuyibing/util/v8/process"
	"sync"
	"time"
)

var (
	// Manager
	// 管理器.
	Manager Management
)

type (
	// Management
	// 管理器接口.
	Management interface {
		// Config
		// 全局配置.
		//
		// 此为读模式, 通过此方法读取全局配置.
		Config() configurer.Configuration

		// Logger
		// 日志操作接口.
		Logger() loggers.OperatorManager

		// Tracer
		// 链路操作接口.
		Tracer() tracers.OperatorManager

		// Start
		// 启动管理器.
		//
		// 当触发启动时将阻塞协程(10ms), 直到设置的日志与链路执行器启动成功(异步批处
		// 理).
		Start(ctx context.Context)

		// Stop
		// 安全退出.
		//
		// 当触发退出时将阻塞协程(30秒超时), 直到设置的日志与链路执行器全部处理完成,
		// 此过程可以避免数据丢失.
		Stop()
	}

	manager struct {
		config configurer.Configuration
		logger loggers.OperatorManager
		tracer tracers.OperatorManager

		name      string
		processor process.Processor
	}
)

// /////////////////////////////////////////////////////////////////////////////
// Interface methods
// /////////////////////////////////////////////////////////////////////////////

func (o *manager) Config() configurer.Configuration { return o.config }
func (o *manager) Logger() loggers.OperatorManager  { return o.logger }
func (o *manager) Tracer() tracers.OperatorManager  { return o.tracer }

func (o *manager) Start(ctx context.Context) { o.start(ctx) }
func (o *manager) Stop()                     { o.stop() }

// /////////////////////////////////////////////////////////////////////////////
// Event methods
// /////////////////////////////////////////////////////////////////////////////

func (o *manager) onBeforeLogger(_ context.Context) (ignored bool) {
	// 基于编码
	if ex := o.logger.GetExecutor(); ex != nil {
		common.InternalInfo(`<%s> logger executor [name="%s"][level="%s"]`,
			o.name, ex.Processor().Name(),
			configurer.Config.GetLoggerLevel(),
		)

		// 加为子进程
		if _, exists := o.processor.Get(ex.Processor().Name()); !exists {
			o.processor.Add(ex.Processor())
		}
		return
	}

	// 基于配置
	if call, ok := builtinLoggers[configurer.Config.GetLoggerExporter()]; ok {
		if ex := call(); ex != nil {
			common.InternalInfo(`<%s> logger executor [name="%s"][level="%s"]`,
				o.name, ex.Processor().Name(),
				configurer.Config.GetLoggerLevel(),
			)

			// 加为子进程
			o.logger.SetExecutor(ex)
			o.processor.Add(ex.Processor())
		}
	}

	return
}

func (o *manager) onBeforeTracer(_ context.Context) (ignored bool) {
	// 基于编码
	if ex := o.tracer.GetExecutor(); ex != nil {
		common.InternalInfo(`<%s> tracer executor [name="%s"][topic="%s"]`,
			o.name, ex.Processor().Name(),
			configurer.Config.GetTracerTopic(),
		)

		// 加为子进程
		if _, exists := o.processor.Get(ex.Processor().Name()); !exists {
			o.processor.Add(ex.Processor())
		}
		return
	}

	// 基于配置
	if call, ok := builtinTracers[configurer.Config.GetTracerExporter()]; ok {
		if ex := call(); ex != nil {
			common.InternalInfo(`<%s> tracer executor [name="%s"][topic="%s"]`,
				o.name, ex.Processor().Name(),
				configurer.Config.GetTracerTopic(),
			)

			// 加为子进程
			o.tracer.SetExecutor(ex)
			o.processor.Add(ex.Processor())
		}
	}

	return
}

func (o *manager) onCall(ctx context.Context) (ignored bool) {
	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}

func (o *manager) onPanic(_ context.Context, v interface{}) {
	common.InternalFatal("<%s> fatal: %v", o.name, v)
}

// /////////////////////////////////////////////////////////////////////////////
// Access and constructor
// /////////////////////////////////////////////////////////////////////////////

func (o *manager) init() *manager {
	o.config = configurer.Config
	o.logger = loggers.Operator
	o.tracer = tracers.Operator

	o.name = "manager"
	o.processor = process.New(o.name).
		Before(o.onBeforeLogger, o.onBeforeTracer).
		Callback(o.onCall).
		Panic(o.onPanic)

	return o
}

func (o *manager) start(ctx context.Context) {
	// 并行启动.
	go func() {
		common.InternalInfo("<%s> start", o.name)

		if err := o.processor.Start(ctx); err != nil {
			common.InternalInfo("<%s> stopped: %v", o.name, err)
		} else {
			common.InternalInfo("<%s> stopped", o.name)
		}
	}()

	// 等待完成.
	mx := 10
	ms := time.Millisecond
	for i := 0; i < mx; i++ {
		time.Sleep(ms)

		if func() bool {
			if o.logger.GetExecutor() != nil {
				return o.logger.GetExecutor().Processor().Healthy()
			}
			return true
		}() && func() bool {
			if o.tracer.GetExecutor() != nil {
				return o.tracer.GetExecutor().Processor().Healthy()
			}
			return true
		}() {
			common.InternalInfo("<%s> started", o.name)
			break
		}
	}
}

func (o *manager) stop() {
	// 并行退出.
	o.processor.Stop()

	// 等待完成.
	mx := 300
	ms := time.Millisecond * 100
	for i := 0; i < mx; i++ {
		time.Sleep(ms)

		if func() bool {
			if o.logger.GetExecutor() != nil {
				return o.logger.GetExecutor().Processor().Stopped()
			}
			return true
		}() && func() bool {
			if o.tracer.GetExecutor() != nil {
				return o.tracer.GetExecutor().Processor().Stopped()
			}
			return true
		}() {
			break
		}
	}
}

func init() { new(sync.Once).Do(func() { Manager = (&manager{}).init() }) }
