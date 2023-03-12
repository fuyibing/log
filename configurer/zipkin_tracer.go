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
	ConfigTracerZipkin interface {
		GetZipkinTracer() ZipkinTracer
	}

	ZipkinTracer interface {
		GetContentType() string
		GetEndpoint() string
	}

	zipkinTracer struct {
		// API Content type.
		// Default: application/json
		ContentType string `yaml:"content-type"`

		// API Address.
		// Example: http://localhost:9411/api/v2/spans
		Endpoint string `yaml:"endpoint"`
	}
)

// Getter

func (o *config) GetZipkinTracer() ZipkinTracer { return o.ZipkinTracer }

func (o *zipkinTracer) GetContentType() string { return o.ContentType }
func (o *zipkinTracer) GetEndpoint() string    { return o.Endpoint }

// Setter.

func (o *Setter) SetZipkinTracerContentType(s string) *Setter {
	o.config.ZipkinTracer.ContentType = s
	return o
}

func (o *Setter) SetZipkinTracerEndpoint(s string) *Setter {
	o.config.ZipkinTracer.Endpoint = s
	return o
}

// Access.

func (o *zipkinTracer) initDefaults() {
	if o.ContentType == "" {
		o.ContentType = "application/json"
	}
}
