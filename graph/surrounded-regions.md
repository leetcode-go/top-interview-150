# 被围绕的区域

题目链接: [https://leetcode.cn/problems/surrounded-regions](https://leetcode.cn/problems/surrounded-regions)

## 解题思路：

1. 在每个矩阵的元素上遍历上下左右四个方向的元素，可以识别出与之相连的区域
2. 对于被'X'包围的'O'的区域，要将'O'替换成'X'，因此只有从边界线上出发的'O'才不需要替换
3. 从四条边界的'O'出发，将与之相连的'O'替换成'A'，这样就将不需要标记成'X'的'O'与需要标记成'X'的'O'区分开来了
4. 再次遍历整个矩阵，将所有标记成'A'的元素替换成'O'，将所有标记成'O'的元素替换成'X'

```go
func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != 'O' {
			return
		}
		board[i][j] = 'A'
		dfs(i, j+1)
		dfs(i, j-1)
		dfs(i+1, j)
		dfs(i-1, j)
	}
	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for i := 1; i < n-1; i++ {
		dfs(0, i)
		dfs(m-1, i)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n*m)$$,$$m$$为矩阵的行数，$$n$$为矩阵的列数
- **空间复杂度：** 空间复杂度为$$O(m*n)$$,堆栈消耗
