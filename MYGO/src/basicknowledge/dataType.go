package basicknowledge

import "fmt"

//go语言内置以下数据类型
//布尔类型 ： bool  不能跟其他类型做转换
//整型：int8,byte,int16,int,uint,uintptr；int 根据操作系统决定
//浮点类型：float32, float64
//复数类型
//字符串类型：string
//字符类型：rune
//错误类型：error

//所有值类型变量在赋值和作为参数传递时都将产生一次复制动作。

//复合类型
//指针（pointer）
//数组（array）- 值类型 赋值任然是值语义，数组的长度也是数组的一部分
//切片（slice）-值类型 赋值仍然是值语义，切片的主体（属性）是一个数组的指针，and 使用切片的时候有语法糖，自动调用切片的指针
//字典（map）-引用类型
//通道（chan）-引用类型
//结构体（struct）-值类型
//接口（interface）－引用类型
//函数 function －引用类型

func Helloworld() {
	// var helloworld1 int
	// helloworld1 = 1
	// fmt.Println(helloworld1)
	// // myFun()

	// array := [...]int{1, 3, 5, 7, 8}

	// var total int
	// for _, value := range array {
	// 	total += value
	// }
	// fmt.Println(total)

	// for i := 0; i < len(array); i++ {
	// 	tempVar := array[i]
	// 	other := 8 - tempVar
	// 	for j := i + 1; j < len(array); j++ {
	// 		if array[j] == other {
	// 			fmt.Printf("%d 元素是%d 和 %d = 8 \n", j, array[i], array[j])
	// 			// i = j
	// 			break
	// 		}
	// 	}
	// }

	var mArraym [2][2][2]int
	var array2 [2][2]int = [2][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
	}
	mArraym[0] = array2
	fmt.Println(mArraym)

	sliceMap := make(map[string][]int, 10)
	sliceMap["hello"] = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sliceMap)
}
