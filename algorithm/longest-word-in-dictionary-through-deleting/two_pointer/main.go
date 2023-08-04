package main

import "fmt"

// findLongestWord 返回ans
func findLongestWord(s string, dictionary []string) (ans string) {
	// 一层遍历数组
	for _, t := range dictionary { //t=ale
		i := 0 //数组中的字符串索引
		// 二层遍历字符串
		for j := range s { //j=0
			if s[j] == t[i] {
				i++              // 贪心地匹配
				if i == len(t) { // 最终如果 ii 移动到 tt 的末尾，则说明 tt 是 ss 的子序列
					if len(t) > len(ans) || len(t) == len(ans) && t < ans { // 通过遍历 \textit{dictionary}dictionary 中的字符串，并维护当前长度最长且字典序最小的字符串
						ans = t
					}
					break
				}
			}
		}
	}
	return
}

func main() {
	dictionary := []string{"ale", "apple", "monkey", "plea"}
	s := "apple"
	fmt.Println("result:", findLongestWord(s, dictionary))
}

// https://leetcode-cn.com/problems/longest-word-in-dictionary-through-deleting/solution/tong-guo-shan-chu-zi-mu-pi-pei-dao-zi-di-at66/
