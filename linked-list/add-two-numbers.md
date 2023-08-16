# 环形链表

题目链接: [https://leetcode.cn/problems/add-two-numbers](https://leetcode.cn/problems/add-two-numbers)

## 解题思路：

1. 遍历两个链表，逐个去除元素加到临时的sum，再加上上一轮相加得出的进位
2. 临时sum/10就是进位，临时sum%10就是这一轮的最终结果


```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	current := head
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		current.Next = new(ListNode)
		current.Next.Val = sum % 10
		current = current.Next
	}
	return head.Next
}
```

## 复杂度分析

- **时间复杂度：** 只遍历了一遍字符串，因此时间复杂度为 $$O(n)$$，其中 $$n$$ 是链表的长度
- **空间复杂度：** 只使用了常数个变量，因此空间复杂度为 $$O(n)$$
