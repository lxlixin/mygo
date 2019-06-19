package util_test

import (
	"practice/util"
	"testing"
)

func TestGetIntFromIP(t *testing.T) {
	ret, err := util.GetIntFromIP("129.211.131.251")
	if err != nil {
		t.Errorf("方法运行错误，错误信息为：%v", err)
	}
	if ret != 2178122747 {
		t.Errorf("ip 运算错误，正确结果为：%v，得到的结果为：%v", 2178122747, ret)
	}
	t.Logf("ip 运算正确，得到的结果为：%v", ret)
}
