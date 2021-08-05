// author: wsfuyibing <websearch@163.com>
// date: 2021-02-22

package interfaces

// 客户端接口
type ClientInterface interface {
	// 添加Debug日志.
	Debug(text string)

	// 添加Debug日志, 支持格式化.
	Debugf(text string, args ...interface{})

	// 添加Debug日志, 支持格式化和请求链.
	Debugfc(ctx interface{}, text string, args ...interface{})

	// 添加Info日志.
	Info(text string)

	// 添加Info日志, 支持格式化.
	Infof(text string, args ...interface{})

	// 添加Info日志, 支持格式化和请求链.
	Infofc(ctx interface{}, text string, args ...interface{})

	// 添加Warn日志.
	Warn(text string)

	// 添加Warn日志, 支持格式化.
	Warnf(text string, args ...interface{})

	// 添加Warn日志, 支持格式化和请求链.
	Warnfc(ctx interface{}, text string, args ...interface{})

	// 添加Error日志.
	Error(text string)

	// 添加Error日志, 支持格式化.
	Errorf(text string, args ...interface{})

	// 添加Error日志, 支持格式化和请求链.
	Errorfc(ctx interface{}, text string, args ...interface{})

	Panic(text string)
	Panicf(text string, args ...interface{})
	Panicfc(ctx interface{}, text string, args ...interface{})
}
