package util

import (
	"errors"
	"strconv"
	"strings"
)

//GetIntFromIP ip string to int64
func GetIntFromIP(ip string) (intIP int64, err error) {
	ipNums := strings.Split(ip, ".")
	for index, str := range ipNums {
		tempInt, err := strconv.Atoi(str)
		if err != nil {
			return 0, errors.New("ip 转换错误")
		}
		intIP += int64(tempInt) << (uint(3-index) * 8)
	}
	return intIP, nil
}
