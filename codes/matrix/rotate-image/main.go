package main

import "fmt"

// https://leetcode.cn/problems/rotate-image
func main() {
	data := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(data)
	for _, item := range data {
		for _, v := range item {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}

func rotate(matrix [][]int) {
	if matrix == nil || len(matrix) == 0 || len(matrix) == 1 {
		return
	}
	n := len(matrix)
	tmp := make([][]int, n)
	for idx := range tmp {
		tmp[idx] = make([]int, n)
	}
	for i, item := range matrix {
		for j, cell := range item {
			tmp[j][n-i-1] = cell
		}
	}
	copy(matrix, tmp)
}
