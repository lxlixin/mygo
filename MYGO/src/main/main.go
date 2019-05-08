package main

import "homework"

func main() {
	mylog := homework.Log{Fname: "main"}
	mylog.Debug("这是一个bug信息！难道是覆盖写入么")
}
