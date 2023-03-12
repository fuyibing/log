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
// date: 2023-03-03

package log

import (
	"github.com/fuyibing/log/v5/common"
)

// Debug
// send DEBUG level log to executor.
func Debug(format string, args ...interface{}) {
	Manager.Logger().Push(nil, common.Debug, format, args...)
}

// Error
// send ERROR level log to executor.
func Error(format string, args ...interface{}) {
	Manager.Logger().Push(nil, common.Error, format, args...)
}

// Fatal
// send FATAL level log to executor.
func Fatal(format string, args ...interface{}) {
	Manager.Logger().Push(nil, common.Fatal, format, args...)
}

// Info
// send INFO level log to executor.
func Info(format string, args ...interface{}) {
	Manager.Logger().Push(nil, common.Info, format, args...)
}

// Warn
// send WARN level log to executor.
func Warn(format string, args ...interface{}) {
	Manager.Logger().Push(nil, common.Warn, format, args...)
}
