package main

import "fmt"

func tt(str string) bool {
	if len(str) == 0 {
		return false
	}
	var r []rune = []rune(str)
	// 上海自来水来自海上
	// 0              len-1
	i, j := 0, len(r)-1
	for i < j {
		if r[i] == r[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func main() {
	str := "上海自来水来自海上"
	fmt.Println(tt(str))
}
