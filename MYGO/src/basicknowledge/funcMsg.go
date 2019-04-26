package basicknowledge

import "fmt"

//语法糖：语法对语言的功能并没有影响，但是更方便程序员使用

func closure(num int) func(x int) int {
	return func(a int) int {
		num++
		return a + num
	}
}

func Test() {
	f := closure(1)
	fmt.Println(f(2))
	fmt.Println(f(2))
	fmt.Println(f(2))
}
