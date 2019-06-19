package basicknowledge

// go 语言内置的testing包，调用go test就可以执行测试代码
// 并且不会被go build编译到最终的可执行文件中
// 测试相关的文件名XXX_test.go
// 基本测试（单元测试）*testing.T：测试程序的逻辑是否正确 函数名以Test开头
// 基准测试 *testing.B：测试函数的性能，函数名以Benchmark开头
// 示例函数 *testing.E：为文档提供示例，函数名以Example开头
// 反射包中的reflect.DeepEqual
// https://juejin.im/post/5ce93447e51d45775746b8b0

func add(x, y int) int {
	return x + y
}
