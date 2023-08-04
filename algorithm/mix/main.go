package main

import "fmt"

func main() {
	output := (1+6)/2*4 ^ 2 + 10%3<<3 // 7/2=3余1, 1100 ^ 0010 = 1110=14, 1<<3=1*2的3次方
	fmt.Printf("output:%v\n", output)
}
