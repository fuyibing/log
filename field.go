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

package log

import (
	"github.com/fuyibing/log/v5/base"
	"github.com/fuyibing/log/v5/conf"
	"github.com/fuyibing/log/v5/exporters"
	"github.com/fuyibing/log/v5/tracer"
	"sync"
)

var (
	fieldPool sync.Pool
)

type (
	// FieldManager 附加 Key/Value 到日志.
	//
	//   log.Field().
	//       Add("key", "value").
	//       Debug("message")
	FieldManager interface {
		// Add 添加 Key/Value 属性.
		Add(key string, value interface{}) FieldManager

		// Debug 记录DEBUG级日志.
		Debug(text string, args ...interface{})

		// Error 记录ERROR级日志.
		Error(text string, args ...interface{})

		// Fatal 记录FATAL级日志.
		Fatal(text string, args ...interface{})

		// Info 记录INFO级日志.
		Info(text string, args ...interface{})

		// Warn 记录WARN级日志.
		Warn(text string, args ...interface{})
	}

	field struct{ Attribute base.Attribute }
)

// Field 创建 Key/Value 空属性.
func Field() FieldManager {
	if g := fieldPool.Get(); g != nil {
		return g.(*field).
			before()
	}

	return (&field{}).
		before()
}

// Add 添加 Key/Value 属性.
func (o *field) Add(key string, value interface{}) FieldManager {
	o.Attribute.Add(key, value)
	return o
}

// Debug 记录DEBUG级日志.
func (o *field) Debug(text string, args ...interface{}) {
	if conf.Config.DebugOn() {
		o.sendLogger(base.Debug, text, args...)
	}
}

// Error 记录ERROR级日志.
func (o *field) Error(text string, args ...interface{}) {
	if conf.Config.ErrorOn() {
		o.sendLogger(base.Error, text, args...)
	}
}

// Fatal 记录FATAL级日志.
func (o *field) Fatal(text string, args ...interface{}) {
	if conf.Config.FatalOn() {
		o.sendLogger(base.Fatal, text, args...)
	}
}

// Info 记录INFO级日志.
func (o *field) Info(text string, args ...interface{}) {
	if conf.Config.InfoOn() {
		o.sendLogger(base.Info, text, args...)
	}
}

// Warn 记录WARN级日志.
func (o *field) Warn(text string, args ...interface{}) {
	if conf.Config.WarnOn() {
		o.sendLogger(base.Warn, text, args...)
	}
}

// /////////////////////////////////////////////////////////////////////////////
// Access and constructor
// /////////////////////////////////////////////////////////////////////////////

func (o *field) before() *field {
	o.Attribute = base.Attribute{}
	return o
}

func (o *field) sendLogger(level base.Level, text string, args ...interface{}) {
	// 释放实例.
	defer func() {
		o.Attribute = nil
		fieldPool.Put(o)
	}()

	// 构建日志.
	x := tracer.NewLog(level, text, args...)
	x.Attribute.Copy(o.Attribute)
	exporters.Exporter.PutLogger(x)
}
