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
	"github.com/fuyibing/log/v5/conf"
	"github.com/fuyibing/log/v5/exporters"
	"github.com/fuyibing/log/v5/tracer"
	"github.com/fuyibing/log/v5/traces"
)

// Debug 记录DEBUG级日志.
func Debug(text string, args ...interface{}) {
	if conf.Config.DebugOn() {
		sendLogger(nil, traces.Debug, text, args...)
	}
}

// Error 记录ERROR级日志.
func Error(text string, args ...interface{}) {
	if conf.Config.ErrorOn() {
		sendLogger(nil, traces.Error, text, args...)
	}
}

// Fatal 记录FATAL级日志.
func Fatal(text string, args ...interface{}) {
	if conf.Config.FatalOn() {
		sendLogger(nil, traces.Fatal, text, args...)
	}
}

// Info 记录INFO级日志.
func Info(text string, args ...interface{}) {
	if conf.Config.InfoOn() {
		sendLogger(nil, traces.Info, text, args...)
	}
}

// Warn 记录WARN级日志.
func Warn(text string, args ...interface{}) {
	if conf.Config.WarnOn() {
		sendLogger(nil, traces.Warn, text, args...)
	}
}

// /////////////////////////////////////////////////////////////////////////////
// Access and constructor
// /////////////////////////////////////////////////////////////////////////////

func sendLogger(attr traces.Attribute, level traces.Level, text string, args ...interface{}) {
	x := tracer.NewLog(level, text, args...)
	x.Attribute.Copy(attr)
	exporters.Exporter.PutLogger(x)
}
