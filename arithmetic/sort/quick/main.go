package main

import "log"

func sort(array []int, low, high int) {
	if low >= high {
		return
	}
	i := low
	j := high
	var index int
	index = array[i]
	for i < j {
		for i < j && array[j] >= index {
			j--
		}
		if i < j {
			array[i] = array[j]
			i++
		}
		for i < j && array[i] < index {
			i++
		}
		if i < j {
			array[j] = array[i]
		}
	}
	array[i] = index
	sort(array, low, i-1)
	sort(array, i+1, high)
}

func QuickSort(array []int) {
	sort(array, 0, len(array)-1)
}

func main() {
	data := []int{5, 4, 9, 8, 7, 6, 0, 1, 3, 2}
	QuickSort(data)
	log.Printf("快速排序结果:%+v", data)
}
