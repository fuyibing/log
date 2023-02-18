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

- `Term` : Print log content on console.
- `File` : Write log to local file with **async** mode.
- `Kafka` : Publish log to kafka with **async** mode. You can subscribe specified topic message by **logstash**, then
  write to aliyun log service or elasticsearch.

### Configurations

##### YAML

Load config file `config/log.yaml` when package initialized.

```
adapter: term
level: debug
time-format: "2006-01-02 15:04:05.999999"

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