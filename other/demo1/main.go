package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	// var sum int
	for _, v := range s {
		// sum += v
		sum += v
	}
	fmt.Println(sum)
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c) //0-3  7, 2, 8  =17
	go sum(s[len(s)/2:], c) //3-5  -9, 4, 0  =-5
	x, y := <-c, <-c        // receive from c

	// fmt.Println(x, y, x+y)
	fmt.Printf("x=%v, y=%v, x+y=%v\n", x, y, x+y)
}
