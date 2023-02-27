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

package built_in

import (
	"github.com/fuyibing/log/v5/cores"
	"github.com/fuyibing/log/v5/exporters/tracer_jaeger"
	"github.com/fuyibing/log/v5/exporters/tracer_term"
)

type (
	// BuiltinTracer
	// 内置Tracer类型.
	BuiltinTracer string
)

// 内置日志枚举.

const (
	TracerTerm   BuiltinTracer = "term"
	TracerJaeger BuiltinTracer = "jaeger"
)

// Exporter
// 基于内建 Tracer 类型, 获取其实例.
func (b BuiltinTracer) Exporter() cores.TracerExporter {
	switch b {
	case TracerJaeger:
		return tracer_jaeger.NewExporter()
	case TracerTerm:
		return tracer_term.NewExporter()
	}
	return nil
}
