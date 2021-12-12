package database

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
    database, err := sqlx.Open("mysql", "root:111111@tcp(127.0.0.1:3306)/nineego")
    if err != nil {
        fmt.Println("open mysql failed,", err)
        return
    }
    db = database
    defer db.Close()
}

func InsertSometing()  {
  sql := "insert into useree(username,sex, email)values (?,?,?)"
	value := [3]string{"user01", "man", "user01@163.com"}

	//执行SQL语句
	r, err := db.Exec(sql, value[0], value[1], value[2])
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}

	//查询最后一天用户ID，判断是否插入成功
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}
	fmt.Println("insert succ", id)
}
