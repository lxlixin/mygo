package basicknowledge

//具名结构体
type Employee struct {
	firstName, lastName string
	age                 int
}

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
