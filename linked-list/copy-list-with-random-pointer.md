# 复制带随机指针的链表

题目链接: [https://leetcode.cn/problems/copy-list-with-random-pointer](https://leetcode.cn/problems/copy-list-with-random-pointer)

## 解题思路：

1. 遍历链表，对每个节点，从cache中查找，若不存在则构建新节点，并存储到cache中，若存在则从cache中直接读取
2. 对于Random的节点，从cache中查找，若cache中不存在Random节点，则构建Random节点，并存储到cache中


```go
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    resNode:=new(Node)
    resHead:=resNode
    cacheNode:=map[*Node]*Node{}
    for item:=head;item!=nil;item=item.Next{
        tmp:=buildNode(item.Val)
        if exist,ok:=cacheNode[item];ok{
            tmp=exist
        }
        cacheNode[item]=tmp
        if item.Random!=nil{
            if exist,ok:=cacheNode[item.Random];ok{
                tmp.Random=exist
            }else{
                tmp.Random=buildNode(item.Random.Val)
                cacheNode[item.Random]=tmp.Random
            }
        }
        
        resHead.Next=tmp
        resHead=resHead.Next
    }
    
    return resNode.Next
}

func buildNode(val int)*Node{
    return &Node{
        Val: val,
    }
}
```

## 复杂度分析

- **时间复杂度：** 只遍历了一遍链表，因此时间复杂度为 $$O(n)$$，其中 $$n$$ 是链表的长度
- **空间复杂度：** 空间复杂度为 $$O(n)$$
