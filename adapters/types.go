// author: wsfuyibing <websearch@163.com>
// date: 2022-10-13

package adapters

// /////////////////////////////////////////////////////////////
// 日志适配
// /////////////////////////////////////////////////////////////

// Adapter
// 适配器类型.
type Adapter string

// 适配器枚举.
const (
	AdapterFile  Adapter = "file"  // 写入TEXT格式到File(本地文件)中
	AdapterKafka Adapter = "kafka" // 发送JSON格式到Kafka
	AdapterRedis Adapter = "redis" // 发送JSON格式到Redis
	AdapterTerm  Adapter = "term"  // 输出TEXT格式到Terminal
)

// /////////////////////////////////////////////////////////////
// 日志级别
// /////////////////////////////////////////////////////////////

type (
	Level    int
	LevelKey string
)

var (
	LevelKeys = map[LevelKey]Level{
		LevelOff:   Off,
		LevelError: Error,
		LevelWarn:  Warn,
		LevelInfo:  Info,
		LevelDebug: Debug,
	}
	LevelText = map[Level]string{
		Debug: "DEBUG",
		Info:  "INFO",
		Warn:  "WARN",
		Error: "ERROR",
	}
)

const (
	Off Level = iota
	Error
	Warn
	Info
	Debug
)

const (
	LevelOff   LevelKey = "off"
	LevelError LevelKey = "error"
	LevelWarn  LevelKey = "warn"
	LevelInfo  LevelKey = "info"
	LevelDebug LevelKey = "debug"
)

func (o Level) String() string {
	if s, ok := LevelText[o]; ok {
		return s
	}
	return "UNKNOWN"
}

// /////////////////////////////////////////////////////////////
// 执行器
// /////////////////////////////////////////////////////////////

type (
	Handler func(line *Line, err error)
)
