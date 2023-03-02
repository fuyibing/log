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

package tracer

import (
	"github.com/fuyibing/log/v5/base"
	"github.com/fuyibing/log/v5/conf"
	"github.com/fuyibing/log/v5/exporters"
)

type (
	// SpanLogger 跨度日志.
	SpanLogger struct {
		Attribute base.Attribute
		Span      *Span
	}
)

// Add 添加 Key/Value 属性.
func (o *SpanLogger) Add(key string, value interface{}) base.SpanLogger {
	o.Attribute.Add(key, value)
	return o
}

// Debug 设置DEBUG级日志.
func (o *SpanLogger) Debug(text string, args ...interface{}) {
	if conf.Config.DebugOn() {
		o.Send(base.Debug, text, args...)
	}
}

// Error 设置ERROR级日志.
func (o *SpanLogger) Error(text string, args ...interface{}) {
	if conf.Config.ErrorOn() {
		o.Send(base.Error, text, args...)
	}
}

// Fatal 设置FATAL级日志.
func (o *SpanLogger) Fatal(text string, args ...interface{}) {
	if conf.Config.FatalOn() {
		o.Send(base.Fatal, text, args...)
	}
}

// Info 设置INFO级日志.
func (o *SpanLogger) Info(text string, args ...interface{}) {
	if conf.Config.InfoOn() {
		o.Send(base.Info, text, args...)
	}
}

// Warn 设置WARN级日志.
func (o *SpanLogger) Warn(text string, args ...interface{}) {
	if conf.Config.WarnOn() {
		o.Send(base.Warn, text, args...)
	}
}

// /////////////////////////////////////////////////////////////////////////////
// Internal methods
// /////////////////////////////////////////////////////////////////////////////

func (o *SpanLogger) Send(level base.Level, text string, args ...interface{}) {
	o.SendLog(level, text, args...)
	o.SendSpan(level, text, args...)
}

func (o *SpanLogger) SendLog(level base.Level, text string, args ...interface{}) {
	x := NewLog(level, text, args...)
	x.Attribute.Copy(o.Attribute)
	exporters.Exporter.PutLogger(x)
}

func (o *SpanLogger) SendSpan(level base.Level, text string, args ...interface{}) {
	x := NewLog(level, text, args...)
	x.Attribute.Copy(o.Attribute)
	o.Span.addLog(x)
}

// /////////////////////////////////////////////////////////////////////////////
// Access and constructor
// /////////////////////////////////////////////////////////////////////////////

func (o *SpanLogger) init() *SpanLogger {
	o.Attribute = base.Attribute{}
	return o
}
