# 二叉树中最大路径和

题目链接: [https://leetcode.cn/problems/binary-tree-maximum-path-sum](https://leetcode.cn/problems/binary-tree-maximum-path-sum)

## 解题思路：

1. 每个节点在路径上的最大贡献为该节点的值加上左右子节点的最大贡献值，即`max(val, left, right)`
2. 以每个节点为根节点的子树的最大路径和为该节点值加上左右子节点的最大贡献值，即`left + right + val`（若子节点的最大贡献值为负数不参与计算）
3. 递归计算每个节点的最大路径和，并与全局的`maxSum`进行对比，若大于，则更新`maxSum`

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var maxSum int

func maxPathSum(root *TreeNode) int {
	maxSum = math.MinInt32
	reverseTree(root)
	return maxSum
}

func reverseTree(root *TreeNode)int{
    if root==nil{
        return 0
    }
    left:=max(reverseTree(root.Left),0)
    right:=max(reverseTree(root.Right),0)
    maxSum=max(maxSum,root.Val+left+right)
    return root.Val+max(left,right)
}

func max(x,y int)int{
    if x>y{
        return x
    }
    return y
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n)$$
- **空间复杂度：** 空间复杂度为$$O(1)$$。
