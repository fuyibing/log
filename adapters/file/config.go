// author: wsfuyibing <websearch@163.com>
// date: 2022-10-13

package file

var Config *Configuration

// Configuration
// 文件配置.
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

// Defaults
// 覆盖默认值.
func (o *Configuration) Defaults(x *Configuration) {
	if x.Path != "" {
		o.Path = x.Path
	}
	if x.Folder != "" {
		o.Folder = x.Folder
	}
	if x.Name != "" {
		o.Name = x.Name
	}
}

// 构造实例.
func (o *Configuration) init() *Configuration {
	o.Path = "./logs"
	o.Folder = "2006-01"
	o.Name = "2006-01-02"
	return o
}
