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

type (
	JaegerTracer interface {
		GetJaegerTracer() ConfigJaegerTracer
	}

	ConfigJaegerTracer interface {
		GetContentType() string
		GetEndpoint() string
		GetPassword() string
		GetUsername() string
	}

	jaegerTracer struct{}
)

func (o *config) GetJaegerTracer() ConfigJaegerTracer { return o.JaegerTracer }

func (o *jaegerTracer) GetContentType() string { return "application/x-thrift" }
func (o *jaegerTracer) GetEndpoint() string    { return "http://localhost:14268/api/traces" }
func (o *jaegerTracer) GetPassword() string    { return "" }
func (o *jaegerTracer) GetUsername() string    { return "" }

func (o *jaegerTracer) initDefaults() {}
