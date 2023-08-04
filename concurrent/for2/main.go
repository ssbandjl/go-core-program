package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func(i int) {

			fmt.Println(i)
			ch <- struct{}{}
			<-ch
		}(i)
	}
	for j := 0; j < 10; j++ {

	}
	time.Sleep(time.Millisecond * 10000)
}
