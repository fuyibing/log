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
// date: 2023-02-27

package conf

import (
	"github.com/fuyibing/log/v5/base"
	"strings"
)

type (
	InterfaceLogger interface {
		GetLoggerExporter() string
		GetLoggerLevel() base.Level

		DebugOn() bool
		ErrorOn() bool
		FatalOn() bool
		InfoOn() bool
		WarnOn() bool
	}
)

func (o *configuration) DebugOn() bool              { return o.debugOn }
func (o *configuration) ErrorOn() bool              { return o.errorOn }
func (o *configuration) FatalOn() bool              { return o.fatalOn }
func (o *configuration) GetLoggerExporter() string  { return strings.ToLower(o.LoggerExporter) }
func (o *configuration) GetLoggerLevel() base.Level { return o.LoggerLevel }
func (o *configuration) InfoOn() bool               { return o.infoOn }
func (o *configuration) WarnOn() bool               { return o.warnOn }

func LogExporter(s string) Option {
	return func(c *configuration) { c.LoggerExporter = s }
}

func LogLevel(l base.Level) Option {
	return func(c *configuration) { c.LoggerLevel = l; c.updateLevelState() }
}
