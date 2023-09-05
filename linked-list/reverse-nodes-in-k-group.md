# K个一组翻转链表

题目链接: [https://leetcode.cn/problems/reverse-nodes-in-k-group](https://leetcode.cn/problems/reverse-nodes-in-k-group)

## 解题思路：

1. 遍历链表，把链表的每个节点取出存放到数组中
2. 遍历数组，把数组中每k个节点位置翻转
3. 遍历数组，把数组中每个节点重新存放到链表中


```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	list := []*ListNode{}
	for l := head; l != nil; l = l.Next {
		list = append(list, l)
	}
	for i := 0; i+k <= len(list); i += k {
		for j := 0; j < k/2; j++ {
			list[i+j], list[i+k-j-1] = list[i+k-j-1], list[i+j]
		}
	}
	res := &ListNode{}
	cur := res
	for i := 0; i < len(list); i++ {
		cur.Next = list[i]
		cur = cur.Next
	}
	cur.Next = nil
	return res.Next
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为 $$O(n)$$，其中 $$head$$ 是链表的长度，只对链表进行了一重遍历
- **空间复杂度：** 空间复杂度为 $$O(n)$$
