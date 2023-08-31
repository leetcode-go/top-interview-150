# 二叉树展开为链表

题目链接: [https://leetcode.cn/problems/flatten-binary-tree-to-linked-list](https://leetcode.cn/problems/flatten-binary-tree-to-linked-list)

## 解题思路：

1. 按先序遍历树，节点存放到数组
2. 遍历数组，将节点连接起来

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    list:=preTravel(root)
    for i:=1;i<len(list);i++{
        pre,node:=list[i-1],list[i]
        pre.Right=node
        pre.Left=nil
    }
}

func preTravel(root *TreeNode)[]*TreeNode{
    list:=[]*TreeNode{}
    if root!=nil{
        list=append(list,root)
        list=append(list,preTravel(root.Left)...)
        list=append(list,preTravel(root.Right)...)
    }
    return list
}
```

## 复杂度分析

- **时间复杂度：** 遍历了数组重所有节点，因此时间复杂度为$$O(n)$$
- **空间复杂度：** 空间复杂度为$$O(n)$$。
