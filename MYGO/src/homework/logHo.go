package homework

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"syscall"
	"time"
)

//日志文件大小
const maxFileSize int64 = 10

//日志文件中的日期格式
const timeFormat = "2006-01-02 15:04:05.000"

//日志文件名中的日期格式
const copyFileTimeFormat = "20060102_150405"

const (
	_ = iota
	dabug
	info
	warn
	errer
	fatal
)

func init() {
	// 解决permission denied 问题
	// https://zhuanlan.zhihu.com/p/33692995
	mask := syscall.Umask(0)
	defer syscall.Umask(mask)
	err := os.MkdirAll("mygolang/log", os.ModePerm)
	check(err)
}

// Log record log msg
// Fname 调用方法的文件名
type Log struct {
	Fname string
}

// Debug msg
func (l Log) Debug(msg string) {
	l.commonLog(dabug, msg)
}

// Info msg
func (l Log) Info(msg string) {
	l.commonLog(info, msg)
}

// Warn msg
func (l Log) Warn(msg string) {
	l.commonLog(warn, msg)
}

// Error msg
func (l Log) Error(msg string) {
	l.commonLog(errer, msg)
}

// Fatal msg
func (l Log) Fatal(msg string) {
	l.commonLog(fatal, msg)
}

// 检查错误
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// 获取文件大小
func getFilesize(fileName string) int64 {
	var result int64
	filepath.Walk(fileName, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	return result
}

// 判断文件是否存在
func isFileExist(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	}
	return true
}

// 复制文件内容到指定文件
func copyFile(fromFilename, toFilename string) {
	originFile, err := os.Open(fromFilename)
	check(err)
	defer originFile.Close()

	newFile, err := os.Create(toFilename)
	check(err)
	defer newFile.Close()

	bytesWritten, err := io.Copy(newFile, originFile)
	check(err)
	fmt.Println(bytesWritten)
	// 清空文件
	originFile.Truncate(0)
}

func (l Log) commonLog(bugType int, msg string) {
	var fileName string
	var copyFilename string
	var logTypeStr string
	switch bugType {
	case dabug:
		fileName = "mygolang/log/debug.log"
		copyFilename = fmt.Sprintf("mygolang/log/debug_%s.log", time.Now().Format(copyFileTimeFormat))
		logTypeStr = "bug"
		break
	case info:
		fileName = "mygolang/log/info.log"
		copyFilename = fmt.Sprintf("mygolang/log/info_%s.log", time.Now().Format(copyFileTimeFormat))
		logTypeStr = "info"
		break
	case errer:
		fileName = "mygolang/log/errer.log"
		copyFilename = fmt.Sprintf("mygolang/log/errer_%s.log", time.Now().Format(copyFileTimeFormat))
		logTypeStr = "errer"
		break
	case warn:
		fileName = "mygolang/log/warn.log"
		copyFilename = fmt.Sprintf("mygolang/log/warn_%s.log", time.Now().Format(copyFileTimeFormat))
		logTypeStr = "warn"
		break
	case fatal:
		fileName = "mygolang/log/fatal.log"
		copyFilename = fmt.Sprintf("mygolang/log/fatal_%s.log", time.Now().Format(copyFileTimeFormat))
		logTypeStr = "fatal"
		break
	}

	//获取当前debug的文件名是否存在
	if !isFileExist(fileName) {
		newfile, err := os.Create(fileName)
		check(err)
		defer newfile.Close()
	} else {
		//获取文件大小
		fileSize := getFilesize(fileName)
		if fileSize >= maxFileSize {
			//拷贝文件到备份文件
			copyFile(fileName, copyFilename)
		}
	}

	// 拼写要写入的内容
	timeStr := time.Now().Format(timeFormat)
	_, msgFile, line, _ := runtime.Caller(1)
	msgFile = filepath.Base(msgFile)
	content := fmt.Sprintf("%s [%s] %s %s:[%d] %s \n", timeStr, l.Fname, logTypeStr, msgFile, line, msg)

	// 写入到文件中
	fd, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	check(err)
	fd.Write([]byte(content))
	fd.Close()
}
