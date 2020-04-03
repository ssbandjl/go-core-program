package main

import (
	"fmt"
	"go_code/account/utils"
)

func main() {
	fmt.Println("面向对象的方式来完成.....")
	utils.NewMyFamilyAccount().MainMenu() //NewMyFamilyAccount返回了一个结构体指针，该结构体绑定了一个MainMenu方法
}
