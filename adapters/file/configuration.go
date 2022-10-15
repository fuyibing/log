// author: wsfuyibing <websearch@163.com>
// date: 2022-10-15

package file

var Config *Configuration

const (
	defaultPath   = "./logs"
	defaultFolder = "2006-01"
	defaultName   = "2006-01-02"
)

// Configuration
// 基础配置.
type Configuration struct {
	// 存储目录.
	// 例如: /var/logs
	Path string `yaml:"path"`

	// 按时间分目录.
	// 格式: 2006-01
	// 参考: 2022-10 (完整: /var/logs/2022-10)
	Folder string `yaml:"folder"`

	// 按时间命名.
	// 格式: 2006-01-02
	// 参考: 2022-10-01 (完整: /var/logs/2022-10/2022-10-01.log)
	Name string `yaml:"name"`
}

// Override
// 覆盖配置.
func (o *Configuration) Override(c *Configuration) *Configuration {
	if c.Path != "" {
		o.Path = c.Path
	}
	if c.Folder != "" {
		o.Folder = c.Folder
	}
	if c.Name != "" {
		o.Name = c.Name
	}
	return o
}

// 构造实例.
func (o *Configuration) init() *Configuration {
	o.Path = defaultPath
	o.Folder = defaultFolder
	o.Name = defaultName
	return o
}
