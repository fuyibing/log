// author: wsfuyibing <websearch@163.com>
// date: 2022-10-15

package term

import "github.com/fuyibing/log/v3/base"

var (
	Colors = map[base.Level][]int{
		base.Debug: {30, 0}, // 白底黑字
		base.Info:  {32, 0}, // 白底绿字
		base.Warn:  {33, 0}, // 白底黄字
		base.Error: {31, 0}, // 白底红字
	}
	Config *Configuration

	defaultColor = false
)

// Configuration
// 基础配置.
type Configuration struct {
	// 终端着色.
	// 是否依据不同级别的日志, 输出不同的颜色.
	Color *bool `yaml:"color"`
}

// Override
// 覆盖配置.
func (o *Configuration) Override(x *Configuration) *Configuration {
	if x.Color != nil {
		o.Color = x.Color
	}
	return o
}

// 构造实例.
func (o *Configuration) init() *Configuration {
	o.Color = &defaultColor
	return o
}
