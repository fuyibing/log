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

package configurer

import (
	"github.com/fuyibing/log/v5/common"
)

type ConfigLogger interface {
	GetLoggerExporter() string
	GetLoggerLevel() common.Level
	LevelEnabled(level common.Level) bool
}

// Getter

func (o *config) LevelEnabled(level common.Level) (yes bool) {
	switch level {
	case common.Debug:
		yes = o.debugOn
	case common.Info:
		yes = o.infoOn
	case common.Warn:
		yes = o.warnOn
	case common.Error:
		yes = o.errorOn
	case common.Fatal:
		yes = o.fatalOn
	}
	return
}

func (o *config) GetLoggerExporter() string    { return o.LoggerExporter }
func (o *config) GetLoggerLevel() common.Level { return o.LoggerLevel }

// Setter

func (o *Setter) SetLoggerExporter(v string) *Setter { o.config.LoggerExporter = v; return o }

func (o *Setter) SetLoggerLevel(v common.Level) *Setter {
	o.config.LoggerLevel = v
	o.config.state()
	return o
}

// Access

func (o *config) defaultLogger() {
	if o.LoggerExporter == "" {
		o.LoggerExporter = defaultLoggerExporter
	}

	if o.LoggerLevel.Upper().Int() == 0 {
		o.LoggerLevel = defaultLoggerLevel
	} else {
		o.LoggerLevel = o.LoggerLevel.Upper()
	}

	o.state()
}

func (o *config) state() {
	li := o.LoggerLevel.Int()
	yes := li > common.Off.Int()

	o.debugOn = yes && li >= common.Debug.Int()
	o.errorOn = yes && li >= common.Error.Int()
	o.fatalOn = yes && li >= common.Fatal.Int()
	o.infoOn = yes && li >= common.Info.Int()
	o.warnOn = yes && li >= common.Warn.Int()
}
