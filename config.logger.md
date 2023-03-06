# Logger

1. [X] `Term` - 打印到终端/控制台
2. [X] `File` - 输出到文件中
3. [ ] `Kafka` - 发布到Kafka

### 公共

```yaml
logger-level: info
```

### 适配

##### File

> `异步/ASync` 日志写入到文件中

```yaml
logger-exporter: file         # 必须
file-logger:
  path: /var/logs             # 存储目录
  folder: 2006-01             # 目录分隔(基于时间)
  name: 2006-01-02            # 日志文件名(基于时间)
  ext: log                    # 日志扩展名
```

##### Term

> 日志打印到终端/控制台, 此模式适合于开发环境, 且此模式是日志是同步打印.

```yaml
logger-exporter: term         # 必须
```
