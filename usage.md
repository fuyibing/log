# 用法

> 日志共分五个级别，分别为 `DEBUG`, `INFO`, `WARN`, `ERROR`, `FATAL`，
> 当使用 `FATAL` 级别时自动加载 `Stack` 堆栈。

### 一、简易日志

```go
log.Info("message")
log.Info("message: %d", 1)
log.Fatal("message")
```

```text
[2023-03-01 09:10:11.123460][ INFO] message
[2023-03-01 09:10:11.123463][ INFO] message 1
[2023-03-01 09:10:11.123470][FATAL] message
</Users/fuyibing/codes/github.com/fuyibing/log/examples/demo/main.go:84> IN <main.main3()>
</Users/fuyibing/codes/github.com/fuyibing/log/examples/demo/main.go:59> IN <main.main()>
```

### 二、字段日志

### 三、链路日志
