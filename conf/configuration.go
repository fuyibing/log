// author: wsfuyibing <websearch@163.com>
// date: 2023-02-16

package conf

type (
	Configuration struct {
		Adapters   []Adapter `yaml:"adapter"`
		Level      Level     `yaml:"level"`
		TimeFormat string    `yaml:"time-format"`

		debugOn, infoOn, warnOn, errorOn, fatalOn bool
	}
)

// /////////////////////////////////////////////////////////////
// Status switch based on: level field configure
// /////////////////////////////////////////////////////////////

func (o *Configuration) DebugOn() bool { return o.debugOn }
func (o *Configuration) InfoOn() bool  { return o.infoOn }
func (o *Configuration) WarnOn() bool  { return o.warnOn }
func (o *Configuration) ErrorOn() bool { return o.errorOn }
func (o *Configuration) FatalOn() bool { return o.fatalOn }
func (o *Configuration) PanicOn() bool { return o.fatalOn }

func (o *Configuration) Set(options ...Option) {
	for _, option := range options {
		option(o)
	}
}

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *Configuration) init() *Configuration {
	return o
}

func (o *Configuration) updateStatus() {
	li := o.Level.Int()
	on := li > 0

	o.fatalOn = on && li >= Fatal.Int()
	o.errorOn = on && li >= Error.Int()
	o.warnOn = on && li >= Warn.Int()
	o.infoOn = on && li >= Info.Int()
	o.debugOn = on && li >= Debug.Int()
}
