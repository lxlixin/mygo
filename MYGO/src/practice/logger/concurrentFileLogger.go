package logger

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"sync"
	"time"
)

//LoggerContent 日志内容实体
type LoggerContent struct {
	logType  int
	content  string
	logTime  string
	funcName string
	line     int
}

//newLoggerContent 初始化日志内容
func newLoggerContent(logType int, format string, args ...interface{}) (*LoggerContent, error) {
	timeStr := time.Now().Format("2006-01-02 15:04:05.000")
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		err := errors.New("获取报错信息定位失败")
		return nil, err
	}
	file = path.Base(file)
	msg := fmt.Sprintf(format, args...)
	return &LoggerContent{
		logType:  logType,
		content:  msg,
		logTime:  timeStr,
		funcName: file,
		line:     line,
	}, nil
}

//ConcurrentFileLogger 文件日志
type ConcurrentFileLogger struct {
	fileName    string
	errFileName string
	logLevel    int
	logDate     string
	logFile     *os.File
	errFile     *os.File
	logBacket   chan *LoggerContent
	lock        *sync.Mutex
}

//NewConcurrentFileLogger  初始化文件日志对象
func NewConcurrentFileLogger(fileName, errFileName string, logLevel int) *ConcurrentFileLogger {
	ff := &ConcurrentFileLogger{
		fileName:    fileName,
		errFileName: errFileName,
		logLevel:    logLevel,
		logDate:     time.Now().Format("2006-01-02"),
		logBacket:   make(chan *LoggerContent, 50000),
		lock:        &sync.Mutex{},
	}
	ff.InitFile()
	return ff
}

//InitFile 初始化文件日志的文件
func (f *ConcurrentFileLogger) InitFile() {
	logFile, err := os.OpenFile(f.fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		panic("生成或者打开日志文件失败！")
	}
	f.logFile = logFile

	errFile, err := os.OpenFile(f.errFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		panic("生成或者打开错误日志文件失败！")
	}
	f.errFile = errFile
}

//Log 记录日志信息
func (f *ConcurrentFileLogger) Log(logType int, format string, args ...interface{}) (err error) {
	if len(f.logBacket) >= 50000 {
		return errors.New("缓冲日志超过50000，此条日志被丢弃！")
	}
	c, err := newLoggerContent(logType, format, args...)
	if err != nil {
		return
	}
	f.logBacket <- c
	return
}

func (f *ConcurrentFileLogger) realLog() (err error) {
	loggerContent := <-f.logBacket
	if loggerContent.logType < f.logLevel {
		return
	}
	logTypeStr := getLogType(loggerContent.logType)
	content := fmt.Sprintf("[%s] [%s] [%s] [%d] %s \n", loggerContent.logTime, loggerContent.funcName, logTypeStr, loggerContent.line, loggerContent.content)
	//将内容写入到file中

	f.lock.Lock()
	newLogFile, newLogDate0, err0 := logMsg(f.logFile, content, f.fileName, f.logDate)
	if err0 != nil {
		err = err0
		return
	}
	if newLogFile != nil {
		f.logFile = newLogFile
		f.logDate = newLogDate0
	}
	f.lock.Unlock()
	if loggerContent.logType >= Error {
		f.lock.Lock()
		newErrorFile, _, err1 := logMsg(f.errFile, content, f.errFileName, f.logDate)
		if err1 != nil {
			err = err1
			return
		}
		if newErrorFile != nil {
			f.errFile = newErrorFile
		}
		f.lock.Unlock()
	}
	return
}

// Close 关闭文件
func (f *ConcurrentFileLogger) Close() (err error) {
	err = f.logFile.Close()
	err = f.errFile.Close()
	return
}

//Start 关闭
func (f *ConcurrentFileLogger) Start() (err error) {
	for {
		if len(f.logBacket) > 0 {
			f.realLog()
		}
	}
}

// func main() {

// 	l := logger.NewConcurrentFileLogger("./cLog.log", "./cErrLog.log", logger.Debug)
// 	i := 0
// 	启动一个专门的goroutine 跑线程
// 	go l.Start()
// 	for {
// 		i++
// 		l.Log(logger.Debug, "这是第%d条日志", i)
// 	}

// }
