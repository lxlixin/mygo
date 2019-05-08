package basicknowledge

import "fmt"

//数组切片不仅仅是一个指针，它拥有自己的结构
//一个指向原生数组的指针
//数组切片中元素的个数
//数组切片已经分配的存储空间
//数组的长度在定义后无法修改，数组切片是指针类型?  个人偏向切片是值类型，go语法糖会简化操作，但切片本身是值类型，或者说切片是内置的一个结构体，结构体中有一个属性是指向数组的元素的指针
//基于数组的一个封装
//https://www.cnblogs.com/liuzhongchao/p/9159896.html

func TempFun() {

	//基于数组创建数组切片，左包含右不包含
	myArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	mySlice1 := myArray[:5]
	// for _, item := range mySlice1 {
	// 	fmt.Println(item)
	// }

	//基于make()方法创建数组切片,元素初始值为0
	//make 用来给引用类型做初始化了
	// mySlice2 := make([]int, 8)
	// for _, item := range mySlice2 {
	// 	fmt.Println(item)
	// }

	//直接创建并初始化包含5个元素的数组切片
	// mySlice3 := []int{1, 2, 3, 4, 5}
	// for _, item := range mySlice3 {
	// 	fmt.Println(item)
	// }

	//获取切片中指针的指向的第一个元素的地址
	fmt.Printf("1: %p\n", &myArray)
	//获取到的地址同上，所以证明是go 对切片的取地址进行了处理
	fmt.Printf("2: %p\n", mySlice1)
	//获取切片变量的地址
	fmt.Printf("3: %p\n", &mySlice1)

	// array1 := []int{1, 2}
	// slice1 := array1[:]
	// slice2 := slice1
	// fmt.Println(slice1)
	// fmt.Println(slice2)
	// slice1[1] = 44
	// fmt.Printf("1：%p\n", &slice1)
	// fmt.Printf("2：%p\n", &slice2)
	// fmt.Printf("3：%p\n", array1)
	// fmt.Printf("4：%p\n", slice1)
	// fmt.Printf("5：%p\n", slice2)
	// slice1 = append(slice1, 3, 4, 5, 6, 7, 8)
	// fmt.Printf("6：%p\n", &slice1)
	// fmt.Printf("7：%p\n", slice1)
	// fmt.Println("8: ", slice1)
	// fmt.Println("9：", slice2)
}
