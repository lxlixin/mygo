package homework

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// MySQLConfig mysql配置信息实体
type MySQLConfig struct {
	Addr     string `ini:"addr"`
	UserName string `ini:"username"`
	Password string `ini:"password"`
	Port     int64  `ini:"port"`
}

// RedisConfig redis 配置信息实体
type RedisConfig struct {
	Host string `ini:"host"`
	Port int64  `ini:"port"`
}

// InitDataConfig 从指定的路径下读取文件信息到配置类中
func InitDataConfig(fPath string) {
	f, err := os.Open(fPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var currentStruct string
	mysql := MySQLConfig{}
	redis := RedisConfig{}
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		// fmt.Println(line)
		sName, content, fild, ok := getContent(line)
		if ok {
			if sName != "" {
				currentStruct = sName
				continue
			}
			switch currentStruct {
			case "[mysql]":
				v := reflect.ValueOf(&mysql)
				t := reflect.TypeOf(mysql)
				setValue(v, t, fild, content)
			case "[redis]":
				v := reflect.ValueOf(&redis)
				t := reflect.TypeOf(redis)
				setValue(v, t, fild, content)
			}
		}
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
	}
	fmt.Printf("%v\n", mysql)
	fmt.Printf("%v\n", redis)
}

func getContent(v string) (sName, content, fild string, ok bool) {
	if v == "" {
		return
	}
	//匹配[%]
	reg := regexp.MustCompile(`^\[.*\]`)
	result := reg.FindAllString(v, -1)
	if len(result) > 0 {
		sName = result[0]
		ok = true
		return
	}
	//判断如果是以# 或者;开头，则认为是注释
	exist := strings.HasPrefix(v, "#")
	if exist {
		return
	}
	exist = strings.HasPrefix(v, ";")
	if exist {
		return
	}
	//匹配 x=模式 并且返回的长度大于1，则说明格式满足条件
	reg = regexp.MustCompile(`[^=]*(?U)=+[^=]*\b`)
	result = reg.FindAllString(v, -1)
	if len(result) < 1 {
		return
	}
	firsStr := result[0]
	firsStr = strings.Trim(firsStr, " ")
	fild = strings.Replace(firsStr, "=", "", -1)
	content = strings.Replace(v, result[0], "", 1)
	content = strings.TrimSpace(content)
	fild = strings.TrimSpace(fild)
	if content == "" {
		return
	}
	ok = true
	return
}

func setValue(v reflect.Value, t reflect.Type, fild, content string) {
	fmt.Println(content)
	tField, ok := t.FieldByName(fild)
	if ok {
		rField := v.Elem().FieldByName(fild)
		if tField.Type.Kind() == reflect.Int64 {
			i, err := strconv.ParseInt(content, 10, 64)
			if err == nil {
				rField.SetInt(i)
			}
		} else if tField.Type.Kind() == reflect.String {
			rField.SetString(content)
		}
	}
}
