# 二叉树的最近公共祖先

题目链接: [https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree](https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree)

## 解题思路：

1. 最近公共祖先有两种情况
    - 当前节点`p`或`q`为其中之一
    - p或q归属于同一个棵子树
2. 从`root`开始遍历树，判断当前节点的值是否与`p`或`q`相等，若相等则返回当前节点，若不相等，则从左右子树中查找
3. 若左右子树中各自查找到`p`或`q`，则说明当前节点为最近公共祖先
4. 若左子树中未查找到`p`或`q`，则说明`p`为`q`的最近公共祖先或`q`为`p`的最近公共祖先，返回从子树中查找到的子节点

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	//
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	// 查找左子树中是否存在p或q
	leftCheck := lowestCommonAncestor(root.Left, p, q)
	// 查找右子树中是否存在p或q
	rightCheck := lowestCommonAncestor(root.Right, p, q)
	// 如果左子树与右子树均查找到p或q，则p与q的公共祖先节点为当前节点，且当前节点必然是最近的公共祖先节点
	if leftCheck != nil && rightCheck != nil {
		return root
	}
	// 左子树未查找到p或q则两个节点的公共祖先节点为右子树中查找到的节点
	if leftCheck == nil {
		return rightCheck
	}
	return leftCheck
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n)$$
- **空间复杂度：** 空间复杂度为$$O(1)$$。
