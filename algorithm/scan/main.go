package main

import "fmt"

func main() {
	var (
		name     string
		age      int
		is_marry bool
	)

	fmt.Scan(&name, &age, &is_marry)
	fmt.Printf("获取结果 name:%s age:%d is_marry:%t \n", name, age, is_marry)
}
