package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func GetDb(host string, port string, username string, password string) (*sql.DB, error) {
	mysqlConnStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/mysql?charset=utf8&parseTime=True&loc=Local", username, password, host, port)
	Db, err = sql.Open("mysql", mysqlConnStr)
	return Db, err
}

func main() {
	//获取数据库连接
	Db, err := GetDb("172.16.24.200", "31422", "root", "test123")
	if err != nil {
		panic(err)
	}
	rows, _ := Db.Query("show slave status")
	defer rows.Close()

	cols, _ := rows.Columns() //列名
	buff := make([]interface{}, len(cols))
	data := make([]string, len(cols))
	for i, _ := range buff {
		buff[i] = &data[i]
	}

	for rows.Next() {
		rows.Scan(buff...)
	}

	for k, col := range data {
		fmt.Printf("%30s:\t%s\n", cols[k], col)
	}
}
