package basicknowledge

import "fmt"

//执行顺序：先做全局的声明 再init 然后main
func init() {
	// fmt.Println("init1 函数用于package的初始化！")
}

func myFun() {
	fmt.Println("特殊标识符:_ \n 使用（import _ 包路径）只是引用该包，仅仅是为了调用init()函数，所以无法通过包名来调用包中的其他函数!")
}
