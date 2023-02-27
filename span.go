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

package log

import (
	"context"
	"github.com/fuyibing/log/v5/base"
	"github.com/fuyibing/log/v5/cores"
)

// Span
// 从上下文上读取链路跨度, 若不存在则返回nil.
//
// 场景: 一般使用在链路跨度的中间过程中, 例如在一个链路跨度中使用了SQL语句, 可以将
// SQL打印在链路的日志列表中.
//
// 效果: /examples/images/log.span.png
func Span(ctx context.Context) (span cores.Span, exists bool) {
	if v := ctx.Value(base.ContextKeySpan); v != nil {
		span, exists = v.(cores.Span)
	}
	return
}