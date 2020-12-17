//这是最复杂的一种情形。我们不能让接收者和发送者中的任何一个关闭用来传输数据的通道，我们也不能让多个接收者之一关闭一个额外的信号通道。
//这两种做法都违反了通道关闭原则。 然而，我们可以引入一个中间调解者角色并让其关闭额外的信号通道来通知所有的接收者和发送者结束工作。 具体实现见下例。
//注意其中使用了一个尝试发送操作来向中间调解者发送信号。
//在此例中，通道关闭原则依旧得到了遵守。
//
//请注意，信号通道toStop的容量必须至少为1。 如果它的容量为0，则在中间调解者还未准备好的情况下就已经有某个协程向toStop发送信号时，此信号将被抛弃。
//
//我们也可以不使用尝试发送操作向中间调解者发送信号，但信号通道toStop的容量必须至少为数据发送者和数据接收者的数量之和，以防止向其发送数据时（有一个极其微小的可能）导致某些发送者和接收者协程永久阻塞。

package main

import (
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const Max = 100000
	const NumReceivers = 10
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int)
	stopCh := make(chan struct{})
	// stopCh是一个额外的信号通道。它的发送
	// 者为中间调解者。它的接收者为dataCh
	// 数据通道的所有的发送者和接收者。

	toStop := make(chan string, 1)
	// toStop是一个用来通知中间调解者让其
	// 关闭信号通道stopCh的第二个信号通道。
	// 此第二个信号通道的发送者为dataCh数据
	// 通道的所有的发送者和接收者，它的接收者
	// 为中间调解者。它必须为一个缓冲通道。

	var stoppedBy string

	// 中间调解者
	go func() {
		stoppedBy = <-toStop
		close(stopCh)
	}()

	// 发送者
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(Max)
				if value == 0 {
					// 为了防止阻塞，这里使用了一个尝试
					// 发送操作来向中间调解者发送信号。
					select {
					case toStop <- "发送者#" + id:
					default:
					}
					return
				}

				// 此处的尝试接收操作是为了让此发送协程尽早
				// 退出。标准编译器对尝试接收和尝试发送做了
				// 特殊的优化，因而它们的速度很快。
				select {
				case <-stopCh:
					return
				default:
				}

				// 即使stopCh已关闭，如果这个select代码块
				// 中第二个分支的发送操作是非阻塞的，则第一个
				// 分支仍很有可能在若干个循环步内依然不会被选
				// 中。如果这是不可接受的，则上面的第一个尝试
				// 接收操作代码块是必需的。
				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// 接收者
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			defer wgReceivers.Done()

			for {
				// 和发送者协程一样，此处的尝试接收操作是为了
				// 让此接收协程尽早退出。
				select {
				case <-stopCh:
					return
				default:
				}

				// 即使stopCh已关闭，如果这个select代码块
				// 中第二个分支的接收操作是非阻塞的，则第一个
				// 分支仍很有可能在若干个循环步内依然不会被选
				// 中。如果这是不可接受的，则上面尝试接收操作
				// 代码块是必需的。
				select {
				case <-stopCh:
					return
				case value := <-dataCh:
					if value == Max-1 {
						// 为了防止阻塞，这里使用了一个尝试
						// 发送操作来向中间调解者发送信号。
						select {
						case toStop <- "接收者#" + id:
						default:
						}
						return
					}

					log.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}

	// ...
	wgReceivers.Wait()
	log.Println("被" + stoppedBy + "终止了")
}
