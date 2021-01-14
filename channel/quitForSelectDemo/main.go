package main

import (
	// "time"
	"fmt"
)

func produce(id int, ch chan string, quitSub chan int) {
	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("msg:%d:%d", id, i)
	}
	quitSub <- id
}

func main() {
	c1 := make(chan string, 1) //定义两个有缓冲的通道，容量为1
	c2 := make(chan string, 1)
	quitForSelect := make(chan int, 2) //用于通知退出for select循环
	quitTag := []int{}                 //用于读取和存储quitForSlect通道的两个值，保证两个product都写入完毕

	quitMain := make(chan bool) //阻塞main gorouting,等待for select 处理完成

	go produce(1, c1, quitForSelect)
	go produce(2, c2, quitForSelect)

	go func() {
		for { //使用select来等待这两个通道的值，然后输出
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case q := <-quitForSelect:
				fmt.Println("got quit tag for Gorouting id:", q)
				quitTag = append(quitTag, q)
				if len(quitTag) == 2 {
					fmt.Println("end to quit Main ")
					quitMain <- true
				}
			}
		}
	}()

	<-quitMain
	fmt.Println("exit from main")
}
