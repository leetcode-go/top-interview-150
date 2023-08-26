# 删除排序链表中的重复元素 II

题目链接: [https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii](https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii)

## 解题思路

1. 遍历链表，对比每个节点与下一个节点的值是否一致，若是则标记这个值为重复值，跳过这个节点
2. 后续节点需要与该重复值进行对比是否相等，若是则也认为重复值，进行跳过


```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	last := &ListNode{-101, nil}
	res := last
	tmp := -101
	for cur != nil {
		if cur.Next != nil && cur.Val == cur.Next.Val {
			tmp = cur.Val
		} else {
			if cur.Val != tmp {
				last.Next = &ListNode{cur.Val, nil}
				last = last.Next
			}
		}
		cur = cur.Next
	}
	return res.Next
}
```

## 复杂度分析

- **时间复杂度：** 只遍历了一遍链表，因此时间复杂度为 $$O(n)$$，其中 $$n$$ 是链表的长度
- **空间复杂度：** 空间复杂度为 $$O(n)$$