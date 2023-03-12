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
	Manager Management
)

type (
	Management interface {
		// Config
		// global configurations, readonly.
		Config() configurer.Configuration

		// Logger
		// log operator.
		Logger() loggers.OperatorManager

		// Tracer
		// trace operator.
		Tracer() tracers.OperatorManager

		// Start
		// start boot manager as async mode.
		Start(ctx context.Context)

		// Stop
		// send stop signal then wait all log and span push completed. Force
		// stop after 30 seconds.
		Stop()
	}

	manager struct {
		config    configurer.Configuration
		logger    loggers.OperatorManager
		name      string
		processor process.Processor
		tracer    tracers.OperatorManager
	}
)

// /////////////////////////////////////////////////////////////////////////////
// Interface methods
// /////////////////////////////////////////////////////////////////////////////

func (o *manager) Config() configurer.Configuration { return o.config }
func (o *manager) Logger() loggers.OperatorManager  { return o.logger }
func (o *manager) Tracer() tracers.OperatorManager  { return o.tracer }
func (o *manager) Start(ctx context.Context)        { o.start(ctx) }
func (o *manager) Stop()                            { o.stop() }

// /////////////////////////////////////////////////////////////////////////////
// Event methods
// /////////////////////////////////////////////////////////////////////////////

func (o *manager) onBeforeLogger(_ context.Context) (ignored bool) {
	// Add logger exporter as child process which configured by user code.
	if ex := o.logger.GetExecutor(); ex != nil {
		common.InternalInfo(`<%s> logger executor [name="%s"][level="%s"]`,
			o.name, ex.Processor().Name(),
			configurer.Config.GetLoggerLevel(),
		)

		if _, exists := o.processor.Get(ex.Processor().Name()); !exists {
			o.processor.Add(ex.Processor())
		}
		return
	}

	// Add logger exporter as child process which configured by config file.
	if call, ok := builtinLoggers[configurer.Config.GetLoggerExporter()]; ok {
		if ex := call(); ex != nil {
			common.InternalInfo(`<%s> logger executor [name="%s"][level="%s"]`,
				o.name, ex.Processor().Name(),
				configurer.Config.GetLoggerLevel(),
			)

			o.logger.SetExecutor(ex)
			o.processor.Add(ex.Processor())
		}
	}

	return
}

func (o *manager) onBeforeTracer(_ context.Context) (ignored bool) {
	// Add tracer exporter as child process which configured by user code.
	if ex := o.tracer.GetExecutor(); ex != nil {
		common.InternalInfo(`<%s> tracer executor [name="%s"][topic="%s"]`,
			o.name, ex.Processor().Name(),
			configurer.Config.GetTracerTopic(),
		)

		if _, exists := o.processor.Get(ex.Processor().Name()); !exists {
			o.processor.Add(ex.Processor())
		}
		return
	}

	// Add tracer exporter as child process which configured by config file.
	if call, ok := builtinTracers[configurer.Config.GetTracerExporter()]; ok {
		if ex := call(); ex != nil {
			common.InternalInfo(`<%s> tracer executor [name="%s"][topic="%s"]`,
				o.name, ex.Processor().Name(),
				configurer.Config.GetTracerTopic(),
			)

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

	o.name = "manager"
	o.processor = process.New(o.name).
		Before(o.onBeforeLogger, o.onBeforeTracer).
		Callback(o.onCall).
		Panic(o.onPanic)
	o.tracer = tracers.Operator

	return o
}

func (o *manager) start(ctx context.Context) {
	go func() {
		common.InternalInfo("<%s> start", o.name)
		if err := o.processor.Start(ctx); err != nil {
			common.InternalInfo("<%s> stopped: %v", o.name, err)
		} else {
			common.InternalInfo("<%s> stopped", o.name)
		}
	}()

	// Wait all processors started.
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
	// Send stop signal.
	o.processor.Stop()

	// Wait all log and span exported done or timed out.
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
