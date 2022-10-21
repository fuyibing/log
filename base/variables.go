// author: wsfuyibing <websearch@163.com>
// date: 2022-10-15

package base

import "os"

var (
	LogHost       = "0.0.0.0"
	LogName       = "FLOG"
	LogPid        = os.Getpid()
	LogPort       = 0
	LogTimeFormat = "2006-01-02 15:04:05.999"
	LogUserAgent  = "FLOG/3.0"
	LogVersion    = "3.0"
)
