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
// date: 2023-03-02

package conf

type ConfigService interface {
	GetServiceName() string
	GetServicePort() int
	GetServiceVersion() string
}

// Getter

func (o *config) GetServiceName() string    { return o.ServiceName }
func (o *config) GetServicePort() int       { return o.ServicePort }
func (o *config) GetServiceVersion() string { return o.ServiceVersion }

// Setter

func (o *FieldManager) SetServiceName(s string) *FieldManager {
	if s != "" {
		o.config.ServiceName = s
	}
	return o
}

func (o *FieldManager) SetServicePort(n int) *FieldManager {
	if n > 0 {
		o.config.ServicePort = n
	}
	return o
}

func (o *FieldManager) SetServiceVersion(s string) *FieldManager {
	if s != "" {
		o.config.ServiceVersion = s
	}
	return o
}
