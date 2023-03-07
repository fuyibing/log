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

// Debug 记录 DEBUG 级日志.
func Debug(format string, args ...interface{}) {
	Manager.Logger().Push(nil, common.Debug, format, args...)
}

// Error 记录 ERROR 级日志.
func Error(format string, args ...interface{}) {
	Manager.Logger().Push(nil, common.Error, format, args...)
}

// Fatal 记录 FATAL 级日志.
func Fatal(format string, args ...interface{}) {
	Manager.Logger().Push(nil, common.Fatal, format, args...)
}

// Info 记录 INFO 级日志.
func Info(format string, args ...interface{}) {
	Manager.Logger().Push(nil, common.Info, format, args...)
}

// Warn 记录 WARN 级日志.
func Warn(format string, args ...interface{}) {
	Manager.Logger().Push(nil, common.Warn, format, args...)
}
