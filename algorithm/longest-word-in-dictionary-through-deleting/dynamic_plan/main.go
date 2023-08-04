package main

import (
	"fmt"
)

// findLongestWord 返回ans, 状态转移方程
func findLongestWord(s string, dictionary []string) (ans string) {
	m := len(s)
	f := make([][26]int, m+1)
	for i := range f[m] {
		f[m][i] = m
	}
	for i := m - 1; i >= 0; i-- {
		f[i] = f[i+1]
		f[i][s[i]-'a'] = i
	}

outer:
	for _, t := range dictionary {
		j := 0
		for _, ch := range t {
			if f[j][ch-'a'] == m {
				continue outer
			}
			j = f[j][ch-'a'] + 1
		}
		if len(t) > len(ans) || len(t) == len(ans) && t < ans {
			ans = t
		}
	}
	return
}

func main() {
	dictionary := []string{"ale", "bpple", "cpple", "apple", "monkey", "plea"}
	s := "apple"
	fmt.Println("result:", findLongestWord(s, dictionary))
}

// https://leetcode-cn.com/problems/longest-word-in-dictionary-through-deleting/solution/tong-guo-shan-chu-zi-mu-pi-pei-dao-zi-di-at66/
