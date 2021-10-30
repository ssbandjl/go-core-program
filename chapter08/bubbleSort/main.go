package main

import (
	"fmt"
)

//冒泡排序
func BubbleSort(arr *[5]int) {

	fmt.Println("排序前arr=", (*arr))
	temp := 0 //临时变量(用于做交换)

	// 冒泡排序..一步一步推导出来的
	// 第一次比较数组长度次,找出最大的, 其次第二大, ...
	for i := 0; i < len(*arr)-1; i++ {

		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				//交换, 将大的往后排
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}

	}

	fmt.Println("排序后arr=", (*arr))

}

func main() {

	//定义数组
	arr := [5]int{24, 69, 80, 57, 13}
	//将数组传递给一个函数，完成排序

	BubbleSort(&arr)

	fmt.Println("main arr=", arr) //有序? 是有序的

	var arr2 = [5]int{20, 11, -1, 44, 9}
	Test(&arr2) //这里需要执行go build
	fmt.Println("手写冒泡排序法排序后:\n", arr2)
}
