package basicknowledge

// go 语言内置的testing包
// 测试先关的文件名XXX_test.go
// 基本测试 单元测试 函数名以Test开头  *testing.T
// 基准测试 *testing.B
// 示例函数
// 反射包中的reflect.DeepEqual
// https://juejin.im/post/5ce93447e51d45775746b8b0

func add(x, y int) int {
	return x + y
}
