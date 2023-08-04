package main

import (
	"fmt"
	"log"
)

type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

func exist(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	log.Printf("行(高)h=%d,列(宽)w=%d", h, w) //h=3,w=4
	vis := make([][]bool, h)              // 为了防止重复遍历相同的位置，需要额外维护一个与 \textit{board}board 等大的 \textit{visited}visited 数组，用于标识每个位置是否被访问过。每次遍历相邻位置时，需要跳过已经被访问的位置
	for i := range vis {
		vis[i] = make([]bool, w) //初始化
	}
	fmt.Println("初始化后的已检查数组:", vis)  // [[false false false false] [false false false false] [false false false false]]
	var check func(i, j, k int) bool // 设函数 \text{check}(i, j, k)check(i,j,k) 表示判断以网格的 (i, j)(i,j) 位置出发，能否搜索到单词 \textit{word}[k..]word[k..]，其中 \textit{word}[k..]word[k..] 表示字符串 \textit{word}word 从第 kk 个字符开始的后缀子串
	check = func(i, j, k int) bool {
		if board[i][j] != word[k] { // 剪枝：当前字符不匹配
			return false
		}
		if k == len(word)-1 { // 单词存在于网格中
			return true
		}
		vis[i][j] = true                     // 先设置为已检查
		defer func() { vis[i][j] = false }() // 回溯时还原已访问的单元格
		for _, dir := range directions {     // 遍历方向
			log.Printf("newI=%d, newI=%d", i+dir.x, j+dir.y)
			if newI, newJ := i+dir.x, j+dir.y; 0 <= newI && newI < h && 0 <= newJ && newJ < w && !vis[newI][newJ] { // i=0, j=0, newI=-1, newJ=0, 没有被检查
				log.Printf("递归: newI=%d, newI=%d", i+dir.x, j+dir.y)
				if check(newI, newJ, k+1) { //递归
					return true
				}
			}
		}
		return false
	}
	for i, row := range board { //行i,列j
		for j := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// 作者：LeetCode-Solution
// 链接：https://leetcode-cn.com/problems/word-search/solution/dan-ci-sou-suo-by-leetcode-solution/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	// [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]
	// fmt.Println([]byte("ABCE"))
	// fmt.Println([]byte("SFCS"))
	// fmt.Println([]byte("ADEF"))
	board := [][]byte{{65, 66, 67, 69}, {83, 70, 67, 83}, {65, 68, 69, 70}}
	word := "ABCCED"
	fmt.Println(exist(board, word))
}
