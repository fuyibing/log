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
// date: 2023-02-26

package cores

import (
	"github.com/fuyibing/log/v5/base"
	"github.com/fuyibing/log/v5/conf"
	"sync"
)

var (
	spanLoggerPool sync.Pool
)

type (
	// SpanLogger
	// is the component for record custom log message.
	SpanLogger struct {
		attr Attr
		span Span
	}
)

// NewSpanLogger
// return a SpanLogger component to record logs.
func NewSpanLogger(span Span) (sl *SpanLogger) {
	if v := spanLoggerPool.Get(); v != nil {
		sl = v.(*SpanLogger)
		sl.attr = NewAttr()
		sl.span = span
	} else {
		sl = &SpanLogger{
			attr: NewAttr(),
			span: span,
		}
	}
	return sl
}

// Add
// key/value pair on logger.
func (o *SpanLogger) Add(key string, value interface{}) *SpanLogger {
	o.attr.Add(key, value)
	return o
}

// Debug
// send custom message to log exporter with debug level.
func (o *SpanLogger) Debug(text string, args ...interface{}) {
	if conf.Config.DebugOn() {
		o.send(base.Debug, text, args...)
	}
}

// Error
// send custom message to log exporter with error level.
func (o *SpanLogger) Error(text string, args ...interface{}) {
	if conf.Config.ErrorOn() {
		o.send(base.Error, text, args...)
	}
}

// Fatal
// send custom message to log exporter with fatal level.
func (o *SpanLogger) Fatal(text string, args ...interface{}) {
	if conf.Config.FatalOn() {
		o.send(base.Fatal, text, args...)
	}
}

// Info
// send custom message to log exporter with info level.
func (o *SpanLogger) Info(text string, args ...interface{}) {
	if conf.Config.InfoOn() {
		o.send(base.Info, text, args...)
	}
}

// Warn
// send custom message to log exporter with warning level.
func (o *SpanLogger) Warn(text string, args ...interface{}) {
	if conf.Config.WarnOn() {
		o.send(base.Warn, text, args...)
	}
}

// send
// custom log message to exporter.
func (o *SpanLogger) send(level base.Level, text string, args ...interface{}) {
	// Create a Line component then add on SpanLogs component.
	x := NewLine(level, text, args...)
	x.GetAttr().Copy(o.attr)
	o.span.GetLogs().Add(x)

	// Call logger exporter and push if enabled.
	if Registry.LoggerEnabled() {
		v := NewLine(level, text, args...)
		v.GetAttr().Copy(o.attr)
		Registry.LoggerExporter().Push(v)
	}

	// Release to pool when ended.
	o.attr = nil
	o.span = nil
	spanLoggerPool.Put(o)
}
