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
[2022-10-14 15:36:41.140847][10.3.6.14:8527][FLOG][DEBUG][pid=45334] debug message
[2022-10-14 15:36:41.141365][10.3.6.14:8527][FLOG][INFO][pid=45334] info message
[2022-10-14 15:36:41.141401][10.3.6.14:8527][FLOG][INFO][pid=45334] warn message
[2022-10-14 15:36:41.141539][10.3.6.14:8527][FLOG][ERROR][pid=45334] error message
```

### 3. Redis 订阅

`异步` `降级`

> Abount

### 4. Kafka 订阅

`异步` `降级`

> About

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
