# logger

1. [x] 包名: `github.com/fuyibing/log/v2`.
1. [x] 版本: `v2`

> 本包在导入后, 扫描 `tmp/log.yaml` 或 `config/log.yaml` 并初始化,
> 详细配置参见本包的 `config/log.yaml` 注释.

### 通用用法.

```
log.Info("message.")
log.Infof("message %d: %s.", 1, "parse", "example")
```

### 带请求链

```
ctx := log.NewContext()
log.Debugfc(ctx, "debug fc")
log.Infofc(ctx, "info fc")
log.Warnfc(ctx, "warn fc")
log.Errorfc(ctx, "error fc")
```

### 在IRIS框架中

> 首先通过中间件, 在入口注册请求链.

```text
log.BindRequest(iris.Request())
```

> 在IRIS框架中使用.

```text
ctx := iris.Request().Context()

log.Debugfc(ctx, "debug fc")
log.Infofc(ctx, "info fc")
log.Warnfc(ctx, "warn fc")
log.Errorfc(ctx, "error fc")

```

### 自定义日志处理.

```text
log.Config.SetHandler(func(line interfaces.LineInterface) {
    println("handler: ", line.SpanVersion(), line.Content())
})

ctx := log.NewContext()
log.Debugfc(ctx, "debug fc")
log.Infofc(ctx, "info fc")
log.Warnfc(ctx, "warn fc")
log.Errorfc(ctx, "error fc")
```

