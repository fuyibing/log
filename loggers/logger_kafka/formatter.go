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
// date: 2023-03-05

package logger_kafka

import (
	"github.com/fuyibing/log/v5/loggers"
)

type formatter struct {
}

func (o *formatter) Byte(vs ...loggers.Log) (body []byte, err error) {
	return
}

func (o *formatter) String(vs ...loggers.Log) (text string, err error) {
	return
}

func (o *formatter) init() *formatter {
	return o
}
