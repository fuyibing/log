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
	// ConfigLoggerFile
	// expose file adapter for logger.
	ConfigLoggerFile interface {
		GetFileLogger() FileLogger
	}

	// FileLogger
	// expose file logger configuration methods.
	FileLogger interface {
		GetExt() string
		GetFolder() string
		GetName() string
		GetPath() string
	}

	fileLogger struct {
		Ext    string `yaml:"ext"`
		Folder string `yaml:"folder"`
		Name   string `yaml:"name"`
		Path   string `yaml:"path"`
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
