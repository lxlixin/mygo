package main

import "practice/logger"

func main() {
	var myLogger logger.Logger = logger.NewConsoleLogger(logger.Error)
	defer myLogger.Close()
	myLogger.Log(logger.Fatal, "这是日志内容：%s", "日志内容就是八点佛塑科技")
}
