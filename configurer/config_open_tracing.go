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

type (
	// ConfigOpenTracing
	// expose opentracing configuration methods.
	ConfigOpenTracing interface {
		GetOpenTracingSampled() string
		GetOpenTracingSpanId() string
		GetOpenTracingTraceId() string
	}
)

// Getter

func (o *config) GetOpenTracingSampled() string { return o.OpenTracingSampled }
func (o *config) GetOpenTracingSpanId() string  { return o.OpenTracingSpanId }
func (o *config) GetOpenTracingTraceId() string { return o.OpenTracingTraceId }

// Setter

func (o *Setter) SetOpenTracingSampled(s string) *Setter { o.config.OpenTracingSampled = s; return o }
func (o *Setter) SetOpenTracingSpanId(s string) *Setter  { o.config.OpenTracingSpanId = s; return o }
func (o *Setter) SetOpenTracingTraceId(s string) *Setter { o.config.OpenTracingTraceId = s; return o }

// Access

func (o *config) defaultOpenTracing() {
	if o.OpenTracingSampled == "" {
		o.OpenTracingSampled = defaultOpenTracingSampled
	}
	if o.OpenTracingSpanId == "" {
		o.OpenTracingSpanId = defaultOpenTracingSpanId
	}
	if o.OpenTracingTraceId == "" {
		o.OpenTracingTraceId = defaultOpenTracingTraceId
	}
}
