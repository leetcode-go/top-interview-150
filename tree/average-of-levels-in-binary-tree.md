# 二叉树的层平均值

题目链接: [https://leetcode.cn/problems/average-of-levels-in-binary-tree](https://leetcode.cn/problems/average-of-levels-in-binary-tree)

## 解题思路：

1. 对树进行层序遍历，并对每层节点的值进行求平均值

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	res := make([]float64, 0)
	queue := []*TreeNode{root}
	levelLast := 1
	sum := 0
	levelCount := float64(1)
	for len(queue) > 0 {
		item := queue[0]
		sum += item.Val
		levelLast--
		if levelLast == 0 {
			res = append(res, float64(sum)/levelCount)
			levelCount = 0
			sum = 0
		}
		queue = queue[1:]
		if item.Left != nil {
			queue = append(queue, item.Left)
		}
		if item.Right != nil {
			queue = append(queue, item.Right)
		}
		if levelLast == 0 {
			levelLast = len(queue)
			levelCount = float64(levelLast)
		}
	}
	return res
}

```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n)$$
- **空间复杂度：** 空间复杂度为$$O(n)$$，$$n$$的大小为树节点最多的那层的节点数
