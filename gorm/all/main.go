package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//定义结构
type User struct {
	gorm.Model
	Name      string `gorm:"unique"`
	Pets      []*Pet //一个用户有多个宠物
	Toys      []Toy  `gorm:"polymorphic:Owner"` //一个用户有多个玩具,多态
	ManagerID *uint
	Manager   *User //单表,属于某个经理Manager管理,取消构成单边关联
	//Team []User `gorm:"foreignkey:ManagerID"`  //管理多个团队Team,单表
	Friends []*User `gorm:"many2many:user_friends"` //拥有很多朋友,单表
}

//宠物
type Pet struct {
	gorm.Model
	UserID int
	Toy    Toy `gorm:"polymorphic:Owner;"`
}

//玩具
type Toy struct {
	gorm.Model
	OwnerID   string
	OwnerType string
}

func main() {
	//连接数据库
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:root@tcp(data:3306)/all?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("数据库连接失败,连接信息:%s\n", dsn)
		panic("failed to connect database")
	}

	// 迁移 schema 新建数据表
	//db.AutoMigrate(&User{}, &Group{}, &Permission{}, &KubernetesInfo{}, &MysqlInstance{}, &MysqlSvcLabel{}, &UserGroup{}, &GroupPermission{})
	db.AutoMigrate(&User{}, &Pet{}, &Toy{})
	fmt.Printf("数据库自动迁移完成\n")

	//新建第1个用户
	//db.Save(&User{
	//	Name: "sxb",
	//})
	//
	////新建第2个用户,管理sxb
	//db.Save(&User{
	//	Name: "jl",
	//	Team: []User{{Name: "sxb"}},
	//})

	//新建第2个用户,指定其上级
	var father User
	db.Where("name = ?", "sxb").First(&father)

	db.Save(&User{
		Name:    "xbb3",
		Manager: &father,
	})

}
