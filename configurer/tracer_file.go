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
	// ConfigTracerFile
	// expose file adapter for tracer.
	ConfigTracerFile interface {
		GetFileTracer() FileTracer
	}

	// FileTracer
	// expose file tracer configuration methods.
	FileTracer interface {
		GetExt() string
		GetFolder() string
		GetName() string
		GetPath() string
	}

	fileTracer struct {
		Ext    string `yaml:"ext"`
		Folder string `yaml:"folder"`
		Name   string `yaml:"name"`
		Path   string `yaml:"path"`
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
