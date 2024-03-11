package main

import "fmt"

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	fmt.Println(searchMatrix([][]int{{1}}, 1))
}

func searchMatrix(matrix [][]int, target int) bool {
	row := searchRow(matrix, target)
	left, right := 0, len(matrix[row])-1
	for left <= right {
		mid := (right-left)/2 + left
		if target > matrix[row][mid] {
			left = mid + 1
		} else if target < matrix[row][mid] {
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}

func searchRow(data [][]int, target int) int {
	left, right := 0, len(data)-1
	for left < right {
		rowMid := (right-left)/2 + left
		if data[rowMid][0] <= target {
			left = rowMid
		} else {
			right = rowMid - 1
		}
	}
	return left
}
