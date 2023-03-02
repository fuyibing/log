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
		GetExt() string
		GetFolder() string
		GetName() string
		GetPath() string
	}

	fileLogger struct {
		// 根目录.
		// 默认: ./
		Path string `yaml:"path"`

		// 目录名.
		// 默认: 2006-01
		Folder string `yaml:"folder"`

		// 文件名.
		// 默认: 2006-01-02
		Name string `yaml:"name"`

		// 扩展名.
		// 默认: log
		Ext string `yaml:"ext"`
	}
)

// Child.

func (o *config) GetFileLogger() ConfigFileLogger { return o.FileLogger }

// Getter

func (o *fileLogger) GetExt() string    { return o.Ext }
func (o *fileLogger) GetFolder() string { return o.Folder }
func (o *fileLogger) GetName() string   { return o.Name }
func (o *fileLogger) GetPath() string   { return o.Path }

// Setter

func (o *FieldManager) SetFileLoggerExt(s string) *FieldManager {
	if s != "" {
		o.config.FileLogger.Ext = s
	}
	return o
}

func (o *FieldManager) SetFileLoggerFolder(s string) *FieldManager {
	if s != "" {
		o.config.FileLogger.Folder = s
	}
	return o
}

func (o *FieldManager) SetFileLoggerName(s string) *FieldManager {
	if s != "" {
		o.config.FileLogger.Name = s
	}
	return o
}

func (o *FieldManager) SetFileLoggerPath(s string) *FieldManager {
	if s != "" {
		o.config.FileLogger.Path = s
	}
	return o
}

// Initialize

func (o *fileLogger) initDefaults() {
	if o.Path == "" {
		o.Path = "./logs"
	}
	if o.Folder == "" {
		o.Folder = "2006-01"
	}
	if o.Name == "" {
		o.Name = "2006-01-02"
	}
	if o.Ext == "" {
		o.Ext = "log"
	}
}
