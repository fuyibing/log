// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package conf

type (
	KafkaConfiguration interface {
		GetAddresses() []string
		GetTopic() string
	}

	kafkaConfiguration struct {
		parent *configuration

		Addresses []string `yaml:"addresses"`
		Topic     string   `yaml:"topic"`
	}
)

// /////////////////////////////////////////////////////////////
// Interface methods
// /////////////////////////////////////////////////////////////

func (o *kafkaConfiguration) GetAddresses() []string { return o.Addresses }
func (o *kafkaConfiguration) GetTopic() string       { return o.Topic }

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *kafkaConfiguration) initDefaults() {
	if o.Addresses == nil {
		o.Addresses = []string{}
	}
}
