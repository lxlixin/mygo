package mynet

import (
	"bufio"
	"fmt"
	"net"
)

//SendMsg 往服务器发送消息
func SendMsg() {
	// tcp  可以换成udp
	conn, err := net.Dial("tcp", "127.0.0.1:9988")
	if err != nil {
		fmt.Println("链接服务器失败", err)
		return
	}
	defer conn.Close()

	conn.Write([]byte("你好呀，服务器 \n"))

	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("读取信息失败")
	}
	fmt.Printf("来自服务器端的消息：%s", status)
}
