# Log

```
import "github.com/fuyibing/log/v8"
```

```
func main(){
    log.Debug("debug info")
    log.Debugf("debug at MessageQueue[topic=%s, queue=%d]", "Topic", 1)

    // With open tracing on context.
    log.Debugfc(ctx, "debug at MessageQueue[topic=%s, queue=%d]", "Topic", 1)
}
```

## Supports

- [x] Term
- [x] File `async`
- [x] Kafka `async`

### Configurations

##### YAML

Load config file `config/log.yaml` when package initialized.

```
adapter:
  - kafka               # send log to kafka
  - file                # send log to file if send to kafka failed
  - term                # send log to terminal if send to file failed
kafka:
file:
term:
  color: false
```

##### CODE

You must config them with coder.

```
func(){
    log.Config.Kafka.Topic = "Topic"

    log.Config.Set(
        conf.WithLevel(conf.Info),    	
        conf.WithAdapter(conf.Kafka, conf.File, conf.Term),    	
    )
}

```

## Formatter

```log
[2023-02-16 09:10:11.235][DEBUG][PID=3721][SERVICE=app][172.16.0.100:80880] debug message
[2023-02-16 09:10:11.235][INFO][PID=3721][SERVICE=app][172.16.0.100:80880] info message
[2023-02-16 09:10:11.241][WARN][PID=3721][SERVICE=app][172.16.0.100:80880] warning message
[2023-02-16 09:10:11.244][ERROR][PID=3721][SERVICE=app][172.16.0.100:80880] error message
[2023-02-16 09:10:11.246][FATAL][PID=3721][SERVICE=app][172.16.0.100:80880] fatal message
```

### Built-in

- [X] Kafka
- [X] File
- [X] Term

### Custom

```
adapter := log.Client.GetAdapterInterface(conf.Kafka)

if adapter == nil {
    return
}

adapter.SetFormatter(func (line *base.Line) (text string, error){
    return "format result", nil
})

```