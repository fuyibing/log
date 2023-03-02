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

package traces

import (
	"strings"
)

type (
	// Level 日志级别.
	Level string
)

// 日志级别枚举.

const (
	Off   Level = "OFF"
	Debug Level = "DEBUG"
	Info  Level = "INFO"
	Warn  Level = "WARN"
	Error Level = "ERROR"
	Fatal Level = "FATAL"

	DefaultLevel = Info
)

var (
	LevelToIntegers = map[Level]int{
		Off:   1,
		Fatal: 2,
		Error: 3,
		Warn:  4,
		Info:  5,
		Debug: 6,
	}
)

func (o Level) Int() int {
	if n, ok := LevelToIntegers[o]; ok {
		return n
	}
	return 0
}

func (o Level) Upper() Level {
	return Level(strings.ToUpper(string(o)))
}