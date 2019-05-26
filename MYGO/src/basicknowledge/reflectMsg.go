package basicknowledge

import (
	"fmt"
)

//反射：通过反射可以获取丰富的类型信息
// go语言的反射没有内置类型工厂，无法做到像java一样通过类型字符串创建对象实例
// 反射是把双刃剑，功能强大但代码可读性低，并且性能低
// go中的反射有两个基本概念：type和value，也是反射中最重要的两个类型
// type 是类型对象，在反射中type还划分为：类型（type）和 种类（kind）
// value
// reflect 是内置包，在reflect包中提供了reflect.TypeOf 和 reflect.ValueOf 两个函数来获取任意对象的Value和Type

// TypeClaim 类型断言
func TypeClaim() {
	var x interface{} // <type,value>
	x = 1234.56
	value, ok := x.(float64)
	fmt.Printf("猜对了?%v，x是一个%T，x的值为：%v \n", ok, value, value)
}

// 修改值需要满足可被寻址的要求，因此我们使用的是 reflect.ValueOf(&v)
// v := 42
// valueV := reflect.ValueOf(&v)
// valueV.Elem().SetInt(1024)
// fmt.Println(v) // 1024

// 修改结构体字段时，要保证字段能被导出，首字母大写
// s := &Stu{
//     "Hank",
//   }
//   valueS := reflect.ValueOf(s)
//   valueS.Elem().FieldByName("Name").SetString("Kang")
//   fmt.Println(s.Name) // Kang
