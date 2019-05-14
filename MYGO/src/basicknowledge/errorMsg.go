package basicknowledge

import (
	"errors"
	"fmt"
)

// 异常处理（error接口，panic，recover）

// 在默认情况下，当发生错误（panic）后，程序就会终止运行
// 如果发生错误后，可以捕获错误信息，并通知管理员（邮件，短信，日志），程序还可以继续执行
// panic错误有时候是不可避免的，如果不捕获，程序就会终止
// 使用recover 来捕获panic错误，但是recover只有在defer中才有效
// 在未发生panic时调用recover，recover会返回nil
// 有时候发生的错误，没有严重到至使程序终止，使用panic抛出错误就会很危险
// go 语言中添加了error 接口，如果只是想要标识有问题而不想造成程序终止，可以使用error
// 通常的用法是，func最后一个返回值是error类型，当调用者调用完成后，可以通过判断error是否为nil来判断方法有没有正确执行
// 错误和异常是不同的概念
// java中没有错误概念，如果遇到了什么问题都是抛异常，这样其实是不对的，error很好的解决了这个问题

// 推荐方法返回error
func errorMsg1(varam string) (result string, err error) {
	if varam == "" {
		err = errors.New("param is nil !")
	}
	result = varam
	return
}

func errorMsg2(password string) (err error) {
	realPassword := "helloworld"
	if password == "" {
		panic("this is dangerous ")
	}
	if realPassword != password {
		err = errors.New("password is wrong, may you can again ")
	}
	return err
}

func errorMsg3() {
	//校验密码是否正确，如果不小心没有输入密码，请不要让程序自杀
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("这里发生了严重的事故，但是我们暂且没有停止该操作，请赶快来检查一下")
		}
	}()
	errorMsg2("")
}
