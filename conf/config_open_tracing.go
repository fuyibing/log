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

type ConfigOpenTracing interface {
	GetOpenTracingSampled() string
	GetOpenTracingSpanId() string
	GetOpenTracingTraceId() string
}

func (o *config) GetOpenTracingSampled() string { return o.OpenTracingSampled }
func (o *config) GetOpenTracingSpanId() string  { return o.OpenTracingSpanId }
func (o *config) GetOpenTracingTraceId() string { return o.OpenTracingTraceId }

func (o *config) initDefaultsOpenTracing() {
	if o.OpenTracingTraceId == "" {
		o.OpenTracingTraceId = traces.OpenTracingTraceId
	}
	if o.OpenTracingSpanId == "" {
		o.OpenTracingSpanId = traces.OpenTracingSpanId
	}
	if o.OpenTracingSampled == "" {
		o.OpenTracingSampled = traces.OpenTracingSampled
	}
}
