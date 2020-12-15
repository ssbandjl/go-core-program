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

	// 以水果库存清单表inventory为例
	// Drop previous table of same name if one exists.  如果之前存在清单表, 则删除该表
	_, err = db.Exec("DROP TABLE IF EXISTS inventory;")
	checkError(err)
	fmt.Println("Finished dropping table (if existed)")

	// Create table. 创建表, 指定id, name, quantity(数量)字段, 其中id为主键
	_, err = db.Exec("CREATE TABLE inventory (id serial PRIMARY KEY, name VARCHAR(50), quantity INTEGER);")
	checkError(err)
	fmt.Println("Finished creating table")

	// Insert some data into table. 插入3条水果数据
	sql_statement := "INSERT INTO inventory (name, quantity) VALUES ($1, $2);"
	_, err = db.Exec(sql_statement, "banana", 150)
	checkError(err)
	_, err = db.Exec(sql_statement, "orange", 154)
	checkError(err)
	_, err = db.Exec(sql_statement, "apple", 100)
	checkError(err)
	fmt.Println("Inserted 3 rows of data")
}
