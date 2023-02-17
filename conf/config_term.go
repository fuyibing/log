// author: wsfuyibing <websearch@163.com>
// date: 2023-02-17

package conf

type (
	TermConfig struct {
		Color bool `xml:"color"`
	}
)

func (o *TermConfig) defaults() {}
