# 求根节点到叶节点数字之和

题目链接: [https://leetcode.cn/problems/sum-root-to-leaf-numbers](https://leetcode.cn/problems/sum-root-to-leaf-numbers)

## 解题思路：

1. 深度遍历，计算出叶子节点的值，将叶子节点值计算入和

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	return sum(root, 0)
}

func sum(root *TreeNode, lastNum int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return lastNum*10 + root.Val
	}
	levelVal := lastNum*10 + root.Val
	num := sum(root.Left, levelVal) + sum(root.Right, levelVal)
	return num
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n)$$
- **空间复杂度：** 空间复杂度为$$O(n)$$。
