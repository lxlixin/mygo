package basicknowledge

import (
	"fmt"
)

//反射：通过反射可以获取丰富的类型信息
// go语言的反射没有内置类型工厂，无法做到像java一样通过类型字符串创建对象实例
// 反射是把双刃剑，功能强大但代码可读性低，并且性能低
// go中的反射有两个基本概念：type和value，也是反射中最重要的两个类型

// TypeClaim 类型断言
func TypeClaim() {
	var x interface{} // <type,value>
	x = 1234.56
	value, ok := x.(float64)
	fmt.Printf("猜对了?%v，x是一个%T，x的值为：%v \n", ok, value, value)
}
