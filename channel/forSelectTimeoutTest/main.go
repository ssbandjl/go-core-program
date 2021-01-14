package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// n := rand.Intn(10) // n will be between 0 and 10
	// fmt.Printf("Sleeping %d seconds...\n", n)
	// time.Sleep(time.Duration(n) * time.Second)
	// fmt.Println("Done")
	// ch := make(chan int, 10)
	ch := make(chan int)
	defer close(ch)
	go func() {
		for i := 0; i <= 100; i++ {
			ch <- i
			if i >= 2 {
				log.Printf("数据发完了, 退出")
				// close(ch)
				break
			}
			n := rand.Intn(8) // n will be between 0 and 10
			log.Printf("本次休眠%d秒", n)
			time.Sleep(time.Duration(n) * time.Second)
		}
	}()

	quitCh := make(chan bool)
	defer close(quitCh)
	go func() {
		for {
			select {
			case num := <-ch:
				// if num == 2 {
				// 	log.Printf("通道已关闭, 不接收数据了")
				// 	quitCh <- true
				// }
				log.Printf("收到数字:%d", num)

			case <-time.After(time.Second * 5):
				log.Printf("超时")
				// close(ch)
				quitCh <- true
				// return
			}
			time.Sleep(time.Second)
		}
	}()

	<-quitCh
	log.Printf("程序退出")

}
