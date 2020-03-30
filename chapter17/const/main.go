package main

import (
	"fmt"
)

func main() {

	var num int
	num = 9 //ok
	//常量声明的时候，必须赋值。
	const tax int = 0
	//常量是不能修改
	//tax = 10
	fmt.Println(num, tax)
	//常量只能修饰bool、数值类型(int, float系列)、string 类型

	//fmt.Println(b)

	const (
		a = iota //附零值，后面的常量依次+1，按行递增；常量不一定大写，但是有访问范围的限制（首字母大写可以导出访问）；掌握运行时和编译时
		b
		c
		d
	)

	fmt.Println(a, b, c, d)
}
