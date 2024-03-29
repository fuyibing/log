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
// date: 2023-03-04

package tracers

import (
	"github.com/fuyibing/log/v5/common"
	"github.com/fuyibing/log/v5/configurer"
	"github.com/fuyibing/log/v5/loggers"
	"sync"
)

var spanLoggerPool = sync.Pool{}

type (
	// SpanLogger
	// log component for span.
	SpanLogger interface {
		Add(key string, value interface{}) SpanLogger
		Debug(text string, args ...interface{})
		Error(text string, args ...interface{})
		Fatal(text string, args ...interface{})
		Info(text string, args ...interface{})
		Warn(text string, args ...interface{})
	}

	spanLogger struct {
		sync.RWMutex

		kv   loggers.Kv
		span *span
	}
)

func (o *spanLogger) Add(key string, value interface{}) SpanLogger { return o.add(key, value) }
func (o *spanLogger) Debug(format string, args ...interface{})     { o.send(common.Debug, format, args...) }
func (o *spanLogger) Info(format string, args ...interface{})      { o.send(common.Info, format, args...) }
func (o *spanLogger) Warn(format string, args ...interface{})      { o.send(common.Warn, format, args...) }
func (o *spanLogger) Error(format string, args ...interface{})     { o.send(common.Error, format, args...) }
func (o *spanLogger) Fatal(format string, args ...interface{})     { o.send(common.Fatal, format, args...) }

// /////////////////////////////////////////////////////////////////////////////
// Access and constructor
// /////////////////////////////////////////////////////////////////////////////

func (o *spanLogger) add(key string, value interface{}) SpanLogger {
	o.Lock()
	defer o.Unlock()

	o.kv.Add(key, value)
	return o
}

func (o *spanLogger) after() {
	o.kv = nil
	o.span = nil
}

func (o *spanLogger) before(span *span) {
	o.kv = loggers.Kv{}
	o.span = span
}

func (o *spanLogger) send(level common.Level, format string, args ...interface{}) {
	// Push to logger executor.
	loggers.Operator.Push(o.kv, level, format, args...)

	// Push to tracer executor.
	if configurer.Config.LevelEnabled(level) {
		log := loggers.NewLog(level, format, args...)
		if len(o.kv) > 0 {
			log.SetKv(o.kv)
		}
		o.span.addLog(log)
	}

	// Release when ended.
	o.after()
}

func spanLoggerAcquire(span *span) SpanLogger {
	if g := spanLoggerPool.Get(); g != nil {
		v := g.(*spanLogger)
		v.before(span)
		return v
	}

	v := &spanLogger{}
	v.before(span)
	return v
}
