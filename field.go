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
	"github.com/fuyibing/log/v5/common"
	"github.com/fuyibing/log/v5/loggers"
)

type (
	// Field 字段组件, 允许在记录日志时附加自定义字段.
	//
	//   log.Field{"key":"value"}.
	//       Debug("message")
	Field map[string]interface{}
)

// Debug 记录 DEBUG 级日志.
func (o Field) Debug(format string, args ...interface{}) {
	sendLog(o, common.Debug, format, args...)
}

// Error 记录 ERROR 级日志.
func (o Field) Error(format string, args ...interface{}) {
	sendLog(o, common.Error, format, args...)
}

// Fatal 记录 FATAL 级日志.
func (o Field) Fatal(format string, args ...interface{}) {
	sendLog(o, common.Fatal, format, args...)
}

// Info 记录 INFO 级日志.
func (o Field) Info(format string, args ...interface{}) {
	sendLog(o, common.Info, format, args...)
}

// Warn 记录 WARN 级日志.
func (o Field) Warn(format string, args ...interface{}) {
	sendLog(o, common.Warn, format, args...)
}

func sendLog(field Field, level common.Level, format string, args ...interface{}) {
	var kv loggers.Kv

	// 复制
	// Key/Value.
	if len(field) > 0 {
		kv = loggers.Kv{}
		for k, v := range field {
			kv[k] = v
		}
	}

	// 发送
	// Log 到 操作台.
	Manager.Logger().Push(kv, level, format, args...)
}
