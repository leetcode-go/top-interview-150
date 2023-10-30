# 有效的数独

题目链接: [https://leetcode.cn/problems/valid-sudoku](https://leetcode.cn/problems/valid-sudoku)

## 解题思路：

1. 遍历矩阵的每个元素，并记录非空的元素，在其所在行、所在列、所在小九宫格的记录中判断是否存在过
2. 若不存在则记录，若存在则返回false

```go
func isValidSudoku(board [][]byte) bool {
	var coll, row [9][9]bool
	var subBox [3][3][9]bool
	for i, rowItem := range board {
		for j, cell := range rowItem {
			if cell == '.' {
				continue
			}
			charIdx := cell - '1'
			if row[i][charIdx] || coll[j][charIdx] || subBox[i/3][j/3][charIdx] {
				return false
			}
			row[i][charIdx] = true
			coll[j][charIdx] = true
			subBox[i/3][j/3][charIdx] = true
		}
	}
	return true
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n)$$,$$n$$为矩阵中所有字符数
- **空间复杂度：** 空间复杂度为$$O(n)$$,$$n$$为矩阵中所有字符数
