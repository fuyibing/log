# 日志

> package `github.com/fuyibing/log/v3`.

## 一、适配器

### 1. 终端打印

`同步`

> 日志内容在终端(`terminal`)上打印, 此过程下内容定向到标准输出流(`stdout`)中.

```yaml
# config/log.yaml
adapter: "term"                             # 适配器名称
term:
  color: false                              # 是否带颜色输出
  time: "15:04:05"                          # 日志时间格式
```

```log
[09:10:11][DEBUG] debug message
[09:10:11][INFO] info message
[09:10:11][WARN] warn message
[09:10:11][ERROR] error message
```

### 2. 文件存储

`同步` `降级`

> 日志内容输出到指定目录/文件中, 此模式下对分步式部署不友好, 不同节点的日志落到各自的机器/容器上. 当写入文件出错时降级为 `TERM` 模式.

```yaml
# config/log.yaml
adapter: "file"                             # 适配器名称
file:
  path: "/var/logs"                         # 日志写到哪个目录
  folder: "2006-01"                         # 按时间拆分目录(/var/logs/2022-10)
  name: "2006-01-02"                        # 按时间切割文件(/var/logs/2022-10/2022-10-01.log)
```

```log
[2022-10-14 15:36:41.140847][10.3.6.14:8527][FLOG][DEBUG][PID=45334] debug message
[2022-10-14 15:36:41.141365][10.3.6.14:8527][FLOG][INFO][PID=45334] info message
[2022-10-14 15:36:41.141401][10.3.6.14:8527][FLOG][INFO][PID=45334] warn message
[2022-10-14 15:36:41.141539][10.3.6.14:8527][FLOG][ERROR][PID=45334] error message
```

### 3. Redis 订阅

`异步` `降级`

> 日志内容写到Redis服务上, 其它服务可以消费Redis并转发到 `Kafka` `Aliyun log service` 等中间件上.

```yaml
# config/log.yaml
adapter: "redis"                            # 适配器名称
redis:
  concurrency: 5                            # 写入Redis最大并发 (默认: 5)
  limit: 100                                # 批量写入时, 每批最多数量 (默认: 100)
  ticker: 500                               # 每隔 350 毫秒, 固定上报一次日志 (默认: 500)
  key-prefix: "logger"                      # 写入到Redis的Key加上前缀.
  key-lifetime: 3600                        # 写入到Redis后, 3600秒后没有被消费时丢弃.
  key-list: "list"                          # 将Key加入到logger:list列表中, 消费时从此Key上使用LPOP
  address:
  password:
  database:
  max-active:
  max-idle:
  timeout:                                  # 连接Redis超时时长 (默认: 5)
  read-timeout:                             # 读数据超时时长 (默认: 3)
  write-timeout:                            # 写数据超时时长 (默认: 10)
  wait: true                                # 连接池不足时, 是否等待 (默认: true)
```

### 4. Kafka 订阅

`异步` `降级`

> About

```yaml
# config/log.yaml
adapter: "kafka"                            # 适配器名称
kafka:
```

## Example

```text
// Simple.
log.Info("Info message.")
log.Infof("info message line %d.", 1)

// Context.
ctx := log.NewContext()
log.Infofc(ctx, "info message.")
log.Infofc(ctx, "info message line %d.", 100)
```
