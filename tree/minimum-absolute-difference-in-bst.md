# 二叉搜索树的最小绝对差

题目链接: [https://leetcode.cn/problems/minimum-absolute-difference-in-bst](https://leetcode.cn/problems/minimum-absolute-difference-in-bst)

## 解题思路：

1. 根据二叉搜索树的特性，中序遍历可以得到一个升序队列，因此我们可以用一个变量记录前一个节点的值，然后遍历队列，计算当前节点与前一个节点的差值，取最小值即可。(下方代码采用的官方解法，原理相同，优化了空间使用情况)

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getMinimumDifference(root *TreeNode) int {
	res, pre := math.MaxInt, -1
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 {
			res = min(node.Val-pre, res)
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return res
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n)$$
- **空间复杂度：** 空间复杂度为$$O(n)$$，$$n$$的大小为树的节点数
