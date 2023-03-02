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
// date: 2023-03-01

// Package logger_file
// 日志[异步]写入到 文件.
package logger_file

import (
	"context"
	"fmt"
	"github.com/fuyibing/log/v5/conf"
	"github.com/fuyibing/log/v5/formatters/logger_file"
	"github.com/fuyibing/log/v5/traces"
	"github.com/fuyibing/util/v8/process"
	"sync/atomic"
	"time"
)

type Exporter struct {
	bucket     traces.Bucket
	formatter  traces.LoggerFormatter
	name       string
	processing int32
	processor  process.Processor
	writer     *writer
}

func New() traces.LoggerExporter { return (&Exporter{}).init() }

// Processor 获取类进程.
func (o *Exporter) Processor() process.Processor { return o.processor }

// Send 发送日志.
func (o *Exporter) Send(log traces.Log) error {
	if !o.processor.Healthy() {
		return o.send(log)
	}

	// 加数据桶.
	backlog, err := o.bucket.Add(log)

	// 入桶出错.
	if err != nil {
		return err
	}

	// 立即上报.
	// 当桶中积压的数据达到阈值时, 立即上报(写入文件).
	if backlog >= conf.Config.GetBucketBatch() {
		go o.pop()
	}
	return nil
}

// SetFormatter 设置格式化.
func (o *Exporter) SetFormatter(formatter traces.LoggerFormatter) {
	o.formatter = formatter
}

// /////////////////////////////////////////////////////////////////////////////
// Bucket operations
// /////////////////////////////////////////////////////////////////////////////

func (o *Exporter) format(log traces.Log) string {
	if o.formatter != nil {
		return o.formatter.String(log)
	}

	return fmt.Sprintf("[%-15s][%5s] %s",
		log.GetTime().Format("15:04:05.999999"),
		log.GetLevel(), log.GetText(),
	)
}

func (o *Exporter) init() *Exporter {
	o.bucket = traces.NewBucket(conf.Config.GetBucketCapacity())
	o.formatter = logger_file.NewFormatter()
	o.writer = (&writer{}).init()

	o.name = "exporter.logger.file"
	o.processor = process.New(o.name).
		After(o.onAfter).
		Before(o.onBefore).
		Callback(o.onCall).
		Panic(o.onPanic)
	return o
}

func (o *Exporter) onAfter(ctx context.Context) (ignored bool) {
	// 处理完成.
	//
	// - 上传完成
	// - 数据桶为空.
	if concurrency := atomic.LoadInt32(&o.processing); concurrency > 0 || o.bucket.Count() > 0 {
		if concurrency < conf.Config.GetBucketConcurrency() {
			go o.pop()
		}
		time.Sleep(time.Millisecond * 400)
		return o.onAfter(ctx)
	}
	return
}

func (o *Exporter) onBefore(_ context.Context) (ignored bool) {
	return
}

func (o *Exporter) onCall(ctx context.Context) (ignored bool) {
	// 定时上报.
	ticker := time.NewTicker(time.Duration(conf.Config.GetBucketFrequency()) * time.Millisecond)

	for {
		select {
		case <-ticker.C:
			go o.pop()
		case <-ctx.Done():
			return
		}
	}
}

func (o *Exporter) onPanic(_ context.Context, v interface{}) {
	traces.InternalError("<%s> %v", o.name, v)
}

func (o *Exporter) pop() {
	if concurrency := atomic.AddInt32(&o.processing, 1); concurrency > conf.Config.GetBucketConcurrency() {
		atomic.AddInt32(&o.processing, -1)
		return
	}

	var (
		count int
		err   error
		items []interface{}
	)

	if items, _, count = o.bucket.Popn(conf.Config.GetBucketBatch()); count == 0 {
		atomic.AddInt32(&o.processing, -1)
		return
	}

	if err = o.send(items...); err != nil {
		traces.InternalError("<%s> %v", o.name, err)
	}

	atomic.AddInt32(&o.processing, -1)
	o.pop()
}

func (o *Exporter) send(logs ...interface{}) (err error) {
	var (
		i    = 0
		path string
		t    time.Time
		text string
	)

	for _, v := range logs {
		if log, ok := v.(traces.Log); ok {
			if i++; i == 1 {
				t = log.GetTime()
			}
			text += fmt.Sprintf("%s\n", o.format(log))
		}
	}

	if i == 0 {
		return
	}

	if _, path, err = o.writer.check(t); err != nil {
		return
	}

	err = o.writer.write(path, text)
	return
}
