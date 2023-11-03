# 矩阵置零

题目链接: [https://leetcode.cn/problems/set-matrix-zeroes](https://leetcode.cn/problems/set-matrix-zeroes)

## 解题思路：

1. 遍历矩阵，找到值为0的元素，将对应的行和列标记为需要置零
2. 遍历需要置零的行和列的记录表，将对应行及对应列的所有元素置零

```go
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
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(m*n)$$,$$m$$为矩阵的行数,$$n$$为矩阵的列数
- **空间复杂度：** 空间复杂度为$$O(m*n)$$,$$m$$为矩阵的行数,$$n$$为矩阵的列数
