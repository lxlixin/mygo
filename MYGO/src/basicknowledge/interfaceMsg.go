package basicknowledge

import "fmt"

//Animal animal common function
type Animal interface {
	Run() func() int
	Eat()
}

//Tiger big cat
type Tiger struct {
	name string
	age  int
	sex  int
}

//Run tiger run
func (a *Tiger) Run() func() int {
	fmt.Println("tiger running ...")
	return func() int {
		a.age++
		return a.age
	}
}

//Eat tiger eat
func (*Tiger) Eat() {
	fmt.Println("tiger eatting ...")
}

type Puma struct {
	name string
	age  int
}

func (a Puma) Run() func() int {
	fmt.Println("puma running ...")
	return func() int {
		a.age++
		return a.age
	}
}

func (Puma) Eat() {
	fmt.Println("puma eatting ...")
}

//Test111 Tiger 和 *Tiger 是两个不同的类型； 是*Tiger 实现了Animal 接口，而不是Tiger 实现了Animal 接口
//如果写成 t := Tiger{name: "cute", age: 12, sex: 1}  那么anm = t 就会报错，因为Tiger没有实现Animal接口
//疑问：：： interface 是值类型还是引用类型？
func Test111() {
	var anm Animal
	t := &Tiger{name: "cute", age: 12, sex: 1}
	anm = t
	f1 := anm.Run()
	fmt.Println(f1())
	fmt.Println(t.age)
	var anm2 Animal
	anm2 = anm
	fmt.Println(anm2.Run()())
	fmt.Println(t.age)
	p := Puma{name: "cool", age: 12}
	anm = p
	f2 := anm.Run()
	fmt.Println(f2())
	fmt.Println(p.age)
}
