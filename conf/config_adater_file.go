// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

package conf

type (
	FileConfiguration interface {
		GetBasePath() string
		GetExtName() string
		GetFileName() string
		GetSeparatorPath() string
	}

	fileConfiguration struct {
		parent *configuration

		BasePath      string `yaml:"base-path"`
		ExtName       string `yaml:"ext-name"`
		FileName      string `yaml:"file-name"`
		SeparatorPath string `yaml:"separator-path"`
	}
)

// /////////////////////////////////////////////////////////////
// Interface methods
// /////////////////////////////////////////////////////////////

func (o *fileConfiguration) GetBasePath() string      { return o.BasePath }
func (o *fileConfiguration) GetExtName() string       { return o.ExtName }
func (o *fileConfiguration) GetFileName() string      { return o.FileName }
func (o *fileConfiguration) GetSeparatorPath() string { return o.SeparatorPath }

// /////////////////////////////////////////////////////////////
// Access methods
// /////////////////////////////////////////////////////////////

func (o *fileConfiguration) initDefaults() {
	if o.BasePath == "" {
		o.parent.Set(SetFileBasePath(DefaultFileBasePath))
	} else {
		o.parent.Set(SetFileBasePath(o.BasePath))
	}
	if o.ExtName == "" {
		o.parent.Set(SetFileExtName(DefaultFileExtName))
	}
	if o.FileName == "" {
		o.parent.Set(SetFileFileName(DefaultFileFileName))
	}
	if o.SeparatorPath == "" {
		o.parent.Set(SetFileSeparatorPath(DefaultFileSeparatorPath))
	}
}
