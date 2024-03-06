# 建立四叉树

题目链接: [https://leetcode.cn/problems/construct-quad-tree](https://leetcode.cn/problems/construct-quad-tree)

## 解题思路：

1. 遍历矩阵，判断是否所有元素都相同，若出现不同，则直接将矩阵一分为四，分别再判断分裂的矩阵是否所有元素都相同
2. 若所有元素都相同，则构建一个叶子节点，并将该节点作为根节点返回

```go
/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */


func construct(grid [][]int) *Node {
	var dfs func([][]int, int, int) *Node
	dfs = func(rows [][]int, c0, c1 int) *Node {
		for _, row := range rows {
			for _, v := range row[c0:c1] {
				if v != rows[0][c0] {
					rMid, cMid := len(rows)/2, (c0+c1)/2
					return &Node{
						true,
						false,
						dfs(rows[:rMid], c0, cMid),
						dfs(rows[:rMid], cMid, c1),
						dfs(rows[rMid:], c0, cMid),
						dfs(rows[rMid:], cMid, c1),
					}
				}
			}
		}
		return &Node{Val: rows[0][c0] == 1, IsLeaf: true}
	}
	return dfs(grid, 0, len(grid))
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(log n)$$,$$n$$为数组行数
- **空间复杂度：** 空间复杂度为$$O(log n)$$,$$n$$为数组行数
