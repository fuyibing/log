// author: wsfuyibing <websearch@163.com>
// date: 2023-02-17

package conf

type (
	Adapter string
)

const (
	Term  Adapter = "term"
	File  Adapter = "file"
	Kafka Adapter = "kafka"
)
