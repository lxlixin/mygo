package logger

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"time"
)

const maxFileSize int64 = 1024 * 1024

//FileLogger 文件日志
type FileLogger struct {
	fileName    string
	errFileName string
	logLevel    int
	logDate     string
	logFile     *os.File
	errFile     *os.File
}

//NewFileLogger 初始化文件日志对象
func NewFileLogger(fileName, errFileName string, logLevel int) *FileLogger {
	ff := &FileLogger{
		fileName:    fileName,
		errFileName: errFileName,
		logLevel:    logLevel,
		logDate:     time.Now().Format("2006-01-02"),
	}
	ff.InitFile()
	return ff
}

//InitFile 初始化文件日志的文件
func (f *FileLogger) InitFile() {
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

//Log 记录信息
func (f *FileLogger) Log(logType int, format string, args ...interface{}) (err error) {
	if logType < f.logLevel {
		return
	}
	logTypeStr := getLogType(logType)
	timeStr := time.Now().Format("2006-01-02 15:04:05.000")
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		err = errors.New("获取报错信息定位失败")
	}
	msg := fmt.Sprintf(format, args...)
	content := fmt.Sprintf("[%s] [%s] [%s] [%s] [%d] %s", timeStr, file, logTypeStr, file, line, msg)
	//将内容写入到file中
	newLogFile, newLogDate0, err0 := logMsg(f.logFile, content, f.fileName, f.logDate)
	if err0 != nil {
		err = err0
		return
	}
	if newLogFile != nil {
		f.logFile = newLogFile
		f.logDate = newLogDate0
	}
	if logType >= Error {
		newErrorFile, _, err1 := logMsg(f.errFile, content, f.errFileName, f.logDate)
		if err1 != nil {
			err = err1
			return
		}
		if newErrorFile != nil {
			f.errFile = newErrorFile
		}
	}
	return
}

func logMsg(logFile *os.File, content, fileName, logDate string) (newFile *os.File, newLogDate string, err error) {
	//判断文件的大小是否达到上限
	fileInfo, _ := logFile.Stat()
	fileSize := fileInfo.Size()
	dateStr := time.Now().Format("2006-01-02")
	if fileSize > maxFileSize || dateStr != logDate {
		//关闭文件
		logFile.Close()
		//复制文件，重新生成文件
		copyFileName := fmt.Sprintf("%s_%s_%d.back", fileName, logDate, time.Now().Unix())
		os.Rename(fileName, copyFileName)
		newFile, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
		if err != nil {
			return
		}
		newFile.Write([]byte(content))
		newLogDate = dateStr
		return
	}
	logFile.Write([]byte(content))
	return
}

// Close 关闭文件
func (f *FileLogger) Close() (err error) {
	err = f.logFile.Close()
	err = f.errFile.Close()
	return
}
