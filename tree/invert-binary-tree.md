# 翻转二叉树

题目链接: [https://leetcode.cn/problems/invert-binary-tree](https://leetcode.cn/problems/invert-binary-tree)

## 解题思路：

1. 深度遍历二叉树，判断每个节点及其子节点的值是否都相同

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root==nil{
        return root
    }
    root.Left,root.Right=invertTree(root.Right),invertTree(root.Left)
    return root
}
```

## 复杂度分析

- **时间复杂度：** 深度遍历二叉树所有的节点，时间复杂度为$$O(n)$$
- **空间复杂度：** 没有使用额外的空间，空间复杂度为$$O(1)$$。
