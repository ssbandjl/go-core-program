package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 除了正常的业务处理时的 wait/notify，我们经常碰到的一个场景，就是程序关闭的时候，我们需要在退出之前做一些清理（doCleanup 方法）的动作。这个时候，我们经常要使用 chan
func doCleanup() {
	log.Printf("连接关闭、文件 close、缓存落盘等一些动作")
}

func main() {
	go func() {
		// 执行业务处理
		log.Printf("执行业务逻辑")
		time.Sleep(time.Second * 2)
	}()

	// 处理CTRL+C等中断信号
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	// 执行退出之前的清理动作
	doCleanup()

	fmt.Println("优雅退出")
}
