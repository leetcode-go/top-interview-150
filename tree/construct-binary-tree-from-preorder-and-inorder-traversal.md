# 从前序和中序遍历序列构造二叉树

题目链接: [https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal](https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal)

## 解题思路：

1. 递归构建，先序遍历的第一个节点为根节点，在中序遍历中找到根节点的位置，根节点左边的为左子树，右边的为右子树
2. 再分别递归构建左子树和右子树

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder)==0{
        return  nil
    }
    root:=&TreeNode{preorder[0],nil,nil}
    i:=0
    for ;i<len(inorder);i++{
        if inorder[i]==preorder[0]{
            break
        }
    }
    root.Left=buildTree(preorder[1:len(inorder[:i])+1],inorder[:i])
    root.Right=buildTree(preorder[len(inorder[:i])+1:],inorder[i+1:])
    return root
}
```

## 复杂度分析

- **时间复杂度：** 遍历了数组重所有节点，因此时间复杂度为$$O(n)$$
- **空间复杂度：** 空间复杂度为$$O(n)$$。
