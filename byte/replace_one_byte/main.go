package main

import "fmt"

func main() {
	str := "hello"
	var bs []byte = []byte(str)
	bs[0] = 'x'
	fmt.Println(string(bs))

	// str := "hello"
	// str[0] = 'x'
	// fmt.Println(str)
}
