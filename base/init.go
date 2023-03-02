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

package base

import (
	"sync"
)

var (
	// Resource 全局资源.
	//
	// 上报链路数据时携带此参数, 如进程ID, 主机名, IP地址等.
	Resource = Attribute{}
)

func init() {
	new(sync.Once).Do(func() {
		Id = (&id{}).init()
	})
}
