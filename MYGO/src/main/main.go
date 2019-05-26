package main

import "practice/logger"

func main() {

	l := logger.NewConcurrentFileLogger("./cLog.log", "./cErrLog.log", logger.Debug)
	i := 0
	go l.Start()
	for {
		i++
		l.Log(logger.Debug, "这是第%d条日志", i)
	}

}
