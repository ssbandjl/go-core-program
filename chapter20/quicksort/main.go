package main

import (
	"fmt"
	_ "math/rand"
	"time"
)

//快速排序
//说明
//1. left 表示 数组左边的下标
//2. right 表示数组右边的下标
//3.  array 表示要排序的数组
//4. 如果左右指针重叠，需要将他们错开
func QuickSort(left int, right int, array *[9]int) {
	l := left
	r := right
	// pivot 是中轴， 支点,中间那个数
	pivot := array[(left+right)/2]
	fmt.Printf("中轴pivot:%v\n", pivot)
	temp := 0

	//for 循环的目标是将比 pivot 小的数放到 左边
	//  比 pivot 大的数放到 右边
	for l < r { //退出条件就是l >= r
		//从  pivot 的左边找到大于等于pivot的值, 因为要交换, 修改这里的两个比较符号可以修改排序方向
		for array[l] < pivot {
			l++
		}
		//从  pivot 的右边边找到小于等于pivot的值
		for array[r] > pivot {
			r--
		}
		// 1 >= r 表明本次分解任务完成, break,说明未找到
		if l >= r {
			break
		}
		//交换
		temp = array[l]
		array[l] = array[r]
		array[r] = temp
		fmt.Printf("交换后:%v\n", array)

		//优化，如果左/右指针走到中间位置,需要移动一位，便于递归
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
		// for循环结束后,l可能比r小
	}
	// 如果  1== r, 再移动下, 不要在比较了, 防止死循环
	if l == r {
		l++
		r--
	}
	// 向左递归
	// fmt.Println(l, r)
	if left < r {
		QuickSort(left, r, array)
	}
	// 向右递归
	if right > l {
		QuickSort(l, right, array)
	}
}

var arr = [9]int{1, 10, 3, 4, 5, 6, 7, 8, 9}

func main() {

	// arr := [9]int{1, 10, 3, 4, 5, 6, 7, 8, 9}
	// arr := [9]int {-9,78,0,23,-567,70, 123, 90, -23}
	fmt.Println("初始", arr)

	// var arr [8000000]int
	// for i := 0; i < 8000000; i++ {
	// 	arr[i] = rand.Intn(900000)
	// }

	//fmt.Println(arr)
	start := time.Now().Unix()
	//调用快速排序
	QuickSort(0, len(arr)-1, &arr)
	end := time.Now().Unix()
	fmt.Println("main..")
	fmt.Printf("快速排序法耗时%d秒", end-start)
	fmt.Println(arr)

}
