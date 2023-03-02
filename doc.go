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

// Package log
// 集成了 Trace 的 Log 中间件, 遵循 OpenTelemetry 规范.
//
// 启动
//   // 在 main() 函数的主口调用此方法, 启动日志服务.
//   log.Logger.Start(ctx)
//
// 退出
//   // 在 main() 函数退出前调用此方法, 以确保异步上报的数据能够正常完成, 避
//   // 免丢失
//   log.Logger.Stop()
//
// 示例
//   func main(){
//       log.Logger.Start(ctx)
//       defer log.Logger.Stop()
//
//       // ... 更多逻辑
//   }
package log
