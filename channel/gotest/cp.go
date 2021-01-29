package gotest

import (
	"log"
)

func Cp() {
	signal1 := make(chan bool)
	defer close(signal1)

	signalExit := make(chan bool)
	defer close(signalExit)
	go func() {
	L:
		for {
			select {
			case <-signal1:
				log.Printf("收到退出信号")
				break L
			default:
				log.Printf("工作中")
			}
			// time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			log.Printf("第%d次检查", i)
		}
		signal1 <- true
		signalExit <- true
	}()
	// time.Sleep(time.Second * 5)
	// signal1 <- true
	<-signalExit
	log.Printf("end")
}
