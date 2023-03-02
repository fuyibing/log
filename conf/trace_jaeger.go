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

	jaegerTracer struct {
		// 上报格式.
		// 默认: application/x-thrift
		ContentType string `yaml:"content-type"`

		// 上报位置.
		// 例如: http://localhost:14268/api/traces
		Endpoint string `yaml:"endpoint"`
		Password string `yaml:"password"`
		Username string `yaml:"username"`
	}
)

// Child.

func (o *config) GetJaegerTracer() ConfigJaegerTracer { return o.JaegerTracer }

// Getter

func (o *jaegerTracer) GetContentType() string { return o.ContentType }
func (o *jaegerTracer) GetEndpoint() string    { return o.Endpoint }
func (o *jaegerTracer) GetPassword() string    { return o.Password }
func (o *jaegerTracer) GetUsername() string    { return o.Username }

// Setter

func (o *FieldManager) SetJaegerTracerContentType(s string) *FieldManager {
	if s != "" {
		o.config.JaegerTracer.ContentType = s
	}
	return o
}

func (o *FieldManager) SetJaegerTracerEndpoint(s string) *FieldManager {
	if s != "" {
		o.config.JaegerTracer.Endpoint = s
	}
	return o
}

func (o *FieldManager) SetJaegerTracerPassword(s string) *FieldManager {
	if s != "" {
		o.config.JaegerTracer.Password = s
	}
	return o
}

func (o *FieldManager) SetJaegerTracerUsername(s string) *FieldManager {
	if s != "" {
		o.config.JaegerTracer.Username = s
	}
	return o
}

// Constructor

func (o *jaegerTracer) initDefaults() {
	if o.ContentType == "" {
		o.ContentType = "application/x-thrift"
	}
}
