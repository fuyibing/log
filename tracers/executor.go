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
// date: 2023-03-04

package tracers

import (
	"github.com/fuyibing/util/v8/process"
)

type (
	// Executor
	// for tracer.
	Executor interface {
		// Processor
		// return executor processor like os process.
		Processor() (processor process.Processor)

		// Publish
		// span component into executor.
		Publish(spans ...Span) (err error)

		// SetFormatter
		// register tracer formatter handler.
		SetFormatter(formatter Formatter)
	}
)
