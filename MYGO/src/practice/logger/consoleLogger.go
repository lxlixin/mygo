package logger

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
)

//ConsoleLogger 控制logger
type ConsoleLogger struct {
	logLevel int
}

// NewConsoleLogger 初始化控制台logger
func NewConsoleLogger(logLevel int) ConsoleLogger {
	return ConsoleLogger{
		logLevel: logLevel,
	}
}

//Log 记录日志信息
func (c ConsoleLogger) Log(logType int, format string, args ...interface{}) (err error) {
	if logType < c.logLevel {
		return
	}
	logTypeStr := getLogType(logType)
	timeStr := time.Now().Format("2006-01-02 15:04:05.000")
	pc, fileName, line, ok := runtime.Caller(1)
	// 从fileName（文件全路径）中剥离出文件名
	fileName = path.Base(fileName) // x/y/xx.txt
	// 根据pc拿到函数名
	funcName := runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName)
	if !ok {
		err = errors.New("获取报错信息定位失败")
	}
	msg := fmt.Sprintf(format, args...)
	content := fmt.Sprintf("[%s] [%s] [%s] [%s] [%d] %s", timeStr, fileName, logTypeStr, funcName, line, msg)
	// 利用fmt包将msg字符串写入f.file文件中
	fmt.Fprintln(os.Stdout, content)
	return
}

//Close 关闭
func (c ConsoleLogger) Close() (err error) {
	return
}

//Start 关闭
func (c ConsoleLogger) Start() (err error) {
	return
}
