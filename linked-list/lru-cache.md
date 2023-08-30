# LRU缓存

题目链接: [https://leetcode.cn/problems/lru-cache](https://leetcode.cn/problems/lru-cache)

## 解题思路：

1. 用链表存储数据，同时采用`map`记录每个`key`对应的`node`
2. 读取数据的时候，将对应的`node`从原位置移除，插入到链表头部
3. 新增节点时，判断`map`长度，超过`capacity`则要在末尾进行移除
4. 新节点或刷新节点均需要在链表头部插入


```go
type Node struct{
    Key int
    Value int
    Next,Pre *Node
}

type LRUCache struct {
    head *Node
    last *Node
    history map[int]*Node
    capacity int
}


func Constructor(capacity int) LRUCache {
    l:=LRUCache{
        head:&Node{},
        last:&Node{},
        history:map[int]*Node{},
        capacity:capacity,
    }
    l.head.Next=l.last
    l.last.Pre=l.head
    return l
}


func (this *LRUCache) Get(key int) int {
    if node,ok:=this.history[key];ok{
        this.deleteNode(node)
        this.insertHead(node)
        return node.Value
    }
    return -1
}

func (this *LRUCache)insertHead(node *Node){
    this.head.Next.Pre=node
    node.Next=this.head.Next
    node.Pre=this.head
    this.head.Next=node
}

func(this *LRUCache)deleteNode(node *Node){
    node.Pre.Next=node.Next
    node.Next.Pre=node.Pre
}


func (this *LRUCache) Put(key int, value int)  {
    if node,ok:=this.history[key];ok{
        node.Value=value
        this.deleteNode(node)
        this.insertHead(node)
    }else{
        if this.capacity==len(this.history){
            rmNode:=this.last.Pre
            this.deleteNode(rmNode)
            delete(this.history,rmNode.Key)
        }
        node=&Node{Key:key,Value:value}
        this.insertHead(node)
        this.history[key]=node
    }
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为 $$O(1)$$
- **空间复杂度：** 空间复杂度为 $$O(n)$$
