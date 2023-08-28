# 相同的树

题目链接: [https://leetcode.cn/problems/same-tree](https://leetcode.cn/problems/same-tree)

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
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p==nil&&q==nil{
        return true
    }
    if (p==nil&&q!=nil)||(p!=nil&&q==nil)||p.Val!=q.Val{
        return false
    }
    return isSameTree(p.Left,q.Left)&&isSameTree(p.Right,q.Right)
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n)$$
- **空间复杂度：** 没有使用额外的空间，空间复杂度为$$O(1)$$。
