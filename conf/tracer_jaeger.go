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
	// JaegerTracerConfiguration
	// 适配 Jaeger 配置参数.
	JaegerTracerConfiguration interface {
		// GetContentType
		// 请求格式.
		GetContentType() string

		// GetEndpoint
		// 请求地址.
		GetEndpoint() string

		// GetPassword
		// 鉴权密码.
		GetPassword() string

		// GetUsername
		// 鉴权用户名.
		GetUsername() string
	}

	jaegerTracerConfiguration struct {
		ContentType string `yaml:"content-type"`
		Endpoint    string `yaml:"endpoint"`
		Password    string `yaml:"password"`
		Username    string `yaml:"username"`
	}
)

// GetJaegerTracer
// 获取 Jaeger 配置集合.
func (o *configuration) GetJaegerTracer() JaegerTracerConfiguration { return o.JaegerTracer }

// GetContentType
// 请求格式.
func (o *jaegerTracerConfiguration) GetContentType() string { return o.ContentType }

// GetEndpoint
// 请求地址.
func (o *jaegerTracerConfiguration) GetEndpoint() string { return o.Endpoint }

// GetPassword
// 鉴权密码.
func (o *jaegerTracerConfiguration) GetPassword() string { return o.Password }

// GetUsername
// 鉴权用户名.
func (o *jaegerTracerConfiguration) GetUsername() string { return o.Username }

// initDefaults
// 构造配置.
func (o *jaegerTracerConfiguration) initDefaults() {}

// JaegerTracerContentType
// 设置请求格式.
//
// 默认: application/x-thrift
func JaegerTracerContentType(s string) Option {
	return func(c *configuration) { c.JaegerTracer.ContentType = s }
}

// JaegerTracerEndpoint
// 设置请求地址.
//
// 例如: http://localhost:14268/api/traces
func JaegerTracerEndpoint(s string) Option {
	return func(c *configuration) { c.JaegerTracer.Endpoint = s }
}

// JaegerTracerPassword
// 设置鉴权密码.
func JaegerTracerPassword(s string) Option {
	return func(c *configuration) { c.JaegerTracer.Password = s }
}

// JaegerTracerUsername
// 设置鉴权用户名.
func JaegerTracerUsername(s string) Option {
	return func(c *configuration) { c.JaegerTracer.Username = s }
}
