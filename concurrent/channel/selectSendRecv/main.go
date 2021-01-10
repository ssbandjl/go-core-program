package main

import "fmt"

// send 和 recv 都可以作为 select 语句的 case clause
func main() {
	var ch = make(chan int, 10)
	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
		case v := <-ch:
			fmt.Println(v)
		}
	}
}
