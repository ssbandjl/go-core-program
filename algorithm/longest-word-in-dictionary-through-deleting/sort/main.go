package main

import (
	"fmt"
	"sort"
)

/*
给定提供的 less 函数，Slice 对切片 x 进行排序。 如果 x 不是切片，它会恐慌。
排序不能保证是稳定的：相等的元素可能会从它们的原始顺序颠倒过来。 对于稳定排序，请使用 SliceStable。
less 函数必须满足与接口类型的 Less 方法相同的要求。
*/
// findLongestWord 返回ans
func findLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		a, b := dictionary[i], dictionary[j]
		return len(a) > len(b) || len(a) == len(b) && a < b //字符个数多的在前面, 字符个数相等时,字母序小的在前面
	})
	fmt.Println("排序后的dictionary:", dictionary)
	for _, t := range dictionary {
		i := 0
		for j := range s {
			if s[j] == t[i] {
				i++
				if i == len(t) {
					return t // 直接返回第一个即可, 因为已经排序
				}
			}
		}
	}
	return ""
}

func main() {
	dictionary := []string{"ale", "bpple", "cpple", "apple", "monkey", "plea"}
	s := "apple"
	fmt.Println("result:", findLongestWord(s, dictionary))
}

// https://leetcode-cn.com/problems/longest-word-in-dictionary-through-deleting/solution/tong-guo-shan-chu-zi-mu-pi-pei-dao-zi-di-at66/
