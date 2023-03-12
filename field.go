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
	// key/value pair on each log.
	//
	//   log.Field{"key":"value"}.
	//       Debug("message")
	Field map[string]interface{}
)

// Debug
// send DEBUG level log to executor.
func (o Field) Debug(format string, args ...interface{}) {
	sendLog(o, common.Debug, format, args...)
}

// Error
// send ERROR level log to executor.
func (o Field) Error(format string, args ...interface{}) {
	sendLog(o, common.Error, format, args...)
}

// Fatal
// send FATAL level log to executor.
func (o Field) Fatal(format string, args ...interface{}) {
	sendLog(o, common.Fatal, format, args...)
}

// Info
// send INFO level log to executor.
func (o Field) Info(format string, args ...interface{}) {
	sendLog(o, common.Info, format, args...)
}

// Warn
// send WARN level log to executor.
func (o Field) Warn(format string, args ...interface{}) {
	sendLog(o, common.Warn, format, args...)
}

func sendLog(field Field, level common.Level, format string, args ...interface{}) {
	var kv loggers.Kv

	// Copy Key/Value pairs into log component.
	if len(field) > 0 {
		kv = loggers.Kv{}
		for k, v := range field {
			kv[k] = v
		}
	}

	// Send to executor by manager dispatcher.
	Manager.Logger().Push(kv, level, format, args...)
}
