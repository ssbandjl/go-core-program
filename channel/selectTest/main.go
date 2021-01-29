package main

import (
	"log"
	"time"
)

func main() {
	signal1 := make(chan bool)
	defer close(signal1)
	signal1 <- true
	go func() {
		for {
			select {
			case <-signal1:
				log.Printf("收到结束信号")
				break
			default:
				log.Printf("1工作中")
			}
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		for {
			select {
			case <-signal1:
				log.Printf("收到结束信号")
				break
			default:
				log.Printf("2工作中")
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		time.Sleep(time.Second * 3)
		signal1 <- true
	}()

	// time.Sleep(time.Second * 5)
	// signal1 <- true
	time.Sleep(time.Second * 5)
	log.Printf("优雅退出")
}
