package homework

import (
	"fmt"
	"os"
)

type student struct {
	name, class string
	age         int
	id          int64
}

func newStudent(name, class string, age int, id int64) *student {
	return &student{
		name:  name,
		class: class,
		age:   age,
		id:    id,
	}
}

type StudentMsg struct {
	students []*student
}

func showMsg() {
	fmt.Println("这是学生管理平台!")
	fmt.Println("1. 添加学生")
	fmt.Println("2. 修改学生")
	fmt.Println("3. 删除学生")
	fmt.Println("4. 展示学生")
	fmt.Println("5. 列出所有学生")
	fmt.Println("6. 退出")
}

//获取学生生成的学生对象
func (s *StudentMsg) getStudent() {
	var (
		name, class string
		age         int
		id          int64
	)
	fmt.Println("添加学生信息开始，请输入学生姓名并换行：")
	fmt.Scanln(&name)
	fmt.Println("请输入学生班级并换行：")
	fmt.Scanln(&class)
	fmt.Println("请输入学生年龄并换行：")
	fmt.Scanln(&age)
	fmt.Println("请输入学生ID并换行：")
	fmt.Scanln(&id)
	s.students = append(s.students, newStudent(name, class, age, id))
}

//修改学生信息
func (s *StudentMsg) updateStudent() {
	var (
		name, class string
		age         int
		id          int64
		conte       bool
	)
	flag := true
	conte = true
	var sstu *student
	for flag && conte {
		fmt.Println("请输入要修改学生信息的ID：")
		fmt.Scanln(&id)
		for _, s := range s.students {
			if s.id == id {
				sstu = s
				flag = false
			}
		}
		if sstu == nil {
			fmt.Println("要修改信息的学生没有找到，是否选择继续？（true/false）")
			fmt.Scanln(&conte)
		}
	}
	if sstu != nil {
		fmt.Println("请输入学生班级并换行：")
		fmt.Scanln(&class)
		sstu.class = class
		fmt.Println("请输入学生年龄并换行：")
		fmt.Scanln(&age)
		sstu.age = age
		fmt.Println("请输入学生姓名并换行：")
		fmt.Scanln(&name)
		sstu.name = name
	}
}

//展示学生信息
func (s *StudentMsg) showStudent() {
	var id int64
	fmt.Println("请输入要查找的学生ID并换行：")
	fmt.Scanln(&id)
	for _, s := range s.students {
		if s.id == id {
			fmt.Printf("找到ID为%d的学生了，姓名为%s，年龄为%d，班级为%s \n", s.id, s.name, s.age, s.class)
			return
		}
	}
	fmt.Println("没有找到对应的学生信息！")
}

//展示所有学生信息
func (s *StudentMsg) listStudent() {
	for _, s := range s.students {
		fmt.Printf("姓名为%s，年龄为%d，班级为%s \n", s.name, s.age, s.class)
	}
}

//删除学生信息
func (ss *StudentMsg) delStudent() {
	var id int64
	fmt.Println("请输入要删除的学生ID并换行：")
	fmt.Scanln(&id)
	for index, s := range ss.students {
		if s.id == id {
			ss.students = append(ss.students[:index], ss.students[index+1:]...)
			return
		}
	}
	fmt.Println("没有找到对应的学生信息！")
}

func (s *StudentMsg) OptStudent() {
	fmt.Printf("%T\n", s.students)
	fmt.Println(s.students)
	if s.students == nil {
		fmt.Println("is nil")
	}
	for {
		showMsg()
		var opt int
		fmt.Scanln(&opt)
		switch opt {
		case 1:
			s.getStudent()
			break
		case 2:
			s.updateStudent()
			break
		case 3:
			s.delStudent()
			break
		case 4:
			s.showStudent()
			break
		case 5:
			s.listStudent()
			break
		case 6:
			os.Exit(0)
		}
	}

}
