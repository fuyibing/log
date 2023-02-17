// author: wsfuyibing <websearch@163.com>
// date: 2023-02-17

package conf

var (
	defaultFilePath    = "./logs"
	defaultFileFolder  = "2006-01"
	defaultFileName    = "2006-01-02"
	defaultFileExtName = "log"
)

type (
	FileConfig struct {
		Path    string `yaml:"path"`
		Folder  string `yaml:"folder"`
		Name    string `yaml:"name"`
		ExtName string `yaml:"ext-name"`
	}
)

func (o *FileConfig) defaults() {
	if o.Path == "" {
		o.Path = defaultFilePath
	}
	if o.Folder == "" {
		o.Folder = defaultFileFolder
	}
	if o.Name == "" {
		o.Name = defaultFileName
	}
	if o.ExtName == "" {
		o.ExtName = defaultFileExtName
	}
}
