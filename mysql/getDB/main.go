package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var (
	Db  *sql.DB
	err error
)

func GetDb(host string, port int, username string, password string) (*sql.DB, error) {

	//参数dataSourceName格式:用户名:密码@[tcp(localhost:3306)]/数据库名, 这里的Db只是句柄，使用的时候才会调用
	//mysqlConnStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/mysql?charset=utf8&parseTime=True&loc=Local", username, password, host, port)

	mysqlConnStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/mysql?&charset=utf8&parseTime=True&loc=Local&timeout=5s", username, password, host, port)
	//fmt.Printf("GetDb mysqlConnStr:%v\n", mysqlConnStr)
	Db, err = sql.Open("mysql", mysqlConnStr)
	fmt.Println("sql.Open驱动正常")
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	//Db.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	//Db.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	//Db.SetConnMaxLifetime(time.Second * 10)
	Db.SetConnMaxIdleTime(time.Second * 10)
	if err != nil {
		fmt.Println("连接出错")
		return Db, err
	}
	fmt.Println("正在检测连通性...")

	err = Db.Ping()

	if err != nil {
		panic(err.Error())
	}

	//time.Sleep(time.Second * 15)
	var ret *sql.Rows
	ret, err = Db.Query("show databases")
	fmt.Println("查询结果", ret)
	return Db, err
}

func main() {
	DB, err := GetDb("192.168.1.1", 3306, "root", "root")
	fmt.Println("DB:", DB)
	fmt.Println("err:", err)
}
