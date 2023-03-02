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
// date: 2023-03-01

package conf

import (
	"github.com/fuyibing/log/v5/traces"
)

type (
	// ConfigLogger
	// 日志配置.
	ConfigLogger interface {
		// GetLoggerExporter
		// 日志上报名称.
		//
		// - term
		// - file
		// - kafka
		// - aliyunsls
		GetLoggerExporter() string

		// GetLoggerLevel
		// 日志记录级别.
		//
		// - DEBUG
		// - INFO
		// - WARN
		// - ERROR
		// - FATAL
		GetLoggerLevel() traces.Level

		DebugOn() bool
		ErrorOn() bool
		FatalOn() bool
		InfoOn() bool
		WarnOn() bool
	}
)

// GetLoggerExporter
// 日志上报名称.
func (o *config) GetLoggerExporter() string { return o.LoggerExporter }

// GetLoggerLevel
// 日志记录级别.
func (o *config) GetLoggerLevel() traces.Level { return o.LoggerLevel }

func (o *config) DebugOn() bool { return true }
func (o *config) ErrorOn() bool { return true }
func (o *config) FatalOn() bool { return true }
func (o *config) InfoOn() bool  { return true }
func (o *config) WarnOn() bool  { return true }

func (o *config) initLoggerDefaults() bool {
	return true
}
