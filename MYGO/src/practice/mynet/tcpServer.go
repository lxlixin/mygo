package mynet

import (
	"bufio"
	"fmt"
	"net"
)

var ln net.Listener

//StartServer 开启器服务器
func StartServer() {
	var err error
	// tcp  可以换成udp
	ln, err = net.Listen("tcp", ":9988")
	if err != nil {
		fmt.Println("启动服务器失败！")
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	msg, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("读取信息失败", err)
		return
	}
	fmt.Printf("来自客户端的消息：%s", msg)
	conn.Write([]byte("收到你的消息了！你也好呀！\n"))
}

// StopServer 停止服务器
func StopServer() {
	ln.Close()
}
