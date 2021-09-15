package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

func main() {
	//nsSvc := fmt.Sprintf("%s.%s", instanceSvc, namespace)
	//mysqlConnStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/mysql?charset=utf8&parseTime=True&loc=Local&timeout=5s", "root", passwd, nsSvc, 3306)
	////util.Log4Zap(zap.InfoLevel).Info(fmt.Sprintf("连通性检查,连接信息:%s\n", mysqlConnStr))
	mysqlConnStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8&parseTime=True&loc=Local&timeout=3s", "root", "crss123456", "172.16.23.85", "32390")

	//mysqlConnStr := "root:crss123456@tcp(:)/mysql?charset=utf8&parseTime=True&loc=Local&timeout=5s"

	DB, err := gorm.Open("mysql", mysqlConnStr)
	defer DB.Close()
	if err != nil {
		//root:cloud1688@tcp(slave-cv-mysql.cloudpepper-v2-cv:3306)/mysql?charset=utf8&parseTime=True&loc=Local&timeout=5s
		log.Printf("无法联通")
		return
	}
	if DB.Error != nil {
		log.Printf("数据库错误")
		return
	}
	log.Printf("连接成功")
}
