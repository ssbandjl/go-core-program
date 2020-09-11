package main

import "fmt"

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	var Slice1 = []int{1, 2, 3, 4, 5}
	fmt.Printf("slice1: %v\n", Slice1)

	Slice2 := remove(Slice1, 2)
	fmt.Printf("slice2: %v\n", Slice2)
}
