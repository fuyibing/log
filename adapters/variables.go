// author: wsfuyibing <websearch@163.com>
// date: 2022-10-13

package adapters

var (
	Host       = "0.0.0.0"                    // 部署IP(计算: 从网卡上解析)
	Name       = "flog"                       // 服务名(解析: 从 app.yaml 中解析)
	Pid        = 0                            // 进程ID(计算: 服务启动时的进程ID)
	Port       = 0                            // 监听端口号(解析: 从 app.yaml 中解析)
	Software   = "flog/3.0"                   // 应用名称(解析: 从 app.yaml 中解析)
	TimeFormat = "2006-01-02 15:04:05.999999" // 日志时间格式(解析: 从 log.yaml 中解析)
	Version    = "3.0"                        // 版本号(解析: 从 app.yaml 中解析)

	NodeId string // 节点ID
)
