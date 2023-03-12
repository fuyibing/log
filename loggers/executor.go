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

package loggers

import (
	"github.com/fuyibing/util/v8/process"
)

type (
	// Executor
	// 日志执行器.
	Executor interface {
		// Processor
		// 类进程.
		//
		// 获取具体执行器的类进程, 基于此类进程启动/退出服务.
		Processor() (processor process.Processor)

		// Publish
		// 发布日志.
		Publish(logs ...Log) (err error)

		// SetFormatter
		// 设置格式.
		SetFormatter(formatter Formatter)
	}
)