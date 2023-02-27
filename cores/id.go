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
// date: 2023-02-25

package cores

import (
	cr "crypto/rand"
	eb "encoding/binary"
	"encoding/hex"
	mr "math/rand"
	"sync"
)

var (
	// Identify
	// 全局ID生成器实例.
	Identify IdentifyGenerator
)

type (
	// IdentifyGenerator
	// ID生成器接口.
	IdentifyGenerator interface {
		// GenSpanId
		// 生成 SpanId 随机值.
		GenSpanId() SpanId

		// GenTraceId
		// 生成 TraceId 随机值.
		GenTraceId() TraceId

		// HexSpanId
		// 基于长度为16的字符串, 反向生成 SpanId.
		HexSpanId(s string) SpanId

		// HexTraceId
		// 基于长度为32的字符串, 反向生成 TraceId.
		HexTraceId(s string) TraceId
	}

	identify struct {
		sync.Mutex
		data   int64
		err    error
		random *mr.Rand
	}
)

// GenSpanId
// 生成 SpanId 随机值.
func (o *identify) GenSpanId() SpanId {
	o.Lock()
	defer o.Unlock()

	s := SpanId{}
	o.random.Read(s[:])
	return s
}

// GenTraceId
// 生成 TraceId 随机值.
func (o *identify) GenTraceId() TraceId {
	o.Lock()
	defer o.Unlock()

	s := TraceId{}
	o.random.Read(s[:])
	return s
}

// HexSpanId
// 基于长度为16的字符串, 反向生成 SpanId.
func (o *identify) HexSpanId(s string) SpanId {
	res := SpanId{}
	if d, de := hex.DecodeString(s); de == nil {
		copy(res[:], d)
	}
	return res
}

// HexTraceId
// 基于长度为32的字符串, 反向生成 TraceId.
func (o *identify) HexTraceId(s string) TraceId {
	res := TraceId{}
	if d, de := hex.DecodeString(s); de == nil {
		copy(res[:], d)
	}
	return res
}

func (o *identify) init() *identify {
	o.err = eb.Read(cr.Reader, eb.LittleEndian, &o.data)
	o.random = mr.New(mr.NewSource(o.data))
	return o
}
