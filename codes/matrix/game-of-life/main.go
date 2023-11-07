package main

import "fmt"

// https://leetcode.cn/problems/game-of-life
func main() {
	matrix := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0}}
	gameOfLife(matrix)
	for _, row := range matrix {
		for _, cell := range row {
			fmt.Printf("%d ", cell)
		}
		fmt.Println()
	}
}

func gameOfLife(board [][]int) {
	if board == nil {
		return
	}
	directs := [][]int{{1, 0}, {1, 1}, {0, 1}, {-1, 0}, {0, -1}, {-1, -1}, {1, -1}, {-1, 1}}
	m, n := len(board), len(board[0])
	for i, row := range board {
		for j, cell := range row {
			die, live := 0, 0
			for _, direct := range directs {
				x, y := i+direct[0], j+direct[1]
				if x >= 0 && x < m && y >= 0 && y < n {
					if board[x][y] == 1 || board[x][y] == -1 {
						live += 1
					} else if board[x][y] == 0 || board[x][y] == 2 {
						die += 1
					}
				}
			}
			if cell == 0 && live == 3 {
				board[i][j] = 2
				continue
			}
			if cell == 1 && (live < 2 || live > 3) {
				board[i][j] = -1
			}
		}
	}
	for i, row := range board {
		for j := range row {
			if board[i][j] == 2 {
				board[i][j] = 1
			} else if board[i][j] == -1 {
				board[i][j] = 0
			}
		}
	}
}
