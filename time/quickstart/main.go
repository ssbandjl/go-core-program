package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("%+v\n", time.Now().Local().UnixNano())
	fmt.Printf("当前时间戳:%+v\n", time.Now().Unix()+10)
	fmt.Printf("当前时间:%+v\n", time.Now())
	fmt.Printf(time.Now().Format("2006-01-02 15:04:05"))

}
