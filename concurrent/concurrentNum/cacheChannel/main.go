package main

import (
	"fmt"
	"log"
)

//缓存通道
var limit = make(chan int, 2)

func main() {
	// work := map[string]func() int{"work1": func() int { return 1 }, "work2": func() int { return 2 }, "work3": func() int { return 3 }}
	work := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, w := range work {
		fmt.Println(w)
		go func() {
			limit <- 1
			log.Printf("%+v", w)
			<-limit
		}()
	}
	select {}
	// 	os.Exit(0)
}
