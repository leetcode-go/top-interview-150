# 分隔链表

题目链接: [https://leetcode.cn/problems/partition-list](https://leetcode.cn/problems/partition-list)

## 解题思路

1. 遍历链表，将链表分成`<=x`和`>x`两部分
2. 合并`<=x`及`>x`两部分的链表，并返回


```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    l1,l2:=new(ListNode),new(ListNode)
    h1,h2:=l1,l2
    for head!=nil{
        if head.Val<x{
            h1.Next=head
            h1=h1.Next
        }else{
            h2.Next=head
            h2=h2.Next
        }
        head=head.Next
    }
    h1.Next=l2.Next
    h2.Next=nil
    return l1.Next
}
```

## 复杂度分析

- **时间复杂度：** 只遍历了一遍链表，因此时间复杂度为 $$O(n)$$，其中 $$n$$ 是链表的长度
- **空间复杂度：** 空间复杂度为 $$O(n)$$
