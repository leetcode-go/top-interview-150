# 二叉搜索树中第K小的元素

题目链接: [https://leetcode.cn/problems/kth-smallest-element-in-a-bst](https://leetcode.cn/problems/kth-smallest-element-in-a-bst)

## 解题思路：

1. 根据二叉搜索树的特性，中序遍历可以得到一个升序队列
2. 因此只需要中序遍历二茬搜索树，并记录遍历的序号，直至遍历到第k个节点，则直接返回

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    if root==nil{
        return 0
    }
    res,idx:=0,0
    find:=false
    var dfs func(node *TreeNode)
    dfs=func(node *TreeNode){
        if node ==nil{
            return
        }
        dfs(node.Left)
        idx++
        if idx==k{
            res=node.Val
            find=true
        }
        if find{
            return
        }
        dfs(node.Right)
    }
    dfs(root)
    return res
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n)$$
- **空间复杂度：** 空间复杂度为$$O(1)$$
