package httpconfig

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func init()  {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/getDetail", getDetail)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("start server fail !")
	}
}

func homeHandler(w http.ResponseWriter, h *http.Request) {
	w.Write([]byte("hello world!"))
}

func getDetail(w http.ResponseWriter, h *http.Request) {
	resp, err := http.Get("http://39.107.60.231:9020/traffic/recruitPlan/detail?recruitPlanId=122373252948721664")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	// 一次性读取
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("错误")
		return
	}
	fmt.Println(string(bs))
}
