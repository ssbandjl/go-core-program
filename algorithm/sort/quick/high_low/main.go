// https://leetcode-cn.com/problems/queue-reconstruction-by-height/solution/gen-ju-shen-gao-zhong-jian-dui-lie-by-leetcode-sol/
package main

import (
	"fmt"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	// fmt.Println(people)
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})
	fmt.Println("从小到大排序:", people)
	ans := make([][]int, len(people))
	// for _, p := range people {
	// 	fmt.Println(p[0])
	// }
	for _, person := range people {
		// fmt.Println(person)
		spaces := person[1] + 1
		// spaces := person[0] + 1
		// fmt.Println(spaces)
		for i := range ans {
			if ans[i] == nil {
				spaces--
				if spaces == 0 {
					ans[i] = person
					break
				}
			}
		}
	}
	return ans
}

func main() {
	// var input [][]int = {4, 1, 5, 2, 3}
	// var input = [][]int{{4}, {1}, {5}, {2}, {3}}
	var input = [][]int{{4, 1, 5, 2, 3}, {1}, {5}, {2}, {3}}
	fmt.Println(reconstructQueue(input))
	// input := "4 1 3 5 2"
	// fmt.Println("4 1 5 2 3")
}
