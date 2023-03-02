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
		DebugOn() bool
		ErrorOn() bool
		FatalOn() bool
		GetLoggerExporter() string
		GetLoggerLevel() traces.Level
		InfoOn() bool
		WarnOn() bool
	}
)

// Getter

func (o *config) DebugOn() bool                { return o.debugOn }
func (o *config) ErrorOn() bool                { return o.errorOn }
func (o *config) FatalOn() bool                { return o.fatalOn }
func (o *config) GetLoggerExporter() string    { return o.LoggerExporter }
func (o *config) GetLoggerLevel() traces.Level { return o.LoggerLevel }
func (o *config) InfoOn() bool                 { return o.infoOn }
func (o *config) WarnOn() bool                 { return o.warnOn }

// Setter

func (o *FieldManager) SetLoggerExporter(s string) *FieldManager {
	o.config.LoggerExporter = s
	return o
}

func (o *FieldManager) SetLoggerLevel(l traces.Level) *FieldManager {
	o.config.LoggerLevel = l
	o.config.updateState()
	return o
}

// Initialize

func (o *config) initLoggerDefaults() {
	// 默认级别.
	if o.LoggerLevel.Upper().Int() == 0 {
		o.LoggerLevel = traces.DefaultLevel
	}

	// 更新状态.
	o.updateState()
}

func (o *config) updateState() {
	var (
		n  = o.LoggerLevel.Int()
		ni = n > traces.Off.Int()
	)

	o.debugOn = ni && n >= traces.Debug.Int()
	o.infoOn = ni && n >= traces.Info.Int()
	o.warnOn = ni && n >= traces.Warn.Int()
	o.errorOn = ni && n >= traces.Error.Int()
	o.fatalOn = ni && n >= traces.Fatal.Int()
}
