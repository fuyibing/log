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
	"time"
)

var (
	nilSpanId SpanId
)

type (
	// Span 链路跨度组件.
	Span interface {
		// Child 生成子跨度组件.
		Child(name string) Span

		// Context 获取上下文.
		Context() context.Context

		// End 结束跨度.
		End()

		// GetAttribute 获取组件属性.
		GetAttribute() Attribute

		// GetDuration 获取跨度时长.
		GetDuration() time.Duration

		// GetLogs 跨度日志列表.
		GetLogs() []Log

		// GetName 跨度名称.
		GetName() string

		// GetSpanId 获取跨度ID.
		GetSpanId() SpanId

		// GetParentSpanId 获取上级跨度ID.
		GetParentSpanId() SpanId

		// GetStartTime 获取开始时间.
		GetStartTime() time.Time

		// GetTrace 获取跟踪组件.
		GetTrace() Trace

		// Logger 跨度日志.
		Logger() SpanLogger
	}

	// SpanId 链路跨度ID.
	SpanId [8]byte

	// SpanLogger 跨度日志.
	//
	// 基于链路跨度设置用户日志.
	SpanLogger interface {
		// Add 添加 Key/Value 属性.
		Add(key string, value interface{}) SpanLogger

		// Debug 记录DEBUG级日志.
		Debug(text string, args ...interface{})

		// Error 记录ERROR级日志.
		Error(text string, args ...interface{})

		// Fatal 记录FATAL级日志.
		Fatal(text string, args ...interface{})

		// Info 记录INFO级日志.
		Info(text string, args ...interface{})

		// Warn 记录WARN级日志.
		Warn(text string, args ...interface{})
	}
)

// IsValid 合法状态.
func (o SpanId) IsValid() bool { return !bytes.Equal(o[:], nilSpanId[:]) }

// String 生成字符串.
func (o SpanId) String() string { return hex.EncodeToString(o[:]) }
