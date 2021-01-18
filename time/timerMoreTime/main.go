package main

import (
	"fmt"
	"time"
)

// Ticker：时间到了，多次执行
func main() {
	// 1.获取ticker对象
	ticker := time.NewTicker(3 * time.Second)
	i := 0
	// 子协程
	go func() {
		for {
			//<-ticker.C
			i++
			fmt.Println(<-ticker.C) // 每隔一段时间,放一个数据到C通道中
			if i == 5 {
				//停止
				ticker.Stop()
			}
		}
	}()
	for {
	}
}
