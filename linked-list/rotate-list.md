# 旋转链表

题目链接: [https://leetcode.cn/problems/rotate-list](https://leetcode.cn/problems/rotate-list)

## 解题思路一

1. 遍历链表，将所有节点放入数组中
2. 再把数组与数组自身进行拼接，行程头尾相连
3. 从length-k的位置开始截取length长度的数组
4. 对k-1位置及数组末尾的Next进行修正


```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if k==0||head == nil || head.Next == nil {
		return head
	}
	listHead := make([]*ListNode, 0)
	l := head
	for ; l != nil; l = l.Next {
		listHead = append(listHead, l)
	}
	length := len(listHead)
	k = k % length
    if k==0{
        return head
    }
	listHead = append(listHead, listHead...)
	listHead = listHead[length-k : length*2-k]
	listHead[k-1].Next = listHead[k]
	listHead[length-1].Next = nil
	return listHead[0]
}
```

## 复杂度分析一

- **时间复杂度：** 只遍历了一遍链表，因此时间复杂度为 $$O(n)$$，其中 $$n$$ 是链表的长度
- **空间复杂度：** 空间复杂度为 $$O(2n)$$


## 解题思路二

1. 将链表头尾相连，组成环形链表
2. 找到第length-k个节点，新的链表的head为该节点的Next，将该节点Next修正为nil


```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	length := 1
	l := head
	for ; l.Next != nil; l = l.Next {
		length++
	}
	k = k % length
	add := length - k
	l.Next = head
	for add > 0 {
		l = l.Next
		add--
	}
	res := l.Next
	l.Next = nil
	return res
}
```

## 复杂度分析二

- **时间复杂度：** 只遍历了一遍链表，因此时间复杂度为 $$O(n)$$，其中 $$n$$ 是链表的长度
- **空间复杂度：** 空间复杂度为 $$O(2n)$$