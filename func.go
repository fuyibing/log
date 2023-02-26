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
// send custom message to log exporter with debug level.
func Debug(text string, args ...interface{}) {
	if conf.Config.DebugOn() {
		cores.Registry.LoggerPush(nil, base.Debug, text, args...)
	}
}

// Error
// send custom message to log exporter with error level.
func Error(text string, args ...interface{}) {
	if conf.Config.ErrorOn() {
		cores.Registry.LoggerPush(nil, base.Error, text, args...)
	}
}

// Fatal
// send custom message to log exporter with fatal level.
func Fatal(text string, args ...interface{}) {
	if conf.Config.FatalOn() {
		cores.Registry.LoggerPush(nil, base.Fatal, text, args...)
	}
}

// Info
// send custom message to log exporter with info level.
func Info(text string, args ...interface{}) {
	if conf.Config.InfoOn() {
		cores.Registry.LoggerPush(nil, base.Info, text, args...)
	}
}

// Warn
// send custom message to log exporter with warning level.
func Warn(text string, args ...interface{}) {
	if conf.Config.WarnOn() {
		cores.Registry.LoggerPush(nil, base.Warn, text, args...)
	}
}
