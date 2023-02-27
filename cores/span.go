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
	"context"
	"github.com/fuyibing/log/v5/base"
	"sync"
)

type (
	// Span
	// 用于Trace的跨度组件.
	Span interface {
		// AddEvent
		// 添加后置事件.
		AddEvent(events ...SpanEvent) Span

		// Child
		// 创建子Span组件.
		Child(name string) Span

		// End
		// 结束Span.
		End()

		// GetAttr
		// 获取Span的属性组件.
		GetAttr() Attr

		// GetContext
		// 获取Span的上下文.
		GetContext() context.Context

		// GetLogs
		// 获取Span的日志组件.
		GetLogs() SpanLogs

		// GetName
		// 获取Span组件名称.
		GetName() string

		// GetSpanId
		// 获取Span的ID组件.
		GetSpanId() SpanId

		// GetParentSpanId
		// 获取Span的上游ID组件.
		GetParentSpanId() SpanId

		// GetSpanTime
		// 获取Span的时间组件.
		GetSpanTime() SpanTime

		// GetTrace
		// 获取Span的隶属Trace组件.
		GetTrace() Trace

		// GetTraceId
		// 获取Span的TraceId组件.
		GetTraceId() TraceId

		// Logger
		// 获取Span的Log组件.
		Logger() *SpanLogger
	}

	span struct {
		sync.RWMutex

		attr                 Attr
		ctx                  context.Context
		events               []SpanEvent
		name                 string
		spanId, parentSpanId SpanId
		spanLogs             SpanLogs
		spanTime             SpanTime
		trace                Trace
		traceId              TraceId
	}
)

// AddEvent
// 添加后置事件.
func (o *span) AddEvent(events ...SpanEvent) Span {
	o.Lock()
	defer o.Unlock()

	o.events = append(o.events, events...)
	return o
}

// Child
// 创建子Span组件.
func (o *span) Child(name string) Span {
	return (&span{name: name}).
		init().
		initRelations(o.trace, o.traceId, o.spanId).
		initContext(o.ctx)
}

// End
// 结束Span.
func (o *span) End() {
	o.spanTime.End()

	// 推送Span到Exporter.
	if Registry.TracerEnabled() {
		Registry.TracerExporter().Push(o)
	}

	// 执行已注册后置事件.
	for _, event := range o.events {
		event.Do(o)
	}
}

// GetAttr
// 获取Span的属性组件.
func (o *span) GetAttr() Attr { return o.attr }

// GetContext
// 获取Span的上下文.
func (o *span) GetContext() context.Context { return o.ctx }

// GetLogs
// 获取Span的日志组件.
func (o *span) GetLogs() SpanLogs {
	o.RLock()
	defer o.RUnlock()

	return o.spanLogs
}

// GetName
// 获取Span组件名称.
func (o *span) GetName() string { return o.name }

// GetSpanId
// 获取Span的ID组件.
func (o *span) GetSpanId() SpanId { return o.spanId }

// GetParentSpanId
// 获取Span的上游ID组件.
func (o *span) GetParentSpanId() SpanId { return o.parentSpanId }

// GetSpanTime
// 获取Span的时间组件.
func (o *span) GetSpanTime() SpanTime { return o.spanTime }

// GetTrace
// 获取Span的隶属Trace组件.
func (o *span) GetTrace() Trace { return o.trace }

// GetTraceId
// 获取Span的TraceId组件.
func (o *span) GetTraceId() TraceId { return o.traceId }

// Logger
// 获取Span的Log组件.
func (o *span) Logger() *SpanLogger { return NewSpanLogger(o) }

// init
// 初始化 Span 跨度.
func (o *span) init() *span {
	o.attr = NewAttr()
	o.events = make([]SpanEvent, 0)
	o.spanId = Identify.GenSpanId()
	o.spanLogs = NewSpanLogs()
	o.spanTime = NewSpanTime()
	return o
}

// initContext
// 初始化 Span 上下文.
func (o *span) initContext(ctx context.Context) *span {
	// 默认上下文.
	if ctx == nil {
		ctx = context.Background()
	}

	// 向 context.Context 绑定 Span 跨度.
	o.ctx = context.WithValue(ctx, base.ContextKeySpan, o)
	return o
}

// initRelations
// 初始化 Span 关系, 绑定 Trace, TraceId, SpanId 参数.
func (o *span) initRelations(t Trace, tid TraceId, pid SpanId) *span {
	o.trace = t
	o.traceId = tid
	o.parentSpanId = pid
	return o
}
