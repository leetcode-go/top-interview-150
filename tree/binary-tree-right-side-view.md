# 二叉树的右视图

题目链接: [https://leetcode.cn/problems/binary-tree-right-side-view](https://leetcode.cn/problems/binary-tree-right-side-view)

## 解题思路：

1. 对树进行层序遍历，从右侧能看到的节点遍为每层的最后一个节点

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	queue := []*TreeNode{root}
	levelLast := 1
	for len(queue) > 0 {
		item := queue[0]
		levelLast--
		if levelLast == 0 {
			res = append(res, item.Val)
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
		}
	}
	return res
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n)$$
- **空间复杂度：** 空间复杂度为$$O(n)$$，$$n$$的大小为树节点最多的那层的节点数
