# Log

Publish application log to target storage with `ASNYC` mode.

```
import "github.com/fuyibing/log/v8"
```

```
func init(){
    // Follow configurations are optional. All of follows can be
    // configured in `config/log.yaml`. Filled with default if not
    // configured.
    log.Config.Set(
        conf.SetTimeFormat("2006-01-02 15:04:05.999999"),
        conf.SetLevel(conf.Error),
        conf.SetPrefix("Prefix"),

        conf.SetServiceAddr("172.16.0.110"),
        conf.SetServicePort(8080),
        conf.SetServiceEnvironment("production"),
        conf.SetServiceName("myapp"),
        conf.SetServiceVersion("1.2.3"),
    )

    // If adapter changed by code, You must call log.Client.Reset()
    // to apply it.
    log.Config.Set(conf.SetAdapter(adapters.AdapterKafka))
    log.Client.Reset()
}

func main(){
    // Wait for a while
    // until all logs publish completed.
    //
    // If the Close method `log.Client.Close()` is not set, Some logs
    // end of the application may be lost.
    defer log.Client.Close()

    // ... ...
    
    log.Debug("debug info")
    log.Infof("info message: adapter=%s, level=%v", conf.Adapter, conf.Level)
    
    // ... ...
}
```

## ASync Supports

- [X] `Term` - Print log content on console.
- [X] `File` - Write log to local file.
- [X] `Kafka` - Publish log to kafka.
- [ ] `SLS` - Aliyun SLS service.

### Configurations

Load config file `config/log.yaml` when package initialized. Use default if not specified.

```
adapter: term
level: debug
time-format: "2006-01-02 15:04:05.999999"
kafka:
  topic: "Log"
  addresses: 
    - "172.16.0.100:9092"
    - "172.16.0.101:9092"
    - "172.16.0.102:9092"
file:
term:
  color: false
```

## Formatter

```log
[2023-02-16 09:10:11.235][DEBUG][PID=3721] debug message
[2023-02-16 09:10:11.235][INFO][PID=3721] info message
[2023-02-16 09:10:11.241][WARN][PID=3721] warning message
[2023-02-16 09:10:11.244][ERROR][PID=3721] error message
[2023-02-16 09:10:11.246][FATAL][PID=3721] fatal message
```

System formatter

```
log.Client.GetAdapterRegistry().SetFormatter(formatters.NewFileFormatter())
log.Client.GetAdapterRegistry().SetFormatter(formatters.NewJsonFormatter())
```

Customer formatter

```
type MyFormatter struct{}

func (o *MyFormatter) Body(line *base.Line) []byte {
    // ...
}

func (o *MyFormatter) String(line *base.Line) string {
    // ...
}

func init(){
    log.Client.GetAdapterRegistry().SetFormatter(&MyFormatter{})
}
```