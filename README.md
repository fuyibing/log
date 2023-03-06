# 链路日志

集成 `Trace` 的 `Log` 中间件, 遵循 `OpenTelemetry` 规范.

```go
import github.com/fuyibing/log/v5
```

## 配置

> 包初始化时，自动从配置文件(config/log.yaml)文件中读取配置参数，若未指定则使用默认值

### 公共参数

H3

### Logger

#### Term

```yaml
logger-exporter: "term"
```

#### File

```yaml
logger-exporter: "file"
file-logger:
  path: "/var/logs"
  folder: "2006-01"
  name: "2006-01-02"
  ext: "log"
```

#### Kafka

### Tracer

### Term

### File

### Kafka
