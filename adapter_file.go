// author: wsfuyibing <websearch@163.com>
// date: 2021-02-10

package log

// File struct.
type adapterFile struct{}

// New file adapter instance.
func NewAdapterFile() *adapterFile {
	return new(adapterFile)
}

// File adapter callback.
func (o *adapterFile) Callback(line LineInterface) {}
