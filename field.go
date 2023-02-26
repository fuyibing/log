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

type (
	// Field additional key/value pairs component for log.
	//
	// Attention: do not use it in different goroutine, it is not security,
	// it would fire a panic error.
	//
	//   log.Field{}.
	//       Add("key", "value").
	//       Debug("message content: %d", 1)
	Field struct {
		data map[string]interface{}
	}
	// map[string]interface{}
)

// Add
// key/value pair on component.
func (f Field) Add(key string, value interface{}) Field {
	if f.data == nil {
		f.data = make(map[string]interface{})
	}

	f.data[key] = value
	return f
}

// Debug
// send key/value component to log exporter with debug level.
func (f Field) Debug(text string, args ...interface{}) {
	if conf.Config.DebugOn() {
		cores.Registry.LoggerPush(f.data, base.Debug, text, args...)
	}
}

// Error
// send key/value component to log exporter with error level.
func (f Field) Error(text string, args ...interface{}) {
	if conf.Config.ErrorOn() {
		cores.Registry.LoggerPush(f.data, base.Error, text, args...)
	}
}

// Fatal
// send key/value component to log exporter with fatal level.
func (f Field) Fatal(text string, args ...interface{}) {
	if conf.Config.FatalOn() {
		cores.Registry.LoggerPush(f.data, base.Fatal, text, args...)
	}
}

// Info
// send key/value component to log exporter with info level.
func (f Field) Info(text string, args ...interface{}) {
	if conf.Config.InfoOn() {
		cores.Registry.LoggerPush(f.data, base.Info, text, args...)
	}
}

// Warn
// send key/value component to log exporter with warning level.
func (f Field) Warn(text string, args ...interface{}) {
	if conf.Config.WarnOn() {
		cores.Registry.LoggerPush(f.data, base.Warn, text, args...)
	}
}
