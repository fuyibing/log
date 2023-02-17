// author: wsfuyibing <websearch@163.com>
// date: 2023-02-17

package conf

func (o *Configuration) DebugOn() bool { return o.debugOn }
func (o *Configuration) InfoOn() bool  { return o.infoOn }
func (o *Configuration) WarnOn() bool  { return o.warnOn }
func (o *Configuration) ErrorOn() bool { return o.errorOn }
func (o *Configuration) FatalOn() bool { return o.fatalOn }
func (o *Configuration) PanicOn() bool { return o.fatalOn }

func (o *Configuration) updateStatus() {
	li := o.Level.Int()
	on := li > 0

	o.fatalOn = on && li >= Fatal.Int()
	o.errorOn = on && li >= Error.Int()
	o.warnOn = on && li >= Warn.Int()
	o.infoOn = on && li >= Info.Int()
	o.debugOn = on && li >= Debug.Int()
}
