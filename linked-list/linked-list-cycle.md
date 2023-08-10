# 环形链表

题目链接: [https://leetcode.cn/problems/linked-list-cycle](https://leetcode.cn/problems/linked-list-cycle)

## 解题思路：

1. 遍历链表每个节点，并记录
2. 若某个节点已经出现过，则认为是环形链表


```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    valMap:=map[*ListNode]bool{}
    for head!=nil{
        if _,ok:=valMap[head];ok{
            return true
        }else{
            valMap[head]=true
            head=head.Next
        }
    }
    return false
}
```

## 复杂度分析

- **时间复杂度：** 只遍历了一遍字符串，因此时间复杂度为 $$O(n)$$，其中 $$n$$ 是链表的长度
- **空间复杂度：** 只使用了常数个变量，因此空间复杂度为 $$O(n)$$
