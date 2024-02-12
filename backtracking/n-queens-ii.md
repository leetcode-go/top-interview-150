# n皇后问题II

题目链接: [https://leetcode.cn/problems/n-queens-ii](https://leetcode.cn/problems/n-queens-ii)

## 解题思路：

1. 遍历每一行的各个点，判断每一个同列、左上向右下的斜线以及左下向右上的斜线是否存在点，若不存在则继续下一行
2. 左上向右下的斜线规律为行下标与列下标之差相等，左下向右上的规律为行下标与列下标只和相等

```go
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

```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n!)$$
- **空间复杂度：** 空间复杂度为$$O(n!)$$
