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

import (
	"github.com/fuyibing/log/v5/base"
)

type Option func(c *configuration)

func LogLevel(l base.Level) Option {
	return func(c *configuration) {
		c.LoggerLevel = l
		c.updateState()
	}
}

func ServiceName(s string) Option {
	return func(c *configuration) {
		c.ServiceName = s
	}
}

func ServicePort(p int) Option {
	return func(c *configuration) {
		c.ServicePort = p
	}
}

func ServiceVersion(s string) Option {
	return func(c *configuration) {
		c.ServiceVersion = s
	}
}

func TracerExporter(s string) Option {
	return func(c *configuration) {
		c.TracerExporter = s
	}
}

func TracerTopic(s string) Option {
	return func(c *configuration) {
		c.TracerTopic = s
	}
}
