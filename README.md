# Log

Publish application log to adapter with `ASNYC` mode. Support kafka, file, terminal.

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

## Configurations

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

### System

```
log.Client.GetAdapterRegistry().SetFormatter(formatters.NewFileFormatter())
log.Client.GetAdapterRegistry().SetFormatter(formatters.NewJsonFormatter())
```

### Custom

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

## Example

### YAML.

```yaml
# path: config/log.yaml
adapter: term
level: debug
time-format: 15:04:05.9999
async-disabled: true
term:
  color: true
```

### Run example

```shell
cd examples/config-file
go run main.go
```

### Code

```text
func main(){
    defer log.Client.Close()

	log.Debug("debug")
	log.Infof("text format, id=%d, type=%s", 100, "demo")
	log.Map{"id": 100, "type": "demo"}.Infof("with map config")

	c1 := log.NewContext()
	log.Debugfc(c1, "debug")
	log.Infofc(c1, "text format, id=%d, type=%s", 100, "demo")
	log.Map{"id": 100, "type": "demo"}.Infofc(c1, "with map config")

	c2 := log.NewChild(c1)
	log.Infofc(c2, "child of map 1")
	log.Infofc(c2, "child of map 2")
}
```

### Output

```text
[19:16:46.4406][DEBUG][P=65221] debug
[19:16:46.4407][INFO][P=65221] text format, id=100, type=demo
[19:16:46.4407][INFO][P=65221] {"id":100,"type":"demo"} with map config
[19:19:34.3841][DEBUG][P=65221][T=693f34e13932458ba956400d2e921bdc][TS=898c4a1d][TP=][TV=0.0] debug
[19:19:34.3841][INFO][P=65221][T=693f34e13932458ba956400d2e921bdc][TS=898c4a1d][TP=][TV=0.1] text format, id=100, type=demo
[19:19:34.3841][INFO][P=65221][T=693f34e13932458ba956400d2e921bdc][TS=898c4a1d][TP=][TV=0.2] {"id":100,"type":"demo"} with map config
[19:19:34.3841][INFO][P=65221][T=693f34e13932458ba956400d2e921bdc][TS=a2534f36][TP=898c4a1d][TV=0.2.0] child of map 1
[19:19:34.3841][INFO][P=65221][T=693f34e13932458ba956400d2e921bdc][TS=a2534f36][TP=898c4a1d][TV=0.2.1] child of map 2
```


