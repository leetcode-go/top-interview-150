# 完全二叉树的节点个数

题目链接: [https://leetcode.cn/problems/count-complete-tree-nodes](https://leetcode.cn/problems/count-complete-tree-nodes)

## 解题思路：

1. 完全二叉树，每层节点均是从左往右填充，中间不会存在空缺，可以直接通过广度遍历，逐层统计节点数
2. 从左往右逐层遍历节点，并计数，若节点存在左子树则左子树加入队列，否则直接跳过进入下一个节点的遍历
3. 若存在左子树，则将左子树加入到队列后判断是否存在右子树，存在则加入队列

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
  if root==nil{
    return 0
  }
	queue := []*TreeNode{root}
	count := 0
	for len(queue) > 0 {
		node := queue[0]
		count += 1
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}else{
            continue
        }
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return count
}
```

## 复杂度分析

- **时间复杂度：** 遍历了数组中所有节点，因此时间复杂度为$$O(n)$$
- **空间复杂度：** 空间复杂度为$$O(n)$$。
