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

package base

// Level
// 日志级别.
type Level string

// 预定义日志级别枚举值.

const (
	Off   Level = "OFF"
	Debug Level = "DEBUG"
	Info  Level = "INFO"
	Warn  Level = "WARN"
	Error Level = "ERROR"
	Fatal Level = "FATAL"
)

// 预定义级别与整型映射.

var (
	levelTextToInteger = map[Level]int{
		Off:   1,
		Fatal: 2,
		Error: 3,
		Warn:  4,
		Info:  5,
		Debug: 6,
	}
)

// Int
// 日志级别转成整型.
func (o Level) Int() int {
	if i, ok := levelTextToInteger[o]; ok {
		return i
	}
	return 0
}

// String
// 日志级别转成字符串.
func (o Level) String() string {
	return string(o)
}
