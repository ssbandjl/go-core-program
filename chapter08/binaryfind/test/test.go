package main

import (
	"fmt"
)

var middleIndex = 0

func run(num int, arr [6]int, leftIndex int, rightIndex int){
	// fmt.Println("原始数组", arr)
	if leftIndex > rightIndex{
		return
	}

	middleIndex	= (leftIndex+rightIndex)/2
	fmt.Println(middleIndex)
	if arr[middleIndex] < num {
		run(num, arr, middleIndex + 1, rightIndex) //因为已经查找过一次了
	}else if arr[middleIndex] > num{
		run(num, arr, leftIndex, middleIndex -1)
	}else{
		fmt.Println("找到index=", middleIndex)
	}
}


func main(){
	arr := [6]int{1,8, 10, 89, 1000, 1234}
	run(1234, arr, 0, len(arr)-1)
	run(10000, arr, 0, len(arr)-1)
}