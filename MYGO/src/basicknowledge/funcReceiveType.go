package basicknowledge

import (
	"fmt"
)

//Person is a struct
type Person struct {
	age int
}

func (p Person) howOld() int {
	return p.age
}

func (p *Person) growUp() {
	p.age++
}

func (p Person) tempGrowUp() {
	p.age++
}

//TestVar is a function of test
//不管接收者类型是值类型还是指针类型，都可以通过值类型或指针类型调用，这里面实际上是通过语法糖起作用的
func TestVar() {
	//eg:
	p1 := Person{age: 10}
	p2 := &Person{age: 10}
	//值类型调用者调用值类型接受者的方法，方法会使用调用者的一个副本，类似“传值”
	fmt.Println(p1.howOld())
	//指针类型调用着调用值类型接收者的方法，指针被解引用为值，实际上是(*p2).howOld()
	fmt.Println(p2.howOld())
	//值类型调用者调用指针类型接收者的方法,使用值的引用来调用方法，实际上是(&p1).growUp()
	p1.growUp()
	//指针类型调用着调用指针类型接收者的方法，实际上也是“传值”，方法里的操作会影响到调用者，类似于指针传值，拷贝了一份指针
	p2.growUp()
	fmt.Println(p1.howOld())
	fmt.Println(p2.howOld())
	p1.tempGrowUp()
	p2.tempGrowUp()
	fmt.Println(p1.howOld())
	fmt.Println(p2.howOld())

}

type coder interface {
	code()
	debug()
}

type gopher struct {
	language string
}

func (g gopher) code() {
	fmt.Println("i am a gopher , i am coding !")
}

func (g *gopher) debug() {
	fmt.Printf("i am a gopher , i am debuging !")
}

// TestVar1 实现了接收者是值类型的方法，相当于自动实现了接收者是指针类型的方法。
// 而实现了接收者是指针类型的方法，不会自动生成接收者是值类型的方法
func TestVar1() {
	// var c1 coder
	var c2 coder
	// g1 := gopher{language: "go"}
	g2 := &gopher{language: "php"}
	// c1 = g1 wrong
	c2 = g2
	c2.code()
}
