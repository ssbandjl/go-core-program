package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price int
}

func main() {
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:root@tcp(data:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println(db)
	// 迁移 schema 新建数据表
	err = db.AutoMigrate(&Product{})
	fmt.Printf("迁移错误:%+v", err)

	var tables []string
	db.Raw("show tables").Scan(&tables)
	fmt.Println(tables)

	// Create 增加数据
	db.Create(&Product{Code: "D42", Price: 100})
	/*
		[405.304ms] [rows:0] CREATE TABLE `products` (`id` bigint unsigned AUTO_INCREMENT,`created_at` datetime(3) NULL,`updated_at` datetime(3) NULL,`deleted_at` datetime(3) NULL,`code` longtext,`price` bigint unsigned,PRIMARY KEY (`id`),INDEX idx_products_deleted_at (`deleted_at`))
	*/

	// Read 读取数据
	var product Product
	db.First(&product, 1) // 根据整形主键查找
	fmt.Printf("根据主键查找结果:\n%+v\n", product)

	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	fmt.Printf("根据code=D42查找结果:\n%+v\n", product)

	// Update - 将 product 的 price 更新为 200
	//db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 软删除 product
	//db.Delete(&product, 1)

	//永久删除
	//db.Unscoped().Delete(&product, 1)
}
