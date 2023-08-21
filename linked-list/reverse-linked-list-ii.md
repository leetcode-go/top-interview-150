# 反转链表 II

题目链接: [https://leetcode.cn/problems/reverse-linked-list-ii](https://leetcode.cn/problems/reverse-linked-list-ii)

## 解题思路：

1. 遍历链表，找到`left`对应的节点的前一个节点`pre`
2. 因为要把`left`到`right`这段翻转，因此我们不能直接把`left`进行移动，而是把`left`的`Next`进行移动，这样我们翻转完`left`的`Next`就是`right`的`Next`
3. 翻转的过程中，不断取出`left`的`Next`，并插入到`pre`的后头


```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    resNode:=new(ListNode)
    resNode.Next=head
    pre:=resNode
    // 查找到left前一个节点
    for i:=0;i<left-1;i++{
        pre = pre.Next
    }
    cur:=pre.Next
    for i:=left; i< right;i++{
        // 不断取出left后一个元素插入到pre后面，并left的next往后移一位
        next := cur.Next
        cur.Next = next.Next
        next.Next = pre.Next
        pre.Next = next
    }
    return resNode.Next
}
```

## 复杂度分析

- **时间复杂度：** 只遍历了一遍链表，因此时间复杂度为 $$O(n)$$，其中 $$n$$ 是链表的长度
- **空间复杂度：** 空间复杂度为 $$O(1)$$
