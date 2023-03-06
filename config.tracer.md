# Tracer

1. [X] `Jaeger` - 上报到 Jaeger
2. [X] `Zipkin` - 上报到 Zipkin
3. [X] `File` - 输出到文件中
4. [X] `Term` - 打印到终端/控制台

### 公共

> 打通上游传递过来的调用链路.

```yaml
open-tracing-sampled: "X-B3-Sampled"
open-tracing-span-id: "X-B3-Spanid"
open-tracing-trace-id: "X-B3-Traceid"
```

### 上报

##### {Jaeger}

```yaml
tracer-topic: "log-trace"                         # 任意
tracer-exporter: "jaeger"                         # 必须
jaeger-tracer:
  content-type: "application/x-thrift"            # API 格式
  endpoint: "http://localhost:14268/api/traces"   # API 地址
  username: ""                                    # Basic 用户名
  password: ""                                    # Basic 密码
```

##### {Zipkin}

```yaml
tracer-topic: "log-trace"                         # 任意
tracer-exporter: "zipkin"                         # 必须
zipkin-tracer:
  content-type: "application/json"                # API 格式
  endpoint: "http://localhost:9411/api/v2/spans"  # API 地址
```

##### {Term}

```yaml
tracer-exporter: "term"                           # 必须
```

##### File

> 链路数据输出到本地文件中. 例如 `/var/logs/2023-03/2023-03-01.trace`

```yaml
tracer-exporter: "file"                           # 必须
file-tracer:
  path: "/var/logs"                               # 存储位置
  folder: "2006-01"                               # 拆分目录
  name: "2006-01-02"                              # 文件格式
  ext: "trace"                                    # 日志扩展名
```
