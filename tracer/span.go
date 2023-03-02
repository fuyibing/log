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
	"context"
	"github.com/fuyibing/log/v5/exporters"
	"github.com/fuyibing/log/v5/traces"
	"sync"
	"time"
)

type (
	// Span 跨度组件.
	Span struct {
		sync.RWMutex

		// 属性.
		Attribute traces.Attribute

		// 上下文.
		Ctx context.Context

		// 结束时间.
		EndTime time.Time

		// 日志列表.
		Logs []traces.Log

		// 名称.
		Name string

		// 随机ID.
		SpanId, ParentSpanId traces.SpanId

		// 开始时间.
		StartTime time.Time

		// 隶属跟踪组件.
		Trace traces.Trace
	}
)

// Child 新建子跨度.
func (o *Span) Child(name string) traces.Span {
	v := (&Span{Name: name}).init()
	v.Trace = o.Trace
	v.ParentSpanId = o.SpanId

	v.initContext(o.Ctx)
	return v
}

func (o *Span) Context() context.Context {
	return o.Ctx
}

// End 结束跨度.
func (o *Span) End() {
	o.Lock()
	o.EndTime = time.Now()
	o.Unlock()

	exporters.Exporter.PutTracer(o)
}

// GetAttribute 获取组件属性.
func (o *Span) GetAttribute() traces.Attribute {
	return o.Attribute
}

// GetDuration 获取跨度时长.
func (o *Span) GetDuration() time.Duration {
	return o.EndTime.Sub(o.StartTime)
}

// GetLogs 跨度日志列表.
func (o *Span) GetLogs() []traces.Log {
	o.RLock()
	defer o.RUnlock()
	return o.Logs
}

// GetName 跨度名称.
func (o *Span) GetName() string {
	return o.Name
}

// GetParentSpanId 获取上级跨度ID.
func (o *Span) GetParentSpanId() traces.SpanId {
	return o.ParentSpanId
}

// GetSpanId 获取跨度ID.
func (o *Span) GetSpanId() traces.SpanId {
	return o.SpanId
}

// GetStartTime 获取开始时间.
func (o *Span) GetStartTime() time.Time {
	return o.StartTime
}

// GetTrace 获取跟踪.
func (o *Span) GetTrace() traces.Trace {
	return o.Trace
}

// Logger 跨度日志.
func (o *Span) Logger() traces.SpanLogger {
	return (&SpanLogger{Span: o}).init()
}

// /////////////////////////////////////////////////////////////////////////////
// Access and constructor
// /////////////////////////////////////////////////////////////////////////////

func (o *Span) addLog(log traces.Log) {
	o.Lock()
	defer o.Unlock()
	o.Logs = append(o.Logs, log)
}

func (o *Span) init() *Span {
	o.Attribute = traces.Attribute{}
	o.Logs = make([]traces.Log, 0)
	o.SpanId = traces.Id.SpanIdNew()
	o.StartTime = time.Now()
	return o
}

func (o *Span) initContext(ctx context.Context) {
	o.Ctx = context.WithValue(ctx, traces.ContextKeyForTrace, o)
}
