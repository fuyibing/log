# Log

Publish application log to target storage with `ASNYC` mode.

```
import "github.com/fuyibing/log/v8"
```

```
func init(){
    log.Config.Set(
        conf.SetAdapter("term"),
        conf.SetLevel(conf.Debug),
        conf.SetTermColor(true),
    )
}

func main(){
    defer func(){
        // Wait all uncompleted logs done.
        log.Client.Stop()
    }()

    log.Debug("debug info")
    log.Debugf("debug at MessageQueue[topic=%s, queue=%d]", "Topic", 1)

    // With open tracing on context.
    log.Debugfc(ctx, "debug at MessageQueue[topic=%s, queue=%d]", "Topic", 1)
}
```

## ASync Supports

- [X] `Term` - Print log content on console.
- [X] `File` - Write log to local file.
- [X] `Kafka` - Publish log to kafka.
- [ ] `SLS` - Aliyunn SLS service.

### Configurations

##### YAML

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

##### CODE

You must config them with coder.

```
func init() {
    log.Config.Set(	
        conf.SetAdapter(adapters.AdapterTerm),
        conf.SetLevel(conf.Info),
        conf.SetTermColor(true),
    )
}

```

## Formatter

```log
[2023-02-16 09:10:11.235][DEBUG][PID=3721] debug message
[2023-02-16 09:10:11.235][INFO][PID=3721] info message
[2023-02-16 09:10:11.241][WARN][PID=3721] warning message
[2023-02-16 09:10:11.244][ERROR][PID=3721] error message
[2023-02-16 09:10:11.246][FATAL][PID=3721] fatal message
```

##### Custom

```
func init(){
    log.Client.GetAdapterRegistry().SetFormatter(&formatter{})
}
```