# 二叉搜索树迭代器

题目链接: [https://leetcode.cn/problems/binary-search-tree-iterator](https://leetcode.cn/problems/binary-search-tree-iterator)

## 解题思路：

1. 中序遍历二叉树，构建成数组，执行操作时直接读取数组即可

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	data []int
}

func Constructor(root *TreeNode) BSTIterator {
	iter := BSTIterator{data: []int{}}
	iter.reverse(root)
	return iter
}

func (this *BSTIterator) reverse(root *TreeNode) {
	if root == nil {
		return
	}
	this.reverse(root.Left)
	this.data = append(this.data, root.Val)
	this.reverse(root.Right)
}

func (this *BSTIterator) Next() int {
	data := this.data[0]
	this.data = this.data[1:]
	return data
}

func (this *BSTIterator) HasNext() bool {
	return len(this.data)> 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
```

## 复杂度分析

- **时间复杂度：** `Next`及`HasNext`操作的时间复杂度为$$O(1)$$
- **空间复杂度：** 空间复杂度为$$O(n)$$。
