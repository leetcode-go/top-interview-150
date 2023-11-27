package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/sort-list
func main() {
	data := &ListNode{Val: -1, Next: &ListNode{5, &ListNode{3, &ListNode{4, &ListNode{Val: 0, Next: nil}}}}}
	sortList(data)
	for data != nil {
		println(data.Val)
		data = data.Next
	}
}

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
