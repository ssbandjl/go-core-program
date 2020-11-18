package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	mysqlConnStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/mysql?charset=utf8&parseTime=True&loc=Local", "root", "root1", "data", 3306)
	DB, err := gorm.Open(mysql.Open(mysqlConnStr), &gorm.Config{})
	if err != nil {
		fmt.Printf("连接数据错误:%s", err.Error())
	}
	fmt.Printf("连接数据成功:%v\n", DB)
}
