package basicknowledge

import "fmt"

//Employee 具名结构体，构造函数是值类型
type Employee struct {
	firstName, lastName string
	age                 int
}

//NewEmployee 构造函数
func NewEmployee(firstName, lastName string, age int) *Employee {
	return &Employee{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}
}

//匿名结构体
// emp1 := struct {
// 	firstName, lastName string
// 	age                 int
// }{"xiao", "宏", 20}

//Hub 是一个练习
type Hub struct {
	cs map[string][]string
}

//NewHub 约定俗成，构造函数以new开头
func NewHub() *Hub {
	return &Hub{
		cs: make(map[string][]string),
	}
}
func main() {
	hub := NewHub()
	if cs, ok := hub.cs["test"]; ok {
		fmt.Printf("cs type is %T\n", cs)
		cs = []string{"1", "2"}
		cs = append(cs, "3")
		hub.cs["test"] = cs
		fmt.Printf("%+v\n", hub)
		cs = append(cs[:1], cs[2:]...)
	}
	cs, ok := hub.cs["test"]
	fmt.Println(cs, ok)
	fmt.Println("testing")
	fmt.Printf("%+v\n", hub)
}
