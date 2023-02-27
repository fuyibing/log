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
	// 用于Span的Logger操作组件.
	SpanLogger struct {
		attr Attr
		span Span
	}
)

// NewSpanLogger
// 创建SpanLogger组件.
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
// 添加日志KV属性.
func (o *SpanLogger) Add(key string, value interface{}) *SpanLogger {
	o.attr.Add(key, value)
	return o
}

// Debug
// 记录为 debug 级日志.
func (o *SpanLogger) Debug(text string, args ...interface{}) {
	if conf.Config.DebugOn() {
		o.send(base.Debug, text, args...)
	}
}

// Error
// 记录为 error 级日志.
func (o *SpanLogger) Error(text string, args ...interface{}) {
	if conf.Config.ErrorOn() {
		o.send(base.Error, text, args...)
	}
}

// Fatal
// 记录为 fatal 级日志.
func (o *SpanLogger) Fatal(text string, args ...interface{}) {
	if conf.Config.FatalOn() {
		o.send(base.Fatal, text, args...)
	}
}

// Info
// 记录为 info 级日志.
func (o *SpanLogger) Info(text string, args ...interface{}) {
	if conf.Config.InfoOn() {
		o.send(base.Info, text, args...)
	}
}

// Warn
// 记录为 warning 级日志.
func (o *SpanLogger) Warn(text string, args ...interface{}) {
	if conf.Config.WarnOn() {
		o.send(base.Warn, text, args...)
	}
}

// send
// 发送Log数据.
func (o *SpanLogger) send(level base.Level, text string, args ...interface{}) {
	// 创建用户Log并加到Span组件中.
	x := NewLine(level, text, args...)
	x.GetAttr().Copy(o.attr)
	o.span.GetLogs().Add(x)

	// 同时推送普通Log.
	if Registry.LoggerEnabled() {
		Registry.LoggerPush(o.attr, level, text, args...)
	}

	// 释放实例回池.
	o.attr = nil
	o.span = nil
	spanLoggerPool.Put(o)
}
