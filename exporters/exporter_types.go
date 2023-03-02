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

package exporters

import (
	"github.com/fuyibing/log/v5/base"
	"github.com/fuyibing/log/v5/exporters/logger_file"
	"github.com/fuyibing/log/v5/exporters/logger_term"
	"github.com/fuyibing/log/v5/exporters/tracer_jaeger"
	"github.com/fuyibing/log/v5/exporters/tracer_term"
)

type (
	// BuiltinLog
	// 内置日志类型.
	BuiltinLog string
)

// 内置日志枚举.

const (
	BuiltinLogTerm BuiltinLog = "term"
	BuiltinLogFile BuiltinLog = "file"
)

// 内置日志构造.

func (o BuiltinLog) Callable() (callable func() base.LoggerExporter) {
	switch o {
	case BuiltinLogTerm:
		callable = logger_term.New
	case BuiltinLogFile:
		callable = logger_file.New
	}
	return
}

type (
	// BuiltinSpan
	// 内置链路类型.
	BuiltinSpan string
)

// 内置链路枚举.

const (
	BuiltinSpanTerm   BuiltinSpan = "term"
	BuiltinSpanJaeger BuiltinSpan = "jaeger"
)

// 内置链路构造.

func (o BuiltinSpan) Callable() (callable func() base.TracerExporter) {
	switch o {
	case BuiltinSpanTerm:
		callable = tracer_term.New
	case BuiltinSpanJaeger:
		callable = tracer_jaeger.New
	}
	return
}
