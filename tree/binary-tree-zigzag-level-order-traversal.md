# 二叉树的锯齿形层序遍历

题目链接: [https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal](https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal)

## 解题思路：

1. 遍历每层节点，并将每个非`nil`的节点存入临时队列，
2. 遍历完当前层节点后，将临时队列中的节点存入`queue`
3. 如果当前层为奇数层，则将当前层的值翻转
4. 将当前层最终的值数组写入到结果数组中

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	level := 0
	res := make([][]int, 0)
	for ; len(queue) > 0; level++ {
		tmpRes := make([]int, 0)
		tmpArr := make([]*TreeNode, 0)
		for _, item := range queue {
			tmpRes = append(tmpRes, item.Val)
			if item.Left != nil {
				tmpArr = append(tmpArr, item.Left)
			}
			if item.Right != nil {
				tmpArr = append(tmpArr, item.Right)
			}

		}
		queue = tmpArr
		if level%2 == 1 {
			for i, n := 0, len(tmpRes); i < n/2; i++ {
				tmpRes[i], tmpRes[n-1-i] = tmpRes[n-1-i], tmpRes[i]
			}
		}
		res = append(res, tmpRes)
	}
	return res
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n)$$
- **空间复杂度：** 空间复杂度为$$O(n)$$，$$n$$的大小为树节点最多的那层的节点数
