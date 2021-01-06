package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "cloud1688\n"
	b := strings.ReplaceAll(a, "\n", "")
	fmt.Println("Start")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("End")
}
