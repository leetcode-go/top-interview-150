# 岛屿数量

题目链接: [https://leetcode.cn/problems/number-of-islands](https://leetcode.cn/problems/number-of-islands)

## 解题思路：

1. 遍历矩阵所有元素，若元素的值为'1'，则为陆地，并以该元素为起点，遍历它的四个方向，若为'1'，则为陆地，将值变更为'2'，并继续遍历，若不为'1'，则停止遍历，返回
2. 每个值为'1'的元素为起点的深度遍历完成，表明一块陆地已经被遍历完成，继续遍历矩阵查找下一个陆地，直至矩阵中所有元素都被遍历完成

```go
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] != '1' {
			return
		}
		grid[i][j] = '2'
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				res++
			}
		}

	}
	return res
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n*m)$$,$$m$$为矩阵的行数，$$n$$为矩阵的列数
- **空间复杂度：** 空间复杂度为$$O(1)$$
