package adapters

import (
	"context"
	"encoding/json"
	"github.com/fuyibing/log/v2/interfaces"
	"github.com/segmentio/kafka-go"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"strings"
)

type kafkaConfig struct {
	Brokers string `yaml:"brokers"`
	Topic   string `yaml:"topic"`
}

type kafkaAdapter struct {
	Conf    *kafkaConfig `yaml:"kafka"`
	ch      chan interfaces.LineInterface
	writer  *kafka.Writer
	handler interfaces.Handler
}

func (o *kafkaAdapter) Run(line interfaces.LineInterface) {
	go func() {
		o.ch <- line
	}()
}

func (o *kafkaAdapter) body(line interfaces.LineInterface) []byte {
	// Init
	data := make(map[string]interface{})
	// Basic.
	data["content"] = line.Content()
	data["duration"] = line.Duration()
	data["level"] = line.Level()
	data["time"] = line.Timeline()
	// Tracing.
	data["action"] = ""
	if line.Tracing() {
		data["parentSpanId"] = line.ParentSpanId()
		data["requestId"] = line.TraceId()
		data["requestMethod"], data["requestUrl"] = line.RequestInfo()
		data["spanId"] = line.SpanId()
		data["traceId"] = line.TraceId()
		data["version"] = line.SpanVersion()
	} else {
		data["parentSpanId"] = ""
		data["requestId"] = ""
		data["requestMethod"] = ""
		data["requestUrl"] = ""
		data["spanId"] = ""
		data["traceId"] = ""
		data["version"] = ""
	}
	// Server.
	data["module"] = line.ServiceName()
	data["pid"] = line.Pid()
	data["serverAddr"] = line.ServiceAddr()
	data["taskId"] = 0
	data["taskName"] = ""
	// JSON string.
	if body, err := json.Marshal(data); err == nil {
		return body
	}
	return nil
}

func (o *kafkaAdapter) listen() {
	go func() {
		defer o.listen()
		for {
			select {
			case line := <-o.ch:
				go o.send(line)
			}
		}
	}()
}

func (o *kafkaAdapter) send(line interfaces.LineInterface) {
	// Catch panic.
	defer func() {
		if r := recover(); r != nil {
			o.handler(line)
		}
	}()
	// Write message
	data := o.body(line)
	if data != nil {
		err := o.writer.WriteMessages(context.Background(),
			kafka.Message{
				//Topic: kw.topic,
				Value: data,
			},
		)
		if err != nil {
			o.handler(line)
		}
	} else {
		o.handler(line)
	}
}

// NewKafka new kafka adapter
func NewKafka() *kafkaAdapter {
	o := &kafkaAdapter{ch: make(chan interfaces.LineInterface), handler: NewTerm().Run}
	// Scan config files and fill kafka config
	for _, file := range []string{"./tmp/log.yaml", "../tmp/log.yaml", "./config/log.yaml", "../config/log.yaml"} {
		body, err := ioutil.ReadFile(file)
		if err != nil {
			continue
		}
		if yaml.Unmarshal(body, o) != nil {
			continue
		}
		break
	}
	// create kafka writer instance
	brokerArr := strings.Split(o.Conf.Brokers, ",")
	kp := &kafka.Writer{
		Addr:     kafka.TCP(brokerArr...),
		Topic:    o.Conf.Topic,
		Balancer: &kafka.LeastBytes{},
		Async:    true,
	}
	o.writer = kp
	//
	o.listen()
	return o
}
