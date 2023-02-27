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
// date: 2023-02-27

// Package log.
//
// 集成 Trace 的 Log 中间件, 遵循 OpenTelemetry 规范. Log 部分支持 term,
// file, kafka, aliyun sls 可配置, Trace 部分支持 term, jaeger, zipkin,
// aliyun sls. 其中 term (输出在终端控制台) 模式一般在开发环境时使用, 此模式下
// Log/Trace 为同步模式, 反之为异步.
//
// 配置: config/log.yaml
package log
