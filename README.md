# Logger

> package `github.com/fuyibing/log/v3`.

### Example.

```text
// Simple.
log.Info("Info message.")
log.Infof("info message line %d.", 1)

// Context.
ctx := log.NewContext()
log.Infofc(ctx, "info message.")
log.Infofc(ctx, "info message line %d.", 100)
```

