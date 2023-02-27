# 链路日志

集成 `Trace` 的 `Log` 中间件, 遵循 `OpenTelemetry` 规范. Log 部分支持 `term`, `file`, `kafka`, `aliyun sls`
可配置, Trace 部分支持 `term`, `jaeger`, `zipkin`, `aliyun sls`. 其中 term (输出在终端控制台) 模式一般在
开发环境时使用, 此模式下 Log/Trace 为同步 `SYNC` 模式, 反之为异步 `ASYNC`.

```go
import github.com/fuyibing/log/v5
```

## 配置

本包在首次加载时, 自动从 `config/log.yaml` 文件中读取配置参数.

```yaml
service-name: "my-app"
service-port: 8080
service-version: 1.2.3
# Log options
logger-exporter: "term"
logger-level: "debug"
# Trace options
tracer-exporter: "jaeger"
tracer-topic: "log-trace"
# Jaeger options
jaeger-tracer:
  content-type: "application/x-thrift"
  endpoint: "http://localhost:14268/api/traces"
```

手动配置

```text
    conf.Config.With(
		conf.ServiceName("my-app"),
		conf.ServicePort(8080),
		conf.ServiceVersion("1.2.3"),

		conf.LogExporter("term"),
		conf.LogLevel("info"),

		conf.TracerExporter("jaeger"),
		conf.TracerTopic("log-trace"),
		conf.JaegerTracerContentType("application/x-thrift"),
		conf.JaegerTracerEndpoint("http://localhost:14268/api/traces"),
	)

	cores.Registry.Update()
```

管理器服务, 使用前调用 `log.Manager.Start` 方法启动服务, 并在服务退出前 `log.Manager.Stop` 方法退出.

```text
log.Manager.Start(ctx)
defer log.Manager.Stop()
```

## 怎么使用

### 简单

```text
log.Info("message")
log.Info("message id=%d", 1)
```

### 混合

```text
log.Field{}.Add("key", "value").Info("message")
log.Field{}.Add("key", "value").Info("message id=%d", 1)
```

## 链路

```text
trace := log.NewTrace("trace")

span := trace.New("span")
defer span.End()

span.Logger().Info("message")
span.Logger().Info("message %d", 1)
span.Logger().Add("key", "value").Info("message %d", 1)
```
