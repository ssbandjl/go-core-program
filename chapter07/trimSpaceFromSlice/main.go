package main

import "fmt"

func TrimSpace(s []byte) []byte {
	b := s[:0]
	for _, x := range s {
		if x != ' ' {
			b = append(b, x)
		}
	}
	return b
}

func main() {
	ret := TrimSpace([]byte{' ', '1', ' ', '3'})
	fmt.Println(ret)
}
