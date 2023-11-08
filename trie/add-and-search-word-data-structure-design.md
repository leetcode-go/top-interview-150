# 添加与搜索单词 - 数据结构设计

题目链接: [https://leetcode.cn/problems/design-add-and-search-words-data-structure](https://leetcode.cn/problems/design-add-and-search-words-data-structure)

## 解题思路：

1. 按照每个字符串内字符出现的先后顺序构建字符串树
2. `search`函数遍历待查找字符串内的所有字符，遍历是否存在一条路线能构成待查找字符串
3. 若某个字符为`.`则需要遍历所有存在的路线
4. 若能则返回末端节点，否则返回`nil`

```go

type WordDictionary struct {
	subChar []*WordDictionary
	end     bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		subChar: make([]*WordDictionary, 26),
	}
}

func search(node *WordDictionary, word string) *WordDictionary {
	if len(word) == 0 {
		if node.end {
			return node
		}
		return nil
	}
	char := word[0]
	if char == '.' {
		for _, item := range node.subChar {
			if item != nil {
				res := search(item, word[1:])
				if res != nil && res.end {
					return res
				}
			}
		}
		return nil
	} else {
		char = char - 'a'
		if node.subChar[char] == nil {
			return nil
		}
		return search(node.subChar[char], word[1:])
	}
}

func (this *WordDictionary) AddWord(word string) {
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

func (this *WordDictionary) Search(word string) bool {
	node := search(this, word)
	if node == nil || !node.end {
		return false
	}
	return true
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n)$$,$$n$$为字符串长度
- **空间复杂度：** 空间复杂度为$$O(n)$$,$$n$$为所有字符串的字符顺序数
