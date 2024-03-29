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
	// SpanId
	// component stored as 8-byte array, encoded / decoded as 16 char
	// string.
	SpanId [8]byte
)

// IsValid
// return true if value is not zero string.
func (o SpanId) IsValid() bool { return !bytes.Equal(o[:], nilSpanId[:]) }

// String
// return hex string, 16 chars.
func (o SpanId) String() string { return hex.EncodeToString(o[:]) }
