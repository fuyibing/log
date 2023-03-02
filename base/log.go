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

package base

import (
	"github.com/fuyibing/util/v8/process"
	"time"
)

type (
	// Log 日志.
	Log interface {
		// GetAttribute
		// 日志属性.
		GetAttribute() Attribute

		// GetLevel
		// 日志级别.
		GetLevel() Level

		// GetTime
		// 日志时间.
		GetTime() time.Time

		// GetText
		// 日志正文.
		GetText() string
	}

	// LoggerExporter 日志导出.
	LoggerExporter interface {
		// Processor 获取类进程.
		Processor() process.Processor

		// Send 发送日志.
		Send(log Log) error

		// SetFormatter 设置格式化.
		SetFormatter(formatter LoggerFormatter)
	}

	// LoggerFormatter 日志格式化.
	LoggerFormatter interface {
		// Byte 转成Byte切片.
		Byte(v Log) []byte

		// String 转成字符串.
		String(v Log) string
	}
)
