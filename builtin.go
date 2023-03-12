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
// date: 2023-03-06

package log

import (
	"github.com/fuyibing/log/v5/loggers"
	"github.com/fuyibing/log/v5/loggers/logger_file"
	"github.com/fuyibing/log/v5/loggers/logger_term"
	"github.com/fuyibing/log/v5/tracers"
	"github.com/fuyibing/log/v5/tracers/tracer_file"
	"github.com/fuyibing/log/v5/tracers/tracer_jaeger"
	"github.com/fuyibing/log/v5/tracers/tracer_term"
	"github.com/fuyibing/log/v5/tracers/tracer_zipkin"
)

var (
	// builtinLoggers
	// builtin executors for logger export.
	builtinLoggers = map[string]func() loggers.Executor{
		"file": logger_file.New,
		"term": logger_term.New,
	}

	// builtinTracers
	// builtin executors for tracer export.
	builtinTracers = map[string]func() tracers.Executor{
		"file":   tracer_file.New,
		"jaeger": tracer_jaeger.New,
		"term":   tracer_term.New,
		"zipkin": tracer_zipkin.New,
	}
)
