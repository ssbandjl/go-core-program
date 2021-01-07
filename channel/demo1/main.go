package main

import (
	"fmt"
	"time"
)

func process(timeout time.Duration) bool {
	ch := make(chan bool) //需要将通道改为带缓存的通道, 防止内存泄漏
	// ch := make(chan bool, 1)

	go func() {
		// 模拟处理耗时的业务
		time.Sleep((timeout + time.Second))
		ch <- true // block
		fmt.Println("exit goroutine")
	}()
	select {
	case result := <-ch:
		return result
	case <-time.After(timeout):
		return false
	}
}

func main() {
	process(time.Second * 3)
}
