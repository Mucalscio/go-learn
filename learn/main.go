package main

import (
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "github.com/jmoiron/sqlx"
  "github.com/pkg/errors"
  "os"
)

var db *sqlx.DB

func init() {
    database, err := sqlx.Open("mysql", "root:111111@tcp(127.0.0.1:33060)/nineego")
    if err != nil {
        fmt.Println("open mysql failed,", err)
        return
    }
    db = database
    fmt.Println("连接Mysql成功！")
}

func main()  {
  fmt.Println("主程序开始咯")
  id, err := doDB()
  if err != nil {
    fmt.Println(err);
    os.Exit(1)
  }
  fmt.Println("插入数据ID：", id);
}

func doDB() (uint, error) {
  id, err := InsertSometing()
  if err != nil {
    return 0, errors.Wrap(err, "调用插入函数发生错误")
  }
  return id, err
}

func InsertSometing() (uint, error) {
  sql := "insert into useree(username,sex, email)values (?,?,?)"
	value := [3]string{"user01", "man", "user01@163.com"}

	r, err := db.Exec(sql, value[0], value[1], value[2])
	if err != nil {
    return 0, errors.Wrap(err, "执行SQL出错了")
	}

  id, err := r.LastInsertId()
  if err != nil {
    return 0, errors.Wrap(err, "插入数据没有成功")
	}

  return uint(id), nil;
}
