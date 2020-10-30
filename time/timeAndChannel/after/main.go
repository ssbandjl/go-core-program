package main

import (
	"fmt"
	"time"
)

func main() {
	//使用After(),返回值为只读通道 <- chanTime, 与Timer.C相同
	ch1 := time.After(time.Second * 5) //相当于创建一个通道 NewTimer(d).C
	fmt.Println("当前时间:", time.Now())
	data := <-ch1
	fmt.Printf("data类型:%T\n", data)
	fmt.Println("data:", data)
}
