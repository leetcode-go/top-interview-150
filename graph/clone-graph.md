# 克隆图

题目链接: [https://leetcode.cn/problems/clone-graph](https://leetcode.cn/problems/clone-graph)

## 解题思路：

1. 深度遍历图，从`queue`中取出待遍历节点，对每个节点的`Neighbors`进行遍历，对于每个`neighbor`
2. 判断每个`neighbor`是否为之前遍历过，若遍历过，则从`tmpMap`中取出克隆节点，加入临时的`neighbors`中
3. 否则，构建一个克隆节点，并存入`tmpMap`及`queue`中，加入`neighbors`中
4. 将临时`neighbors`赋值给当前的节点在`tmpMap`中的克隆节点
5. 直至`queue`为空

```go
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node==nil{
        return nil
    }
	res := &Node{Val: node.Val}
	tmpMap := map[*Node]*Node{node.Val: res}
	queue := []*Node{node}
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		neighbors := make([]*Node, len(item.Neighbors))
		for idx, sub := range item.Neighbors {
			if tmpMap[sub.Val] != nil {
				neighbors[idx] = tmpMap[sub.Val]
			} else {
				subNode := &Node{Val: sub.Val}
				tmpMap[subNode.Val] = subNode
				neighbors[idx] = subNode
				queue = append(queue, sub)
			}
		}
        tmpMap[item.Val].Neighbors=neighbors
	}
	return res
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n)$$,$$n$$为图的节点个数
- **空间复杂度：** 空间复杂度为$$O(n)$$,$$n$$为图的节点个数
