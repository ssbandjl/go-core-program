package main

import (
	"gorm.io/gorm"
)

func main() {
	nsSvc := fmt.Sprintf("%s.%s", instanceSvc, namespace)
	mysqlConnStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/mysql?charset=utf8&parseTime=True&loc=Local&timeout=5s", "root", passwd, nsSvc, 3306)
	//util.Log4Zap(zap.InfoLevel).Info(fmt.Sprintf("连通性检查,连接信息:%s\n", mysqlConnStr))
	mysqlConnStr := "root:cloud1688@tcp(slave-cv-mysql.cloudpepper-v2-cv:3306)/mysql?charset=utf8&parseTime=True&loc=Local&timeout=5s"

	DB, err := gorm.Open("mysql", mysqlConnStr)
	defer DB.Close()
	if err != nil {
		//root:cloud1688@tcp(slave-cv-mysql.cloudpepper-v2-cv:3306)/mysql?charset=utf8&parseTime=True&loc=Local&timeout=5s
		return false
	}
	if DB.Error != nil {
		return false
	}
	return true
}
