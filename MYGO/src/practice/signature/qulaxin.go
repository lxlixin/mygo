package signature

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//SendToQulaxin 向趣拉新请求数据
func SendToQulaxin() {

	data := make(map[string]interface{})
	data["code"] = "test1"
	data["extend"] = "144240"
	data["id"] = "47"

	inputs := make([]map[string]string, 4)
	input0 := make(map[string]string)
	input0["label"] = "姓名"
	input0["type"] = "text"
	input0["value"] = "name"
	inputs[0] = input0

	input1 := make(map[string]string)
	input1["label"] = "图片1"
	input1["type"] = "image"
	input1["value"] = "http://static.appshike.com/static/images/supertask/17057758/20190110110351/8f2f98c18d3141aabbf1d7c809e4f7f1.png"
	inputs[1] = input1

	input2 := make(map[string]string)
	input2["label"] = "图片2"
	input2["type"] = "image"
	input2["value"] = "http://static.appshike.com/static/images/supertask/17057758/20190110110351/8f2f98c18d3141aabbf1d7c809e4f7f1.png"
	inputs[2] = input2

	input3 := make(map[string]string)
	input3["label"] = "手机号"
	input3["type"] = "text"
	input3["value"] = "18910003832"
	inputs[3] = input3

	data["input"] = inputs

	dataList := make([]map[string]interface{}, 1)
	dataList[0] = data

	jsonStr, err := json.Marshal(dataList)
	if err != nil {
		fmt.Printf("json 序列化失败，error 是：%v\n", err)
	}

	request := make(map[string]interface{}, 5)
	request["appKey"] = "duodianguanggao"
	request["cmd"] = "userSubmit"
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	fmt.Printf("timestamp is : %v \n", timestamp)
	request["timestamp"] = timestamp
	request["dataList"] = dataList

	stringA := "202cb962ac59075b964b07152d234b70" + "appKey=" + request["appKey"].(string) + "cmd=" + request["cmd"].(string) + "dataList=" + string(jsonStr) + "timestamp=" + request["timestamp"].(string) + "202cb962ac59075b964b07152d234b70"

	//获取MD5加密后的字符串
	sign := strings.ToUpper(MD5(stringA))

	request["sign"] = sign

	requestStr, err := json.Marshal(request)

	if err != nil {
		fmt.Printf("request json序列化报错了：%v\n", err)
	}

	// fmt.Printf("value is : %v \n", string(requestStr))
	// fmt.Printf("jsonString : %s\n", stringA)

	requestS := bytes.NewBuffer(requestStr)
	r, _ := http.NewRequest("POST", "http://openapi.qulaxin.com/api/channel", requestS)
	r.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	response, _ := client.Do(r)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	}
}

//SendToDuodian  模拟回调多点接口
func SendToDuodian() {
	data := make(map[string]interface{})
	data["extend"] = "414353"
	data["checkTime"] = "2019-01-01 11:11:11"
	data["id"] = "51"
	data["checkResult"] = 1
	data["checkRemark"] = "暂且假装不通过吧"

	dataList := make([]map[string]interface{}, 1)
	dataList[0] = data

	jsonStr, err := json.Marshal(dataList)
	if err != nil {
		fmt.Printf("json 序列化失败，error 是：%v\n", err)
	}

	request := make(map[string]interface{}, 5)
	request["appKey"] = "duodianguanggao"
	request["cmd"] = "userSubmit"
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	fmt.Printf("timestamp is : %v \n", timestamp)
	request["timestamp"] = timestamp
	request["dataList"] = string(jsonStr)

	stringA := "202cb962ac59075b964b07152d234b70" + "appKey=" + request["appKey"].(string) + "cmd=" + request["cmd"].(string) + "dataList=" + request["dataList"].(string) + "timestamp=" + request["timestamp"].(string) + "202cb962ac59075b964b07152d234b70"

	//获取MD5加密后的字符串
	sign := strings.ToUpper(MD5(stringA))

	request["sign"] = sign

	requestStr, err := json.Marshal(request)

	if err != nil {
		fmt.Printf("request json序列化报错了：%v\n", err)
	}

	fmt.Printf("value is : %v \n", string(requestStr))
	fmt.Printf("jsonString : %s\n", stringA)

	requestS := bytes.NewBuffer(requestStr)
	r, _ := http.NewRequest("POST", "http://localhost:8089/shike/api/android/cpaSettlement/qulaxin", requestS)
	r.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	response, _ := client.Do(r)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	} else {
		fmt.Println("返回的状态码", response.StatusCode)
	}
}

//MD5 生成32位MD5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}
