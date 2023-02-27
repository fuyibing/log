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

package log

import (
	"github.com/fuyibing/log/v5/base"
	"github.com/fuyibing/log/v5/conf"
	"github.com/fuyibing/log/v5/cores"
)

// Debug
// 记录为 debug 级日志.
func Debug(text string, args ...interface{}) {
	if conf.Config.DebugOn() {
		cores.Registry.LoggerPush(nil, base.Debug, text, args...)
	}
}

// Error
// 记录为 error 级日志.
func Error(text string, args ...interface{}) {
	if conf.Config.ErrorOn() {
		cores.Registry.LoggerPush(nil, base.Error, text, args...)
	}
}

// Fatal
// 记录为 fatal 级日志.
func Fatal(text string, args ...interface{}) {
	if conf.Config.FatalOn() {
		cores.Registry.LoggerPush(nil, base.Fatal, text, args...)
	}
}

// Info
// 记录为 info 级日志.
func Info(text string, args ...interface{}) {
	if conf.Config.InfoOn() {
		cores.Registry.LoggerPush(nil, base.Info, text, args...)
	}
}

// Warn
// 记录为 warning 级日志.
func Warn(text string, args ...interface{}) {
	if conf.Config.WarnOn() {
		cores.Registry.LoggerPush(nil, base.Warn, text, args...)
	}
}
