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
	// Field
	// 应用于日志输出的键/值对附加信息.
	//
	// 此支持非必须, 便于输出易读的格式化数据. 当日志输出选择在中间件(如:aliyun sls)中时
	// 可将此结构标准化成易查询、易索引的数据.
	Field struct {
		data map[string]interface{}
	}
)

// Add
// 添加 K/V 键值对.
func (f Field) Add(key string, value interface{}) Field {
	if f.data == nil {
		f.data = make(map[string]interface{})
	}

	f.data[key] = value
	return f
}

// Debug
// 记录为 debug 级日志.
func (f Field) Debug(text string, args ...interface{}) {
	if conf.Config.DebugOn() {
		cores.Registry.LoggerPush(f.data, base.Debug, text, args...)
	}
}

// Error
// 记录为 error 级日志.
func (f Field) Error(text string, args ...interface{}) {
	if conf.Config.ErrorOn() {
		cores.Registry.LoggerPush(f.data, base.Error, text, args...)
	}
}

// Fatal
// 记录为 fatal 级日志.
func (f Field) Fatal(text string, args ...interface{}) {
	if conf.Config.FatalOn() {
		cores.Registry.LoggerPush(f.data, base.Fatal, text, args...)
	}
}

// Info
// 记录为 info 级日志.
func (f Field) Info(text string, args ...interface{}) {
	if conf.Config.InfoOn() {
		cores.Registry.LoggerPush(f.data, base.Info, text, args...)
	}
}

// Warn
// 记录为 warning 级日志.
func (f Field) Warn(text string, args ...interface{}) {
	if conf.Config.WarnOn() {
		cores.Registry.LoggerPush(f.data, base.Warn, text, args...)
	}
}
