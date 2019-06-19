package mynet

import (
	"net/http"
)

//StartHttpServer  启动服务
func StartHttpServer() {
	http.ListenAndServe("127.0.0.1:8888", nil)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

	})
}
