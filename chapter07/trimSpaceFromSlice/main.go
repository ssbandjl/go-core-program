package main

import "fmt"

//利用长度为0的空切片特性，删除字节切片中的空格
func TrimSpace(s []byte) []byte {
	b := s[:0]
	fmt.Println(b)
	for _, x := range s {
		if x != ' ' {
			b = append(b, x)
			fmt.Println(b)
		}
	}
	return b
}

func main() {
	ret := TrimSpace([]byte{' ', '1', ' ', '3'})
	fmt.Println(ret)
}
