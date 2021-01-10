package main

import "fmt"

// send 和 recv 都可以作为 select 语句的 case clause
func main() {
	var ch = make(chan int, 10)
	// chan 还可以应用于 for-range 语句中
	for v := range ch {
		fmt.Println(v)
	}

	//或者是忽略读取的值，只是清空 chan
	for range ch {
	}

	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
		case v := <-ch:
			fmt.Println(v)
		}
	}
}
