package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "HelloWord你好"
	l1 := len([]rune(str))
	l2 := bytes.Count([]byte(str), nil) - 1
	l3 := strings.Count(str, "") - 1
	l4 := utf8.RuneCountInString(str)
	fmt.Println(l1)
	fmt.Println(l2)
	fmt.Println(l3)
	fmt.Println(l4)
	fmt.Println(len(str))
}
