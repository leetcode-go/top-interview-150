# 螺旋矩阵

题目链接: [https://leetcode.cn/problems/spiral-matrix](https://leetcode.cn/problems/spiral-matrix)

## 解题思路：

1. 用数组存储每个方向要变动的`x`、`y`，以及变动的大小，遍历矩阵
2. 碰到边界或者同方向上下一个位置已经被遍历过，则改变方向
3. 记录改变方向的次数，对4取余，获取这次该往哪个方向变动，每个遍历过的位置都标记为-101(矩阵中的值不超过-100-100)
4. 直至所有元素都被遍历完成

```go
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	var direction = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var i, j, m, n = 0, 0, len(matrix), len(matrix[0])
    var result =make([]int,m*n)
	var direct = 0
	for idx:=0;idx < m*n;idx++ {
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
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(m*n)$$,$$n$$为矩阵的行数，$$m$$为矩阵的列数
- **空间复杂度：** 空间复杂度为$$O(m*n)$$,$$n$$为矩阵的行数，$$m$$为矩阵的列数
