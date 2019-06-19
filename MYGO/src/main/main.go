package main

import (
	"practice/signature"
)

// _ "github.com/go-sql-driver/mysql"

//User  用户实体表
type User struct {
	id   int
	age  int
	name string
}

func main() {
	// err := mysql.MyInit("root:123@tcp(127.0.0.1:3306)/go_test")
	// if err != nil {
	// 	fmt.Printf("链接数据库失败，error：%v\n", err)
	// }

	// // insert := "insert into user(id,name,age) value(?,?,?)"
	// // mysql.DB.Prepare(insert)

	// seleStr := "select id,name,age from user where id = ?"
	// user := &User{}
	// row := mysql.DB.QueryRow(seleStr, 1)
	// err = row.Scan(&user.id, &user.name, &user.age)
	// if err != nil {
	// 	fmt.Printf("获取用户失败，错误信息：%v \n", err)
	// 	return
	// }
	// fmt.Println(user)

	signature.SendToDuodian()
}
