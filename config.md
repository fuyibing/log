# 配置

> 包初始化时，自动从配置文件(config/log.yaml)文件中读取配置参数，若未指定则使用默认值

### 异步批处理

```yaml
bucket-batch: 100                       # 每批次最大数量
bucket-capacity: 10000                  # 内存队列最大容量
bucket-concurrency: 5                   # 最大并发/并行上报
bucket-frequency: 500                   # 自动上报频率(单位: 毫秒)
```

### 更多适配项

1. [Logger](./config.logger.md) - 上报日志
2. [Tracer](./config.tracer.md) - 上报链路
