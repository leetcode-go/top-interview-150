# 二叉树的最大深度

题目链接: [https://leetcode.cn/problems/maximum-depth-of-binary-tree](https://leetcode.cn/problems/maximum-depth-of-binary-tree)

## 解题思路：

1. 深度遍历二叉树，获取每个节点的左右子树的深度，逐层返回每个左右子树的较大的深度

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root==nil{
        return 0
    }
    return 1+max(maxDepth(root.Left),maxDepth(root.Right))
}

func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}
```

## 复杂度分析

- **时间复杂度：** 对二树进行了深度遍历，因此时间复杂度为$$O(log n)$$
- **空间复杂度：** 没有使用额外的空间，空间复杂度为$$O(1)$$。
