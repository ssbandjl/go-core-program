package main

import (
	"log"
	"time"
)

type Token struct{}

func newWorker(id int, ch chan Token, nextCh chan Token) {
	for {
		token := <-ch // 阻塞, 直到取得令牌
		// fmt.Println((id + 1)) // id从1开始
		log.Printf("Goroutine%d开始执行", id+1)
		time.Sleep(time.Second)
		nextCh <- token //将token给下一个通道
	}
}
func main() {
	chs := []chan Token{make(chan Token), make(chan Token), make(chan Token), make(chan Token)}

	// 创建4个worker
	for i := 0; i < 4; i++ {
		log.Printf("执行协程newWorker, 通道ID:i=%d,(i+1)对4取余=%d", i, (i+1)%4)
		go newWorker(i, chs[i], chs[(i+1)%4])
	}

	//首先把令牌交给第一个worker
	chs[0] <- struct{}{}

	select {}
}

/*
我来给你具体解释下这个实现方式。首先，我们定义一个令牌类型（Token），接着定义一个创建 worker 的方法，这个方法会从它自己的 chan 中读取令牌。哪个 goroutine 取得了令牌，就可以打印出自己编号，因为需要每秒打印一次数据，所以，我们让它休眠 1 秒后，再把令牌交给它的下家。接着，在第 16 行启动每个 worker 的 goroutine，并在第 20 行将令牌先交给第一个 worker。如果你运行这个程序，就会在命令行中看到每一秒就会输出一个编号，而且编号是以 1、2、3、4 这样的顺序输出的。这类场景有一个特点，就是当前持有数据的 goroutine 都有一个信箱，信箱使用 chan 实现，goroutine 只需要关注自己的信箱中的数据，处理完毕后，就把结果发送到下一家的信箱中
*/
