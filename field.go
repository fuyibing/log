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
	// Field
	// 自定义字段.
	//
	//   log.Field{"key":"value"}.
	//       Debug("message")
	Field map[string]interface{}
)

func (o Field) Debug(text string, args ...interface{}) { sendLog(o, common.Debug, text, args...) }
func (o Field) Info(text string, args ...interface{})  { sendLog(o, common.Info, text, args...) }
func (o Field) Warn(text string, args ...interface{})  { sendLog(o, common.Warn, text, args...) }
func (o Field) Error(text string, args ...interface{}) { sendLog(o, common.Error, text, args...) }
func (o Field) Fatal(text string, args ...interface{}) { sendLog(o, common.Fatal, text, args...) }

// sendLog
// 发送日志.
func sendLog(field Field, level common.Level, text string, args ...interface{}) {
	var kv loggers.Kv

	// 复制 Key/Value.
	if len(field) > 0 {
		kv = loggers.Kv{}
		for k, v := range field {
			kv[k] = v
		}
	}

	// 发送 Log.
	Manager.Logger().Push(kv, level, text, args...)
}
