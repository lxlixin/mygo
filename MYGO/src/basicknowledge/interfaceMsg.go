package basicknowledge

import "fmt"

//Animal animal common function
type Animal interface {
	Run()
	Eat()
}

//Tiger big cat
type Tiger struct {
	name string
	age  int
	sex  int
}

//Run tiger run
func (*Tiger) Run() {
	fmt.Println("tiger running ...")
}

//Eat tiger eat
func (*Tiger) Eat() {
	fmt.Println("tiger eatting ...")
}

//Test111 Tiger 和 *Tiger 是两个不同的类型； 是*Tiger 实现了Animal 接口，而不是Tiger 实现了Animal 接口
//如果写成 t := Tiger{name: "cute", age: 12, sex: 1}  那么anm = t 就会报错，因为Tiger没有实现Animal接口
func Test111() {
	var anm Animal
	t := &Tiger{name: "cute", age: 12, sex: 1}
	anm = t
	anm.Run()
}
