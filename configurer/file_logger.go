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
	ConfigLoggerFile interface {
		GetFileLogger() FileLogger
	}

	FileLogger interface {
		GetExt() string
		GetFolder() string
		GetName() string
		GetPath() string
	}

	fileLogger struct {
		// 扩展名.
		// 默认: log
		Ext string `yaml:"ext"`

		// 目录分隔.
		// 默认: 2006-01 (即按月分隔)
		Folder string `yaml:"folder"`

		// 文件全名.
		// 默认: 2006-01-02
		Name string `yaml:"name"`

		// 存储位置.
		// 默认: ./logs (在项目目录)
		Path string `yaml:"path"`
	}
)

// Getter

func (o *config) GetFileLogger() FileLogger { return o.FileLogger }

func (o *fileLogger) GetExt() string    { return o.Ext }
func (o *fileLogger) GetFolder() string { return o.Folder }
func (o *fileLogger) GetName() string   { return o.Name }
func (o *fileLogger) GetPath() string   { return o.Path }

// Setter.

func (o *Setter) SetFileLoggerExt(s string) *Setter {
	o.config.FileLogger.Ext = s
	return o
}

func (o *Setter) SetFileLoggerFolder(s string) *Setter {
	o.config.FileLogger.Folder = s
	return o
}

func (o *Setter) SetFileLoggerName(s string) *Setter {
	o.config.FileLogger.Name = s
	return o
}

func (o *Setter) SetFileLoggerPath(s string) *Setter {
	o.config.FileLogger.Path = s
	return o
}

// Defaults

func (o *fileLogger) initDefaults() {
	if o.Ext == "" {
		o.Ext = defaultFileLoggerExt
	}
	if o.Folder == "" {
		o.Folder = defaultFileLoggerFolder
	}
	if o.Name == "" {
		o.Name = defaultFileLoggerName
	}
	if o.Path == "" {
		o.Path = defaultFileLoggerPath
	}
}
