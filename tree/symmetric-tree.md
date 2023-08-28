# 对称二叉树

题目链接: [https://leetcode.cn/problems/symmetric-tree](https://leetcode.cn/problems/symmetric-tree)

## 解题思路：

1. 分别判断左子树的左节点与右子树的右节点是否相等，左子树的右节点与右子树的左节点是否相等
2. 递归判断

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    return isMirror(root,root)
}

func isMirror(left,right *TreeNode)bool{
    if left==nil&&right==nil{
        return true

    }
    if left==nil||right==nil{
        return false
    }
    return left.Val==right.Val&&isMirror(left.Left,right.Right)&&isMirror(left.Right,right.Left)
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n)$$
- **空间复杂度：** 没有使用额外的空间，空间复杂度为$$O(1)$$。
