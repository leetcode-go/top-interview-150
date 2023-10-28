# 蛇梯棋

题目链接: [https://leetcode.cn/problems/snakes-and-ladders](https://leetcode.cn/problems/snakes-and-ladders)

## 解题思路：

1. 从1开始遍历到n\*n，每个数字代表棋盘的一个位置，通过数字可以计算出这个位置的行号以及列号
2. 通过行号可以判断是从左到右遍历还是从右到左遍历
3. 从该数字开始往后推算接下来的6个数字对应的棋盘位置，以及上面是否有蛇或梯子，判断出若下一步走这个位置最终能到棋盘的哪个位置
4. 并将每个位置的下一个位置及到该位置使用的步数存入队列，再按照步骤3遍历队列里的元素
5. 直至到终点

```go
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
        for i:=1;i<=6;i++{
		next := item.id + i
		if next > n*n {
			break
		}
		row, col := findIdRowAndCol(next, n)
		if board[row][col] > 0 {
			next = board[row][col]
		}
		if next == n*n {
			return item.stepNum+1
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

```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n^2)$$,$$n$$为`board`的边长
- **空间复杂度：** 空间复杂度为$$O(n^2)$$,$$n$$为`board`的边长
