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
// date: 2023-03-05

// Package logger_file
// 输出到文件中, 例如: /var/logs/2023-03/2023-03-01.log.
package logger_file

import (
	"context"
	"fmt"
	"github.com/fuyibing/log/v5/common"
	"github.com/fuyibing/log/v5/configurer"
	"github.com/fuyibing/log/v5/loggers"
	"github.com/fuyibing/util/v8/process"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

type executor struct {
	sync.RWMutex

	bucket     common.Bucket
	folders    map[string]bool
	formatter  loggers.Formatter
	name       string
	processor  process.Processor
	processing int32
}

func New() loggers.Executor { return (&executor{}).init() }

// /////////////////////////////////////////////////////////////////////////////
// Interface methods
// /////////////////////////////////////////////////////////////////////////////

func (o *executor) Processor() process.Processor      { return o.processor }
func (o *executor) Publish(logs ...loggers.Log) error { return o.publish(logs...) }
func (o *executor) SetFormatter(v loggers.Formatter)  { o.formatter = v }

// /////////////////////////////////////////////////////////////////////////////
// Event methods
// /////////////////////////////////////////////////////////////////////////////

func (o *executor) onAfter(ctx context.Context) (ignored bool) {
	cc := atomic.LoadInt32(&o.processing)

	// 处理完成.
	// - 并行降低
	// - 空数据桶.
	if cc == 0 && o.bucket.IsEmpty() {
		return
	}

	// 加大并行.
	if cc < configurer.Config.GetBucketConcurrency() {
		go o.pop()
	}

	// 定时延后.
	time.Sleep(time.Millisecond * 100)
	return o.onAfter(ctx)
}

func (o *executor) onCall(ctx context.Context) (ignored bool) {
	common.InternalInfo("<%s> signal listening", o.name)

	// 定时收取.
	ti := time.NewTicker(time.Duration(configurer.Config.GetBucketFrequency()) * time.Millisecond)

	// 监听信号.
	for {
		select {
		case <-ti.C:
			go o.pop()
		case <-ctx.Done():
			return
		}
	}
}

func (o *executor) onPanic(_ context.Context, v interface{}) {
	common.InternalFatal("<%s> fatal: %v", o.name, v)
}

// /////////////////////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////////////////////

func (o *executor) init() *executor {
	o.bucket = common.NewBucket(configurer.Config.GetBucketCapacity())
	o.folders = make(map[string]bool)
	o.formatter = (&formatter{}).init()
	o.name = "logger.file"
	o.processor = process.New(o.name).
		After(o.onAfter).
		Callback(o.onCall).
		Panic(o.onPanic)

	return o
}

func (o *executor) pop() {
	// 限流控制.
	if cc := atomic.AddInt32(&o.processing, 1); cc > configurer.Config.GetBucketConcurrency() {
		atomic.AddInt32(&o.processing, -1)
		return
	}

	// 取出数据.
	var (
		list []loggers.Log
		redo = false
	)
	if items, _, count := o.bucket.Popn(configurer.Config.GetBucketBatch()); count > 0 {
		list = make([]loggers.Log, 0)
		redo = true
		// 遍历数据.
		for _, item := range items {
			if v, ok := item.(loggers.Log); ok {
				list = append(list, v)
			}
		}
		// 处理日志.
		if len(list) > 0 {
			if err := o.send(list...); err != nil {
				common.InternalInfo("<%s> send: %v", o.name, err)
			}
		}
	}

	// 恢复并行.
	atomic.AddInt32(&o.processing, -1)
	if redo {
		o.pop()
	}
}

func (o *executor) publish(logs ...loggers.Log) (err error) {
	var total int

	// 健康进程.
	if o.processor.Healthy() {
		// 数据入桶.
		for _, log := range logs {
			if total, err = o.bucket.Add(log); err != nil {
				return
			}
		}

		// 立即消费.
		if total >= configurer.Config.GetBucketBatch() {
			go o.pop()
		}
		return
	}

	// 立即写入.
	return o.send(logs...)
}

func (o *executor) send(logs ...loggers.Log) (err error) {
	// 暂无日志.
	if len(logs) == 0 {
		return
	}

	var (
		fp   *os.File
		path string
		text string
	)

	// 校验目录.
	if path, err = o.validate(logs[0].Time()); err != nil {
		return
	}

	// 格式日志.
	if text, err = o.formatter.String(logs...); err != nil {
		return
	}

	// 打开文件.
	if fp, err = os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm); err != nil {
		return
	}

	// 关闭文件.
	defer func() { _ = fp.Close() }()

	// 写入日志.
	_, err = fp.WriteString(fmt.Sprintf("%s\n", text))
	return
}

func (o *executor) validate(t time.Time) (string, error) {
	var (
		dir  = fmt.Sprintf("%s/%s", configurer.Config.GetFileLogger().GetPath(), t.Format(configurer.Config.GetFileLogger().GetFolder()))
		err  error
		path = fmt.Sprintf("%v/%s.%s", dir, t.Format(configurer.Config.GetFileLogger().GetName()), configurer.Config.GetFileLogger().GetExt())
	)

	o.Lock()
	defer o.Unlock()

	if _, ok := o.folders[dir]; !ok {
		if err = os.MkdirAll(dir, os.ModePerm); err == nil {
			o.folders[dir] = true
		}
	}

	return path, err
}
