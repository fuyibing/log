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
// date: 2023-03-04

package tracers

import (
	"bytes"
	"encoding/hex"
)

type (
	// TraceId
	// 链路ID.
	TraceId [16]byte
)

// IsValid
// 校验链路ID.
func (o TraceId) IsValid() bool {
	return !bytes.Equal(o[:], nilTraceId[:])
}

// String
// 转成16进制字符串.
func (o TraceId) String() string {
	return hex.EncodeToString(o[:])
}
