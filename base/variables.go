// author: wsfuyibing <websearch@163.com>
// date: 2022-10-15

package base

import "os"

var (
	LogHost      = "0.0.0.0"   // 计算
	LogPid       = os.Getpid() // 计算
	LogUserAgent = "FLOG/3.0"  // 计算

	LogName       = "FLOG"                    // YAML
	LogPort       = 0                         // YAML
	LogTimeFormat = "2006-01-02 15:04:05.999" // YAML
	LogVersion    = "3.0"                     // YAML
)
