# 二叉树的层序遍历

题目链接: [https://leetcode.cn/problems/binary-tree-level-order-traversal](https://leetcode.cn/problems/binary-tree-level-order-traversal)

## 解题思路：

1. 遍历每层节点，并将每个非`nil`的节点存入临时队列，
2. 遍历完当前层节点后，将临时队列中的节点存入`queue`，并将当层节点的值存入`result`中

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		tmpArr := make([]*TreeNode, 0)
		tmpRes := make([]int, len(queue))
		for idx, node := range queue {
			tmpRes[idx] = node.Val
			if node.Left != nil {
				tmpArr = append(tmpArr, node.Left)
			}
			if node.Right != nil {
				tmpArr = append(tmpArr, node.Right)
			}
		}
		res = append(res, tmpRes)
		queue = tmpArr
	}
	return res
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n)$$，$$n$$的大小为树的节点数量
- **空间复杂度：** 空间复杂度为$$O(n)$$
