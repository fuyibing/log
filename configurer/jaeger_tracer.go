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

package configurer

type (
	ConfigTracerJaeger interface {
		GetJaegerTracer() JaegerTracer
	}

	JaegerTracer interface {
		GetContentType() string
		GetEndpoint() string
		GetPassword() string
		GetUsername() string
	}

	jaegerTracer struct {
		// API Content type.
		// Default: application/x-thrift
		ContentType string `yaml:"content-type"`

		// API Address
		// Example: http://localhost:14268/api/traces
		Endpoint string `yaml:"endpoint"`

		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
)

// Getter

func (o *config) GetJaegerTracer() JaegerTracer { return o.JaegerTracer }

func (o *jaegerTracer) GetContentType() string { return o.ContentType }
func (o *jaegerTracer) GetEndpoint() string    { return o.Endpoint }
func (o *jaegerTracer) GetPassword() string    { return o.Password }
func (o *jaegerTracer) GetUsername() string    { return o.Username }

// Setter.

func (o *Setter) SetJaegerTracerContentType(s string) *Setter {
	o.config.JaegerTracer.ContentType = s
	return o
}

func (o *Setter) SetJaegerTracerEndpoint(s string) *Setter {
	o.config.JaegerTracer.Endpoint = s
	return o
}

func (o *Setter) SetJaegerTracerPassword(s string) *Setter {
	o.config.JaegerTracer.Password = s
	return o
}

func (o *Setter) SetJaegerTracerUsername(s string) *Setter {
	o.config.JaegerTracer.Username = s
	return o
}

// Access.

func (o *jaegerTracer) initDefaults() {
	if o.ContentType == "" {
		o.ContentType = "application/x-thrift"
	}
}
