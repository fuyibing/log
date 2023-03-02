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

package traces

import (
	"bytes"
	"context"
	"encoding/hex"
	"github.com/fuyibing/util/v8/process"
)

var (
	nilTraceId TraceId
)

type (
	// Trace 跟踪组件.
	Trace interface {
		// Context 获取上下文.
		Context() context.Context

		// GetAttribute 获取组件属性.
		GetAttribute() Attribute

		// GetName 获取跟踪名称.
		GetName() string

		// GetSpanId 获取跨度ID.
		GetSpanId() SpanId

		// GetTraceId
		// 获取跟踪ID.
		GetTraceId() TraceId

		// Span 派生跨度组件.
		//
		// 从跟踪组件上派生一个跨度组件.
		Span(name string) Span
	}

	// TraceId 跟踪ID.
	TraceId [16]byte

	// TracerExporter 跨度导出.
	TracerExporter interface {
		// Processor 获取类进程.
		Processor() process.Processor

		// Send 发送跨度.
		Send(span Span) error

		// SetFormatter 设置格式.
		SetFormatter(formatter TracerFormatter)
	}

	// TracerFormatter 跟踪格式化.
	TracerFormatter interface {
		// Byte 转成字符码.
		Byte(_ Span) []byte

		// String 转成字符串.
		String(v Span) (text string)
	}
)

// IsValid 合法状态.
func (o TraceId) IsValid() bool { return !bytes.Equal(o[:], nilTraceId[:]) }

// String 跟踪ID字符串.
func (o TraceId) String() string { return hex.EncodeToString(o[:]) }
