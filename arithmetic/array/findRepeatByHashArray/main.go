package main

import (
	"fmt"
)

//查找重复的数,使用hash法

func FindDupByHash(arr []int) int {
	if arr == nil {
		return -1
	}
	data := map[int]bool{}
	for _, v := range arr {
		if _, ok := data[v]; ok { //如果当前index之前设置为true表示重复,比如data[3]=True,表示之前在该位置设置过一次
			return v
		} else {
			data[v] = true
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 4, 2, 5, 3}
	fmt.Println("Hash法")
	fmt.Println(FindDupByHash(arr))
}
