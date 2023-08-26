# 删除链表的倒数第 N 个结点

题目链接: [https://leetcode.cn/problems/remove-nth-node-from-end-of-list](https://leetcode.cn/problems/remove-nth-node-from-end-of-list)

## 解题思路一

1. 遍历链表，把每个节点都存入数组
2. 根据入参n，找到倒数第n个节点
3. 将第n-1个节点的Next指向第n+1个节点


```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    list:=make([]*ListNode,0)
    l:=head
    for l!=nil{
        list=append(list,l)
        l=l.Next
    }
    if n==len(list){
        return head.Next
    }
    idx:=len(list)-n
    pre:=idx-1
    list[pre].Next=list[idx].Next
    return head
}
```

## 复杂度分析一

- **时间复杂度：** 只遍历了一遍链表，因此时间复杂度为 $$O(n)$$，其中 $$n$$ 是链表的长度
- **空间复杂度：** 空间复杂度为 $$O(n)$$

## 解题思路二

1. 遍历链表，计算出长度
2. 再次遍历链表到倒数第n+1个节点，将倒数第n+1个节点的Next指向倒数第n个节点的Next


```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    length:=0
    l:=head
    for ;l!=nil;l=l.Next{
        length++
    }
    l=&ListNode{0,head}
    cur:=l
    for i:=0;i<length-n;i++{
        cur=cur.Next
    }
    cur.Next=cur.Next.Next
    return l.Next
}
```

## 复杂度分析二

- **时间复杂度：** 只遍历了一遍链表，因此时间复杂度为 $$O(n)$$，其中 $$n$$ 是链表的长度
- **空间复杂度：** 空间复杂度为 $$O(1)$$
