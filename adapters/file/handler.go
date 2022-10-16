// author: wsfuyibing <websearch@163.com>
// date: 2022-10-15

package file

import (
	"context"
	"fmt"
	"sync"

	"github.com/fuyibing/log/v3/base"
	"github.com/fuyibing/log/v3/formatters"
)

type (
	handler struct {
		engine base.AdapterEngine
		mapper map[string]Writer
		mu     *sync.RWMutex
	}
)

// New
// 创建执行器.
func New() base.AdapterEngine {
	return (&handler{}).init()
}

// Log
// 写入日志.
func (o *handler) Log(line *base.Line) {
	// 1. 无效日志.
	if line == nil {
		return
	}

	// 2. 写入文件.
	var err error

	// 2.1 捕获结果.
	defer func() {
		// 捕获异常.
		if r := recover(); r != nil {
			err = fmt.Errorf("panic on file adapter: %v", r)
		}

		// 写入完成.
		if err == nil || o.engine == nil {
			line.Release()
			return
		}

		// 降级处理
		o.engine.Log(line.WithError(err))
	}()

	// 2.2 写入过程.
	//     若文件写入失败, 将失败原因追加在内容后面, 并转发给终端
	//     适配器负责打印.
	err = o.get(line).Write(formatters.Formatter.AsFile(line))
}

// Parent
// 绑定上级/降级.
func (o *handler) Parent(engine base.AdapterEngine) base.AdapterEngine {
	o.engine = engine
	return o
}

// Start
// 启动服务.
func (o *handler) Start(ctx context.Context) {
	if o.engine != nil {
		o.engine.Start(ctx)
	}
}

// 检查文件.
func (o *handler) get(line *base.Line) Writer {
	o.mu.Lock()
	defer o.mu.Unlock()

	key := fmt.Sprintf("%s/%s", line.Time.Format(Config.Folder), line.Time.Format(Config.Name))

	if v, ok := o.mapper[key]; ok {
		return v
	}

	v := (&writer{}).init(line.Time)
	o.mapper[key] = v
	return v
}

// 构造实例.
func (o *handler) init() *handler {
	o.mapper = make(map[string]Writer, 0)
	o.mu = new(sync.RWMutex)
	return o
}
