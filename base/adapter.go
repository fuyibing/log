// author: wsfuyibing <websearch@163.com>
// date: 2022-10-15

package base

import (
	"context"
	"strings"
)

type (
	// Adapter
	// 适配器类型.
	Adapter int

	// AdapterName
	// 适配器名称.
	AdapterName string

	// AdapterLog
	// 适配器执行.
	AdapterLog func(line *Line, err error)

	// AdapterEngine
	// 适配器引擎.
	AdapterEngine interface {
		// Log
		// 发送日志.
		Log(line *Line, err error)

		// Parent
		// 绑定上级/降级.
		Parent(engine AdapterEngine) AdapterEngine

		// Start
		// 启动服务.
		Start(ctx context.Context)
	}
)

// 适配器类型枚举.

const (
	UnknownAdapter Adapter = iota
	Term
	File
	Redis
	Kafka
)

// Adapter
// 适配器名称转类型.
func (s AdapterName) Adapter() (a Adapter) {
	switch strings.ToLower(string(s)) {
	case "term", "terminal":
		a = Term
	case "file":
		a = File
	case "redis":
		a = Redis
	case "kafka":
		a = Kafka
	default:
		a = UnknownAdapter
	}
	return
}
