package basicknowledge

//return 内部执行：先初始化返回值变量 x ，执行defer 代码，再执行汇编指令 RET

//F1 aha
func F1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

//F2 en
func F2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

//F3 ui
func F3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

//F4 90
func F4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
