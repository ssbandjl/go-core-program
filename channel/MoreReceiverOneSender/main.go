package main

//情形一：M个接收者和一个发送者。发送者通过关闭用来传输数据的通道来传递发送结束信号
//这是最简单的一种情形。当发送者欲结束发送，让它关闭用来传输数据的通道即可。

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

// GoID 获取goroutine id
func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false) //all如果为真, 则记录所有goroutine, 为假时表示当前goroutine
	// log.Printf("协程栈帧信息:%s", string((buf[:n])))
	// log.Printf("去掉前缀后:%s", strings.TrimPrefix(string(buf[:n]), "goroutine "))
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0] //去除前缀, 并按默认分隔符分割为切片, 获取切片第一个元素, 即得到goroutine id
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func main() {
	log.Printf("main goroutine id:%d", GoID())
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const Max = 10
	const NumReceivers = 5

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int)

	// 发送者
	go func() {
		for {
			if value := rand.Intn(Max); value == 0 {
				log.Printf("发送value=%d", value)
				// 此唯一的发送者可以安全地关闭此数据通道。
				close(dataCh)
				return
			} else {
				log.Printf("发送value=:%d", value)
				dataCh <- value
			}
		}
	}()

	// 接收者
	var wg sync.WaitGroup
	for i := 0; i < NumReceivers; i++ {
		wg.Add(1)
		go func() {
			defer wgReceivers.Done()

			// 接收数据直到通道dataCh已关闭
			// 并且dataCh的缓冲队列已空。
			defer wg.Done()
			log.Printf("goroutine id:%d", GoID())
			for value := range dataCh {
				log.Printf("接受者, 收到value=%d", value)
			}
		}()
	}

	wgReceivers.Wait()
}
