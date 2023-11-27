# 排序链表

题目链接: [https://leetcode.cn/problems/sort-list](https://leetcode.cn/problems/sort-list)

## 解题思路：

1. 将队列从中间一分为二，分割方法用两个指针指向链表头，一个向后遍历一个元素一个向后遍历两个元素，当遍历到尾部时，另一个指针指向的链表就是后半部分的头元素
2. 对拆分后的链表继续拆分，直至到单个元素，再两两进行排序合并，由小到大升序合并，得出多个相对有序链表
3. 合并到最后，将两个相对有序的链表合并成一个完整的有序链表并返回

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	return split(head, nil)
}

func split(head, end *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == end {
		head.Next = nil
		return head
	}
	slow, fast := head, head
	for fast != end {
		slow = slow.Next
		fast = fast.Next
		if fast != end {
			fast = fast.Next
		}
	}
	mid := slow
	return merge(split(head, mid), split(mid, end))
}

func merge(head1, head2 *ListNode) *ListNode {
	rootHead := &ListNode{}
	tmp, l1, l2 := rootHead, head1, head2
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}
	if l1 != nil {
		tmp.Next = l1
	} else if l2 != nil {
		tmp.Next = l2
	}
	return rootHead.Next
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n*log n)$$，$$n$$为链表长度
- **空间复杂度：** 空间复杂度为$$O(log n)$$，$$n$$为链表长度
