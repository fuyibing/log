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
	"github.com/fuyibing/log/v5/base"
)

type (
	// ConfigTracer
	// 链路配置.
	ConfigTracer interface {
		// GetTracerExporter
		// 链路上报名称.
		//
		// - term
		// - file
		// - jaeger
		// - zipkin
		GetTracerExporter() string

		// GetTracerTopic
		// 链路上报位置.
		GetTracerTopic() string
	}
)

// GetTracerExporter
// 链路上报名称.
func (o *config) GetTracerExporter() string { return o.TracerExporter }

// GetTracerTopic
// 链路上报位置.
func (o *config) GetTracerTopic() string { return o.TracerTopic }

func (o *config) initTracerDefaults() {
	if o.TracerTopic == "" {
		o.TracerTopic = base.TracerTopic
	}
}
