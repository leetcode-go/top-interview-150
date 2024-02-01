package main

import (
	"fmt"
)

func main() {
	fmt.Println(totalNQueens(4))

}

func totalNQueens(n int) (res int) {
	columns, diagonals1, diagonals2 := make([]bool, n), make([]bool, 2*n-1), make([]bool, 2*n-1)
	var backtrack func(int)
	backtrack = func(row int) {
		if row == n {
			res++
			return
		}
		for col, hasQueen := range columns {
			// d1增加了一个n-1,当row-col为负数时,通过哈希,转换为正值,避免数组溢出
			d1, d2 := row+n-1-col, col+row
			if hasQueen || diagonals1[d1] || diagonals2[d2] {
				continue
			}
			columns[col], diagonals1[d1], diagonals2[d2] = true, true, true
			backtrack(row + 1)
			columns[col], diagonals1[d1], diagonals2[d2] = false, false, false
		}
	}
	backtrack(0)
	return
}
