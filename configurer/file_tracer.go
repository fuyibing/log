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
	ConfigFileTracer interface {
		GetFileTracer() FileTracer
	}

	FileTracer interface {
		GetExt() string
		GetFolder() string
		GetName() string
		GetPath() string
	}

	fileTracer struct {
		// 扩展名.
		// 默认: trace
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

func (o *config) GetFileTracer() FileTracer { return o.FileTracer }

func (o *fileTracer) GetExt() string    { return o.Ext }
func (o *fileTracer) GetFolder() string { return o.Folder }
func (o *fileTracer) GetName() string   { return o.Name }
func (o *fileTracer) GetPath() string   { return o.Path }

// Setter.

func (o *Setter) SetFileTracerExt(s string) *Setter {
	o.config.FileTracer.Ext = s
	return o
}

func (o *Setter) SetFileTracerFolder(s string) *Setter {
	o.config.FileTracer.Folder = s
	return o
}

func (o *Setter) SetFileTracerName(s string) *Setter {
	o.config.FileTracer.Name = s
	return o
}

func (o *Setter) SetFileTracerPath(s string) *Setter {
	o.config.FileTracer.Path = s
	return o
}

// Defaults

func (o *fileTracer) initDefaults() {
	if o.Ext == "" {
		o.Ext = defaultFileTracerExt
	}
	if o.Folder == "" {
		o.Folder = defaultFileTracerFolder
	}
	if o.Name == "" {
		o.Name = defaultFileTracerName
	}
	if o.Path == "" {
		o.Path = defaultFileTracerPath
	}
}
