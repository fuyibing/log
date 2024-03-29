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

package tracers

import (
	cr "crypto/rand"
	eb "encoding/binary"
	"encoding/hex"
	mr "math/rand"
	"sync"
)

type (
	id struct {
		sync.Mutex
		data   int64
		err    error
		random *mr.Rand
	}
)

// SpanIdFromHex
// create and return SpanId component based on specified hex string.
func (o *id) SpanIdFromHex(s string) SpanId {
	r := SpanId{}
	if d, de := hex.DecodeString(s); de == nil {
		copy(r[:], d)
	}
	return r
}

// SpanIdNew
// create and return SpanId component.
func (o *id) SpanIdNew() SpanId {
	o.Lock()
	defer o.Unlock()

	s := SpanId{}
	o.random.Read(s[:])
	return s
}

// TraceIdFromHex
// create and return TraceId component based on specified hex string.
func (o *id) TraceIdFromHex(s string) TraceId {
	r := TraceId{}
	if d, de := hex.DecodeString(s); de == nil {
		copy(r[:], d)
	}
	return r
}

// TraceIdNew
// create and return TraceId component.
func (o *id) TraceIdNew() TraceId {
	o.Lock()
	defer o.Unlock()

	s := TraceId{}
	o.random.Read(s[:])
	return s
}

// /////////////////////////////////////////////////////////////////////////////
// Access and constructor
// /////////////////////////////////////////////////////////////////////////////

func (o *id) init() *id {
	o.err = eb.Read(cr.Reader, eb.LittleEndian, &o.data)
	o.random = mr.New(mr.NewSource(o.data))
	return o
}
