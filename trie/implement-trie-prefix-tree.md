# 实现 Trie (前缀树)

题目链接: [https://leetcode.cn/problems/implement-trie-prefix-tree](https://leetcode.cn/problems/implement-trie-prefix-tree)

## 解题思路：

1. 按照每个字符串内字符出现的先后顺序构建字符串树
2. `search`函数遍历待查找字符串内的所有字符，遍历是否存在一条路线能构成待查找字符串
3. 若能则返回末端节点，否则返回`nil`

```go
type Trie struct {
	subChar []*Trie
	end     bool
}

func Constructor() Trie {
	return Trie{
		subChar: make([]*Trie, 26),
	}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, char := range word {
		idx := char - 'a'
		if node.subChar[idx] == nil {
			sub := Constructor()
			node.subChar[idx] = &sub
		}
		node = node.subChar[idx]
	}
	node.end = true
}

func (this *Trie) search(word string) *Trie {
	node := this
	for _, char := range word {
		idx := char - 'a'
		if node.subChar[idx] == nil {
			return nil
		}
		node = node.subChar[idx]
	}
	return node
}
func (this *Trie) Search(word string) bool {
	node := this.search(word)
	if node == nil || !node.end {
		return false
	}
	return true
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.search(prefix) != nil
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n)$$,$$n$$为字符串长度
- **空间复杂度：** 空间复杂度为$$O(n)$$,$$n$$为所有字符串的字符顺序数
