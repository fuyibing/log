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
// date: 2023-02-25

package conf

type (
	JaegerTracerConfiguration interface {
		GetContentType() string
		GetEndpoint() string
		GetPassword() string
		GetUsername() string
	}

	jaegerTracerConfiguration struct {
		ContentType string `yaml:"content-type"`
		Endpoint    string `yaml:"endpoint"`
		Password    string `yaml:"password"`
		Username    string `yaml:"username"`
	}
)

func (o *jaegerTracerConfiguration) GetContentType() string { return o.ContentType }
func (o *jaegerTracerConfiguration) GetEndpoint() string    { return o.Endpoint }
func (o *jaegerTracerConfiguration) GetPassword() string    { return o.Password }
func (o *jaegerTracerConfiguration) GetUsername() string    { return o.Username }

// /////////////////////////////////////////////////////////////////////////////
// Set option
// /////////////////////////////////////////////////////////////////////////////

func JaegerTracerContentType(s string) Option {
	return func(c *configuration) { c.JaegerTracer.ContentType = s }
}

func JaegerTracerEndpoint(s string) Option {
	return func(c *configuration) { c.JaegerTracer.Endpoint = s }
}

func JaegerTracerPassword(s string) Option {
	return func(c *configuration) { c.JaegerTracer.Password = s }
}

func JaegerTracerUsername(s string) Option {
	return func(c *configuration) { c.JaegerTracer.Username = s }
}

// /////////////////////////////////////////////////////////////////////////////
// Jaeger Trace Configuration: access and constructor
// /////////////////////////////////////////////////////////////////////////////

func (o *jaegerTracerConfiguration) initDefaults() {
}
