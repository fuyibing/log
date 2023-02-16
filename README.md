# Log

```shell
import "github.com/fuyibing/log/v8"
```

## Adapters

- [x] Term
- [x] File
- [x] Kafka

### Configurations

#### Based on config file `config/log.yaml`

```yaml
adapter:
  - file
  - term
```

Based on code

```go

```

## Formatter

Built-in

```
When adapter registered
```

Custom

```go
adapter := log.Client.GetAdapterInterface(conf.Kafka)

if adapter == nil {
return
}

adapter.SetFormatter(func (line *base.Line) (text string, error){

return "format result", nil
})

```