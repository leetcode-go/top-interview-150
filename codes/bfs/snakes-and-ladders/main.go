package main

import "fmt"

// https://leetcode.cn/problems/snakes-and-ladders
func main() {
	fmt.Println(snakesAndLadders([][]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 35, -1, -1, 13, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 15, -1, -1, -1, -1}}))
}

type edge struct {
	id      int
	stepNum int
}

func snakesAndLadders(board [][]int) int {
	n := len(board)
	visited := make([]bool, n*n+1)
	queue := []edge{{1, 0}}
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		for i := 1; i <= 6; i++ {
			next := item.id + i
			if next > n*n {
				break
			}
			row, col := findIdRowAndCol(next, n)
			if board[row][col] > -1 {
				next = board[row][col]
			}
			if next == n*n {
				return item.stepNum + 1
			}
			if !visited[next] {
				visited[next] = true
				queue = append(queue, edge{id: next, stepNum: item.stepNum + 1})
			}
		}
	}
	return -1
}

func findIdRowAndCol(id, n int) (row, col int) {
	// 根据每个点的id计算所在的行和列，奇数行从右到左遍历行
	row, col = (id-1)/n, (id-1)%n
	if row%2 == 1 {
		col = n - 1 - col
	}
	row = n - 1 - row
	return
}
