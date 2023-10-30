package main

import "fmt"

// https://leetcode.cn/problems/valid-sudoku
func main() {
	fmt.Println(isValidSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}))
}

func isValidSudoku(board [][]byte) bool {
	var coll, row [9][9]bool
	var subBox [3][3][9]bool
	for i, rowItem := range board {
		for j, cell := range rowItem {
			if cell == '.' {
				continue
			}
			charIdx := cell - '1'
			if row[i][charIdx] || coll[j][charIdx] || subBox[i/3][j/3][charIdx] {
				return false
			}
			row[i][charIdx] = true
			coll[j][charIdx] = true
			subBox[i/3][j/3][charIdx] = true
		}
	}
	return true
}
