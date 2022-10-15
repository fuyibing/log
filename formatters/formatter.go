// author: wsfuyibing <websearch@163.com>
// date: 2022-10-15

package formatters

import "github.com/fuyibing/log/v3/base"

var (
	Formatter FormatterManager
)

type (
	// FormatterHandler
	// 格式化执行器.
	FormatterHandler func(line *base.Line, err error) string

	// FormatterManager
	// 格式化管理器.
	FormatterManager interface {
		// AsFile
		// 用于写入文件.
		AsFile(line *base.Line, err error) string

		// AsJson
		// 用于写入Redis/Kafka.
		AsJson(line *base.Line, err error) string

		// AsTerm
		// 用于终端打印.
		AsTerm(line *base.Line, err error) string

		// SetFileFormatter
		// 自定义File格式.
		SetFileFormatter(x FormatterHandler) FormatterManager

		// SetJsonFormatter
		// 自定义JSON格式.
		SetJsonFormatter(x FormatterHandler) FormatterManager

		// SetTermFormatter
		// 自定义Term格式.
		SetTermFormatter(x FormatterHandler) FormatterManager
	}

	formatter struct {
		term, file, json FormatterHandler
	}
)

// AsFile
// 用于写入文件.
func (o *formatter) AsFile(line *base.Line, err error) string {
	return o.file(line, err)
}

// AsJson
// 用于写入Redis/Kafka.
func (o *formatter) AsJson(line *base.Line, err error) string {
	return o.json(line, err)
}

// AsTerm
// 用于终端打印.
func (o *formatter) AsTerm(line *base.Line, err error) string {
	return o.term(line, err)
}

// SetFileFormatter
// 自定义File格式.
func (o *formatter) SetFileFormatter(x FormatterHandler) FormatterManager {
	if x != nil {
		o.file = x
	}
	return o
}

// SetJsonFormatter
// 自定义JSON格式.
func (o *formatter) SetJsonFormatter(x FormatterHandler) FormatterManager {
	if x != nil {
		o.json = x
	}
	return o
}

// SetTermFormatter
// 自定义Term格式.
func (o *formatter) SetTermFormatter(x FormatterHandler) FormatterManager {
	if x != nil {
		o.term = x
	}
	return o
}

// 构造实例.
func (o *formatter) init() *formatter {
	o.file = (&fileFormatter{}).init().Format
	o.json = (&jsonFormatter{}).init().Format
	o.term = (&termFormatter{}).init().Format
	return o
}
