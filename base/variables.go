// author: wsfuyibing <websearch@163.com>
// date: 2022-10-15

package base

import "os"

var (
	LogHost = "0.0.0.0"
	LogPid  = os.Getpid()
	LogPort = 0

	LogName       = "FLOG"
	LogUserAgent  = "FLOG/3.0"
	LogVersion    = "3.0"
	LogTimeFormat = "2006-01-02 15:04:05.999"
)
