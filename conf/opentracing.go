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

package conf

type (
	InterfaceOpentracing interface {
		GetOpenTracingSample() string
		GetOpenTracingSpanId() string
		GetOpenTracingTraceId() string
	}
)

func (o *configuration) GetOpenTracingSample() string  { return o.OpenTracingSample }
func (o *configuration) GetOpenTracingSpanId() string  { return o.OpenTracingSpanId }
func (o *configuration) GetOpenTracingTraceId() string { return o.OpenTracingTraceId }

func OpenTracingSample(s string) Option {
	return func(c *configuration) { c.OpenTracingSample = s }
}

func OpenOpenTracingSpanId(s string) Option {
	return func(c *configuration) { c.OpenTracingSpanId = s }
}

func OpenTracingTraceId(s string) Option {
	return func(c *configuration) { c.OpenTracingTraceId = s }
}
