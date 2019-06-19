package basicknowledge

//go语言工作空间：编译工具对源码有严格的要求，每个工作空间（workspace）必须由bin,pkg,src三个目录组成
//package main 表示一个可独立执行的程序，每个Go应用程序都包含一个名为main的包。
//获取第三方库：go get + 完整包名，在执行go get命令之前，确保电脑配置了环境变量GOPATH，并且安装git
//在go语言中根据首字母的大小写来确定可以访问的权限。首字母大写是公有的，首字母小写只能在本包中使用
/*
	go程序启动执行顺序如下：
		按顺序导入所有被main包引用的其他包，然后在每个包中执行如下流程：
		如果该包中又导入了其他的包，则从第一步开始递归执行，但是每个包只会被导入一次。
		然后以相反的顺序在每个包中初始化常量和变量，如果该包中含有init函数的话，则调用该函数。
		在完成这一切之后，main也执行同样的过程，最后调用main函数开始执行程序。
*/

//go语言是编译型语言，与编译型语言相对的是解释型语言。
// 编译型语言在程序执行之前，有一个单独的编译过程，将程序翻译成机器语言，以后执行这个程序的时候，就不需要再进行翻译了
// 解释型语言，是在运行的时候，将程序翻译成机器语言，所以运行速度相对编译型语言要慢。
// java属于解释型语言，虽然Java在程序运行之前也有一个编译过程，但是并不是将程序编译成机器语言，而是将它翻译成字节码（可以理解为一个中间语言.class）
// java 跨平台使用的是虚拟机。go 没有跨平台，不是一次编译处处使用。

// go 常用命令
// 1. go run：直接运行程序／编译且运行，不生成二进制文件
// 2. go build：测试编译，检查是否有编译错误／ 以当前所在“包”命名，生成“包名“二进制文件
// 3. go fmt：格式化源码
// 4. go install：编译包文件并编译整个程序／生成二进制文件到$GOPATH/src目录下
// 5. go test：运行测试文件
// 6. go doc：查询文档，godoc -http=:6060，然后在本地访问：localhost:6060，启动本地文档

//go 的工具不完善 比如切片slice 没有delete方法，删除一个元素比较笨拙

// 标准库中文文档：https://studygolang.com/pkgdoc

//一个推荐的go学习博客：https://www.cnblogs.com/liuzhongchao/category/1226355.html

//李文周老师的博客地址：https://www.liwenzhou.com

//2019 go https://github.com/gopherchina/conference/tree/master/2019

// .Go Web轻量级框架Gin学习系列：HTTP请求日志 https://juejin.im/post/5cc84cfa51882562457fe311

// 协程初步讨论 https://segmentfault.com/a/1190000019114754

// go框架-beego基础 https://c.isme.pub/2019/03/25/go-beego/

// 在 struct 內的 pointers 跟 values 差异 https://blog.wu-boy.com/2019/05/what-is-different-between-pointer-and-value-in-golang/

// 从零开始搭建创业公司后台技术栈 https://mp.weixin.qq.com/s/jQ0X66DK1DDW9DFlyd9bdw

// Golang并发模型：轻松入门流水线模型 https://mp.weixin.qq.com/s?__biz=Mzg3MTA0NDQ1OQ==&mid=2247483671&idx=1&sn=1706ffa6deee44a367c34ef84448f55f&scene=21#wechat_redirect

// Golang并发模型：轻松入门流水线FAN模式https://mp.weixin.qq.com/s?__biz=Mzg3MTA0NDQ1OQ==&mid=2247483680&idx=1&sn=de463ebbd088c0acf6c2f0b5f179f38d&scene=21#wechat_redirect

// golang 正则表达式 https://www.cnblogs.com/golove/p/3269099.html

// beego
