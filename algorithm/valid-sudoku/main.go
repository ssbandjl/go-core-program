package main

// https://leetcode-cn.com/problems/valid-sudoku/solution/you-xiao-de-shu-du-by-leetcode-solution-50m6/

import (
	"encoding/json"
	"fmt"
)

func isValidSudoku(board [][]byte) bool {
	var rows, columns [9][9]int
	var subboxes [3][3][9]int
	for i, row := range board {
		for j, c := range row {
			if c == '.' {
				continue
			}
			index := c - '1'
			rows[i][index]++
			columns[j][index]++
			subboxes[i/3][j/3][index]++
			if rows[i][index] > 1 || columns[j][index] > 1 || subboxes[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}

type Input struct {
	TwoDArray [][]string `json:"twoDArray"`
}

func main() {
	input := `{"twoDArray":[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]}`
	var board [][]byte = make([][]byte, 9)
	var inputObj Input
	err := json.Unmarshal([]byte(input), &inputObj)
	if err != nil {
		fmt.Println(err.Error())
	}
	for i := 0; i < len(inputObj.TwoDArray); i++ {
		board[i] = make([]byte, 9)
		for j := 0; j < len(inputObj.TwoDArray[i]); j++ {
			board[i][j] = []byte(inputObj.TwoDArray[i][j])[0]
		}
	}
	fmt.Println("board:", board)
	fmt.Println(isValidSudoku(board))
}
