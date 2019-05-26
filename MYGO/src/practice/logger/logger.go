package logger

//日志级别
const (
	_ = iota
	Debug
	Info
	Warning
	Error
	Fatal
)

func getLogType(logType int) string {
	switch logType {
	case 1:
		return "debug"
	case 2:
		return "info"
	case 3:
		return "warning"
	case 4:
		return "error"
	case 5:
		return "fatal"
	default:
		return "debug"
	}

}

//Logger 日志调用接口
type Logger interface {
	Log(logType int, format string, args ...interface{}) error
	Close() error
	Start() error
}
