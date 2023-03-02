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
	FileLogger interface {
		GetFileLogger() ConfigFileLogger
	}

	// ConfigFileLogger
	// 文件日志.
	ConfigFileLogger interface {
		// GetExt
		// 获取扩展名.
		GetExt() string

		// GetFolder
		// 获取目录名.
		GetFolder() string

		// GetName
		// 获取文件名.
		GetName() string

		// GetPath
		// 存储位置.
		GetPath() string
	}

	fileLogger struct {
		Path   string `yaml:"path"`
		Folder string `yaml:"folder"`
		Name   string `yaml:"name"`
		Ext    string `yaml:"ext"`
	}
)

func (o *config) GetFileLogger() ConfigFileLogger { return o.FileLogger }

func (o *fileLogger) GetExt() string    { return "log" }
func (o *fileLogger) GetFolder() string { return "2006-01" }
func (o *fileLogger) GetName() string   { return "2006-01-02" }
func (o *fileLogger) GetPath() string   { return "./logs" }

func (o *fileLogger) initDefaults() {}
