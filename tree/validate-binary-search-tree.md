# 验证二叉搜索树

题目链接: [https://leetcode.cn/problems/validate-binary-search-tree](https://leetcode.cn/problems/validate-binary-search-tree)

## 解题思路：

1. 根据二叉搜索树的特性，中序遍历可以得到一个升序队列
2. 因此只需要中序遍历二叉搜索树，只要发现前序节点大于等于当前节点，则非二叉搜索树，直接中断递归返回结果

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	isBST, pre := true, math.MinInt
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != math.MinInt && pre >= node.Val {
			isBST = false
			return
		}
        if isBST==false{
            return
        }
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return isBST
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n)$$
- **空间复杂度：** 空间复杂度为$$O(1)$$
