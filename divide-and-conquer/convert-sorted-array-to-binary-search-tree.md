# 将有序数组转换为二叉搜索树

题目链接: [https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree](https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree)

## 解题思路：

1. 根据二叉搜索树的特性，根节点为数组中间的元素，左子树为小于根节点的元素，右子树为大于根节点的元素
2. 数组的`mid`为二叉搜索树的根节点，`mid`左边的元素为左子树，`mid`右边的元素为右子树
3. 递归遍历左右子树，构建相应的根节点，直至`left`大于`right`

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    return buildTree(nums,0,len(nums)-1)
}

func buildTree(nums []int,left,right int)*TreeNode{
    if left>right{
        return nil
    }
    mid:=(left+right)/2
    node:=&TreeNode{Val:nums[mid]}
    node.Right=buildTree(nums,mid+1,right)
    node.Left=buildTree(nums,left,mid-1)
    return node
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n)$$,$$n$$为数组长度
- **空间复杂度：** 空间复杂度为$$O(log n)$$,$$n$$为数组长度
