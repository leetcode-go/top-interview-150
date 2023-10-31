package main

import "fmt"

// https://leetcode.cn/problems/spiral-matrix
func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	var direction = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var i, j, m, n = 0, 0, len(matrix), len(matrix[0])
	var result = make([]int, m*n)
	var direct = 0
	for idx := 0; idx < m*n; idx++ {
		result[idx] = matrix[i][j]
		matrix[i][j] = -101
		nextI := i + direction[direct][0]
		nextJ := j + direction[direct][1]
		if nextI == m || nextI < 0 || nextJ == n || nextJ < 0 || matrix[nextI][nextJ] == -101 {
			direct = (direct + 1) % 4
		}
		i = direction[direct][0] + i
		j = direction[direct][1] + j
	}
	return result
}
