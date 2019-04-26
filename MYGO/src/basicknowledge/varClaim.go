package basicknowledge

import (
	"fmt"
)

//变量名只是在我们变成中方便识别记忆的名称，实际上程序在编译时，所有的变量名都会转换成变量名所对应的地址空间。
//对变量名的处理是编译器的工作。编译器中有符号表，记录变量名和对应的空间索引

//变量的声明 go是静态类型语言，不能在运行期间改变变量类型
//第一种  指定变量类型，声明后若不赋值，使用默认值
var name string

//第二种 根据值自行判断变量类型
var age = 18

//第三种 变量声明并赋值
var aha string = "lixin"

//第四种 一次性声明多个变量
var Mon, Tues, Wed, Thur, Fri, Sat, Sun int

//第五种 短声明变量，短声明变量不能在函数外使用
func shorClaim() {
	name := "helloworld!"
	fmt.Println(name)
}

//第六种 匿名变量，不像Java中，因为要返回值不得不定义好多变量
func getName() (firstName, lastName, nickName string) {
	return "xin", "li", "sweet heart"
}

// _, _, nickName := getName()
// fmt.Println(nickName)

//const 声明常量
// iota 是特殊的常量，是可被编译器修改的常量
//每次出现一次iota换行 ，counter会增加1，直到再次出现const iota会被重置为0
//go 语言的常量是无类型的，只要这个常量在相应类型的值域范围内，就可以赋值给该类型的变量
const Pi = 3.1415926
const KB = 1024

// const {
// 	MB KB<<iota*10
// 	GB
// }

//枚举，枚举指一系列相关常量
