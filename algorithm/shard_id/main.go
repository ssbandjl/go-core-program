package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(16777216 | (uint64(i) << 4))
	}

}
