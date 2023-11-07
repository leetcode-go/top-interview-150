# 生命矩阵

题目链接: [https://leetcode.cn/problems/game-of-life](https://leetcode.cn/problems/game-of-life)

## 解题思路：

1. 遍历矩阵，统计每个位置周边八个方向中活细胞及死细胞数量
2. 按照条件及统计结果，判断当前细胞是否存活，若为死细胞复活则标记为2，若为活细胞死亡则标记为-1(步骤1的统计中，值为-1、1的则认为是活细胞，值为2、0的则认为是死细胞)
3. 遍历矩阵，将标记为2的活细胞标记为1，将标记为-1的死细胞标记为0

```go
func gameOfLife(board [][]int) {
	if board == nil {
		return
	}
	directs := [][]int{{1, 0}, {1, 1}, {0, 1}, {-1, 0}, {0, -1}, {-1, -1}, {1, -1}, {-1, 1}}
	m, n := len(board), len(board[0])
	for i, row := range board {
		for j, cell := range row {
			die, live := 0, 0
			for _, direct := range directs {
				x, y := i+direct[0], j+direct[1]
				if x >= 0 && x < m && y >= 0 && y < n {
					if board[x][y] == 1 || board[x][y] == -1 {
						live += 1
					} else if board[x][y] == 0 || board[x][y] == 2 {
						die += 1
					}
				}
			}
			if cell == 0 && live == 3 {
				board[i][j] = 2
				continue
			}
			if cell == 1 && (live < 2 || live > 3) {
				board[i][j] = -1
			}
		}
	}
	for i, row := range board {
		for j := range row {
			if board[i][j] == 2 {
				board[i][j] = 1
			} else if board[i][j] == -1 {
				board[i][j] = 0
			}
		}
	}
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(m*n*8)$$,$$m$$为矩阵的行数,$$n$$为矩阵的列数
- **空间复杂度：** 空间复杂度为$$O(m*n*8)$$,$$m$$为矩阵的行数,$$n$$为矩阵的列数
