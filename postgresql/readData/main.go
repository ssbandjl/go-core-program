package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	// 	connStr := "postgres://postgres:cloud1688@172.16.24.200:31976?sslmode=disable&connect_timeout=3"
	// Initialize connection constants.
	HOST     = "172.16.24.200"
	PORT     = 31976
	DATABASE = "postgres"
	USER     = "postgres"
	PASSWORD = "cloud1688"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Initialize connection string. 初始化连接字符串, 参数包含主机,端口,用户名,密码,数据库名,SSL模式(禁用),超时时间
	var connectionString string = fmt.Sprintf("host=%s  port=%d user=%s password=%s dbname=%s sslmode=disable connect_timeout=3", HOST, PORT, USER, PASSWORD, DATABASE)

	// Initialize connection object. 初始化连接对象, 驱动名为postgres
	db, err := sql.Open("postgres", connectionString)
	checkError(err)

	err = db.Ping() //连通性检查
	checkError(err)
	fmt.Println("Successfully created connection to database")

	//读取数据
	// Read rows from table.
	var id int
	var name string
	var quantity int

	sql_statement := "SELECT * from inventory;"
	rows, err := db.Query(sql_statement)
	checkError(err)
	defer rows.Close()

	for rows.Next() {
		switch err := rows.Scan(&id, &name, &quantity); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			fmt.Printf("Data row = (%d, %s, %d)\n", id, name, quantity)
		default:
			checkError(err)
		}
	}
}
