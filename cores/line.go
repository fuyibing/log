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
// date: 2023-02-26

package cores

import (
	"fmt"
	"github.com/fuyibing/log/v5/base"
	"time"
)

type (
	// Line
	// is the component of the log.
	Line interface {
		// GetAttr
		// returns an Attr component of the log.
		GetAttr() Attr

		// GetLevel
		// returns a base.Level of the log.
		GetLevel() base.Level

		// GetText
		// returns content of the log.
		GetText() string

		// GetTime
		// returns a time.Time of the log begin.
		GetTime() time.Time
	}

	line struct {
		attr  Attr
		level base.Level
		text  string
		time  time.Time
	}
)

// NewLine
// return a component for custom log.
func NewLine(level base.Level, text string, args ...interface{}) Line {
	return &line{
		attr:  NewAttr(),
		level: level,
		text:  fmt.Sprintf(text, args...),
		time:  time.Now(),
	}
}

// GetAttr
// returns an Attr component of the log.
func (o *line) GetAttr() Attr {
	return o.attr
}

// GetLevel
// returns a base.Level of the log.
func (o *line) GetLevel() base.Level {
	return o.level
}

// GetText
// returns content of the log.
func (o *line) GetText() string {
	return o.text
}

// GetTime
// returns a time.Time of the log begin.
func (o *line) GetTime() time.Time {
	return o.time
}
