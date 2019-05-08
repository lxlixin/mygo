package basicknowledge

//闭包相关知识理解基本概念
// 闭包是可以包含自由（未绑定到特定对象）变量的代码块，这些变量不在这个代码块内或者
// 任何全局上下文中定义，而是在定义代码块的环境中定义。要执行的代码块（由于自由变量包含
// 在代码块中，所以这些自由变量以及它们引用的对象没有被释放）为自由变量提供绑定的计算环
// 境（作用域）。
//  闭包的价值
// 闭包的价值在于可以作为函数对象或者匿名函数，对于类型系统而言，这意味着不仅要表示
// 数据还要表示代码。支持闭包的多数语言都将函数作为第一级对象，就是说这些函数可以存储到
// 变量中作为参数传递给其他函数，最重要的是能够被函数动态创建和返回。
func Adder(x int) func(int) int {
	num := x + 1
	return func(delta int) int {
		num += delta
		return num
	}
}

// func main() {
// 	c := basicknowledge.Adder(0)  此处c为一个函数，c函数中饮用了Adder 函数中的num 变量，所以num没有被释放
//  d := basicknowledge.Adder(1)
// 	fmt.Println(c(1))
// 	fmt.Println(c(1))
// 	fmt.Println(c(1))
// 	fmt.Println(c(1))
// fmt.Println(d(1))
// fmt.Println(d(1))
// fmt.Println(d(1))
// fmt.Println(d(1))
// }
