package mysql

import (
	"database/sql"
)

//DB  数据库连接池
var DB *sql.DB

//MyInit  数据库初始化
func MyInit(dsn string) (err error) {
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		return err
	}
	//已经连接上数据库，设置最大连接数，最大空闲连接数
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(100)
	return nil
}
