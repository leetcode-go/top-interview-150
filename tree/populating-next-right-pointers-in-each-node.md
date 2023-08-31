# 填充每个节点的下一个右侧节点指针

题目链接: [https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii](https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii)

## 解题思路：

1. 采用`bfs`方式遍历每层节点
2. 用`queue`存放每层节点，逐个从头部出队，并将节点的`Next`指向`queue`中的第一个节点
3. 为了区分层别，需要用`size`记录每层节点数

```go
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
    if root==nil{
        return root
    }
    queue:=[]*Node{root}
    for len(queue)>0{
        size:=len(queue)
        for i:=0;i<size;i++{
            node:=queue[0]
            queue=queue[1:]
            if i==size-1{
                node.Next=nil
            }else{
                node.Next=queue[0]
            }
            if node.Left!=nil{
                queue=append(queue,node.Left)
            }
            if node.Right!=nil{
                queue=append(queue,node.Right)
            }
        }
    }
    return root
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n)$$，$$n$$为树的节点数
- **空间复杂度：** 空间复杂度为$$O(n)$$。
