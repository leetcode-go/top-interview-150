# 合并两个有序链表

题目链接: [https://leetcode.cn/problems/merge-two-sorted-lists](https://leetcode.cn/problems/merge-two-sorted-lists)

## 解题思路：

1. 同时遍历两个链表，判断哪个链表的元素小，小的元素入结果链表中，并将结果链表以及小的元素所在链表指针均向后移一位
2. 若某一链表为空，则将另一个链表剩余的元素一并合并到结果链表zhong


```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    res:=new(ListNode)
    head:=res
    for list1!=nil&&list2!=nil{
        if list1.Val>list2.Val{
            head.Next=list2
            list2=list2.Next
        }else{
            head.Next=list1
            list1=list1.Next
        }
        head=head.Next
    }
    if list1==nil{
        head.Next=list2
    }else{
        head.Next=list1
    }
    return res.Next
}
```

## 复杂度分析

- **时间复杂度：** 只遍历了一遍链表，因此时间复杂度为 $$O(n)$$，其中 $$n$$ 是链表的长度
- **空间复杂度：** 空间复杂度为 $$O(n)$$
