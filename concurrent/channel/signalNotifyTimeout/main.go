package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 在实践中，我们需要设置一个最长的等待时间。只要超过了这个时间，程序就不再等待，可以直接退出
func main() {
	var closing = make(chan struct{})
	var closed = make(chan struct{})

	go func() {
		// 模拟业务处理
		for {
			select {
			case <-closing:
				log.Printf("从closing中获取到元素(通道关闭信号已收到)")
				return
			default:
				// ....... 业务计算
				time.Sleep(100 * time.Millisecond)
				log.Printf("默认case语句")
			}
		}
	}()

	// 处理CTRL+C等中断信号
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	close(closing)
	// 执行退出之前的清理动作
	go doCleanup(closed)

	select {
	case <-closed:
	case <-time.After(time.Second):
		fmt.Println("清理超时，不等了")
	}
	fmt.Println("优雅退出")
}

func doCleanup(closed chan struct{}) {
	time.Sleep((time.Minute))
	close(closed)
}
