# 链路日志

`OpenTelemetry`
`Telemetry`
`Trace`

----

集成 `Trace` 的 `Log` 中间件, 遵循 `OpenTelemetry` 规范. Log 部分支持 `term`, `file`, `kafka`, `aliyun sls`
可配置, Trace 部分支持 `term`, `jaeger`, `zipkin`, `aliyun sls`. 其中 term (输出在终端控制台) 模式一般在
开发环境时使用, 此模式下 Log/Trace 为同步 `SYNC` 模式, 反之为异步 `ASYNC`.

```go
import github.com/fuyibing/log/v5
```