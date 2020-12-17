package main

//情形二：一个接收者和N个发送者，此唯一接收者通过关闭一个额外的信号通道来通知发送者不要在发送数据了
//此情形比上一种情形复杂一些。我们不能让接收者关闭用来传输数据的通道来停止数据传输，因为这样做违反了通道关闭原则。 但是我们可以让接收者关闭一个额外的信号通道来通知发送者不要在发送数据了。
//如此例中的注释所述，对于此额外的信号通道stopCh，它只有一个发送者，即dataCh数据通道的唯一接收者。 dataCh数据通道的接收者关闭了信号通道stopCh，这是不违反通道关闭原则的。
//
//在此例中，数据通道dataCh并没有被关闭。是的，我们不必关闭它。 当一个通道不再被任何协程所使用后，它将逐渐被垃圾回收掉，无论它是否已经被关闭。 所以这里的优雅性体现在通过不关闭一个通道来停止使用此通道。

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const Max = 100
	const NumSenders = 10

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(1)

	// ...
	dataCh := make(chan int)
	stopCh := make(chan struct{})
	// stopCh是一个额外的信号通道。它的
	// 发送者为dataCh数据通道的接收者。
	// 它的接收者为dataCh数据通道的发送者。

	// 发送者
	for i := 0; i < NumSenders; i++ {
		go func() {
			for {
				// 这里的第一个尝试接收用来让此发送者
				// 协程尽早地退出。对于这个特定的例子，
				// 此select代码块并非必需。
				select {
				case <-stopCh:
					return
				default:
				}

				// 即使stopCh已经关闭，此第二个select
				// 代码块中的第一个分支仍很有可能在若干个
				// 循环步内依然不会被选中。如果这是不可接受
				// 的，则上面的第一个select代码块是必需的。
				randNum := rand.Intn(Max)
				log.Println("randNum:", randNum)
				select {
				case <-stopCh:
					return
				case dataCh <- randNum:
				}
			}
		}()
	}

	// 接收者
	go func() {
		defer wgReceivers.Done()

		for value := range dataCh {
			if value == Max-1 {
				// 此唯一的接收者同时也是stopCh通道的
				// 唯一发送者。尽管它不能安全地关闭dataCh数
				// 据通道，但它可以安全地关闭stopCh通道。
				close(stopCh)
				return
			}

			log.Println(value)
		}
	}()

	// ...
	wgReceivers.Wait()
}
