// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package kafka

import (
	"github.com/Shopify/sarama"
	_ "github.com/Shopify/sarama"
	"github.com/fuyibing/log/v8/base"
	"github.com/fuyibing/log/v8/conf"
	"github.com/fuyibing/log/v8/formatters"
)

type (
	// Executor
	// send log lines to kafka with specified topic.
	//
	// You can subscribe kafka message than redirect to aliyun
	// log service or elasticsearch by logstash.
	Executor struct {
		formatter     formatters.Formatter
		producer      sarama.SyncProducer
		producerError error
	}
)

func New() *Executor {
	return (&Executor{}).init()
}

// /////////////////////////////////////////////////////////////
// Exported methods
// /////////////////////////////////////////////////////////////

func (o *Executor) Logs(lines ...*base.Line) error {
	if o.producerError != nil {
		return o.producerError
	}

	// Generate producer messages.
	list := make([]*sarama.ProducerMessage, 0)
	for _, line := range lines {
		list = append(list, &sarama.ProducerMessage{
			Topic: conf.Config.GetKafka().GetTopic(),
			Value: sarama.StringEncoder(o.formatter.String(line)),
		})
	}

	// Send to kafka.
	return o.producer.SendMessages(list)
}

func (o *Executor) SetFormatter(formatter formatters.Formatter) {
	o.formatter = formatter
}

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *Executor) init() *Executor {
	o.initProducer()
	return o
}

func (o *Executor) initProducer() {
	cfg := sarama.NewConfig()
	cfg.ClientID = conf.Agent
	cfg.Producer.RequiredAcks = sarama.NoResponse
	cfg.Producer.Return.Errors = true
	cfg.Producer.Return.Successes = true

	o.producer, o.producerError = sarama.NewSyncProducer(
		conf.Config.GetKafka().GetAddresses(),
		cfg,
	)
}
