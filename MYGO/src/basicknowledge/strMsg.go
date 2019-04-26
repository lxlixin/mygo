package basicknowledge

import (
	"fmt"
)

/*
	strings 包实现了用于操作字符的简单函数
	https://mp.weixin.qq.com/s?__biz=MzU0ODc4MjE0Nw==&mid=2247483785&idx=1&sn=d00ce09f798dd285ebaf35407357aa75&scene=21#wechat_redirect
	数组
	在go语言中，数组是固定长度的数据类型！数据类型！
*/

//在go语言中支持两种字符类型，
//一个是byte(实际上是unit8的别名)，代表UTF-8字符串单个字节的值；
//另一个是rune，代表单个Unicode字符
//UTF-8是对Unicode编码规范的一种实现，实现根据变长存储

//通过制定数据类型和元素个数来声明数组
var array [5]int

//声明一个长度为5的整数数组并初始化
var array1 [5]int = [5]int{1, 2, 3, 4, 5}

func Myfun() {
	array2 := [5]string{"1", "2", "3", "4", "5"}
	fmt.Print(array2)
	array3 := [...]string{"1"}
	fmt.Println(array3)

	//string 是值类型，不可修改；
	s1 := "hello"
	fmt.Printf("%p\n", &s1)
	byteA := []byte(s1)
	fmt.Printf("%p\n", &byteA)
	le := len(byteA) - 1
	for i := 0; i < len(byteA)/2; i++ {
		if i != (le - i) {
			temp := byteA[i]
			byteA[i] = byteA[le-i]
			byteA[le-i] = temp
		}
	}
	fmt.Println(string(byteA))
	fmt.Printf("%p\n", &s1)
	fmt.Println("修改以后", s1)
	s66 := fmt.Sprint("hellow", "world")
	fmt.Println(s66)
}
