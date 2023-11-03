package main

import "fmt"

// https://leetcode.cn/problems/set-matrix-zeroes
func main() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	for _, row := range matrix {
		for _, cell := range row {
			fmt.Printf("%d ", cell)
		}
		fmt.Println()
	}
}

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	if matrix == nil || m == 0 || (m == 1 && n == 1) {
		return
	}
	zeroX, zeroY := make([]int, m), make([]int, n)
	for x, row := range matrix {
		for y, cell := range row {
			if cell == 0 {
				zeroX[x] = 1
				zeroY[y] = 1
			}
		}
	}
	for x, item := range zeroX {
		if item == 1 {
			for y := 0; y < n; y++ {
				matrix[x][y] = 0
			}
		}
	}

	for y, item := range zeroY {
		if item == 1 {
			for x := 0; x < m; x++ {
				matrix[x][y] = 0
			}
		}
	}
}
