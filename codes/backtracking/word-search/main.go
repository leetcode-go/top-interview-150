package main

import "fmt"

// https://leetcode.cn/problems/word-search
func main() {
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"))
}

var directs = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

func exist(board [][]byte, word string) bool {
	m, n, wordLen := len(board), len(board[0]), len(word)
	if m*n < len(word) {
		return false
	}

	var backTrack func(i, j int, wordIndex int) bool

	backTrack = func(i, j int, wordIndex int) bool {
		if wordIndex == wordLen {
			return true
		}
		if i >= m || j >= n || i < 0 || j < 0 {
			return false
		}
		tmp := board[i][j]
		if tmp == word[wordIndex] {
			board[i][j] = '-'
			res := false
			for _, direct := range directs {
				res = res || backTrack(i+direct[0], j+direct[1], wordIndex+1)
			}
			board[i][j] = tmp
			return res
		} else {
			return false
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backTrack(i, j, 0) {
				return true
			}
		}
	}
	return false
}
