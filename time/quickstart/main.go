package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("%+v\n", time.Now().Local().UnixNano())
	fmt.Printf("当前时间戳:%+v\n", time.Now().Unix()+10)
}
