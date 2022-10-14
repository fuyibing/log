// author: wsfuyibing <websearch@163.com>
// date: 2022-10-13

package term

var Config = &Configuration{}

// Configuration
// 终端配置.
type Configuration struct {
	// 终端着色.
	// 是否依据不同级别的日志, 输出不同的颜色.
	Color bool `yaml:"color"`

	// 时间格式.
	TimeFormat string `yaml:"time"`
}

func (o *Configuration) Defaults(x *Configuration) {
	if x.Color {
		o.Color = x.Color
	}

	if x.TimeFormat != "" {
		o.TimeFormat = x.TimeFormat
	}
}
