# 日志

> package `github.com/fuyibing/log/v3`.

## 适配器

### TERM

`同步` `降级`

> 日志内容在终端(`terminal`)上打印, 此过程下内容定向到标准输出流(`stdout`)中.

```yaml
# config/log.yaml
adapter: "term"                     # 适配器名称
term:
  color: false                      # 是否带颜色输出
  time: "15:04:05"                  # 日志时间格式
```

```log
[09:10:11][DEBUG] debug message
[09:10:11][INFO] info message
[09:10:11][WARN] warn message
[09:10:11][ERROR] error message
```

**file**

**redis**

**kafka**

1. `term` - 在终端打印日志, 通常在开发机器上使用. `同步`
2. `file` - 写入到指定目录/文件中, 对分步式结构下不友好. `同步` `降级`
3. `redis` - 写入到Redis中. `异步` `降级`
4. `kafka` - 写入到Kafka中. `异步` `降级`

### Example

```text
// Simple.
log.Info("Info message.")
log.Infof("info message line %d.", 1)

// Context.
ctx := log.NewContext()
log.Infofc(ctx, "info message.")
log.Infofc(ctx, "info message line %d.", 100)
```
