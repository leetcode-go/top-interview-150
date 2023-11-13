# 单词搜索 II

题目链接: [https://leetcode.cn/problems/word-search-ii](https://leetcode.cn/problems/word-search-ii)

## 解题思路：

1. 对于所有的`words`构建一个前缀树，前缀树的每个节点对应一个单词，每个单词的最后一个字母对应一个节点，单词的其余部分对应子节点
2. 遍历`board`，判断每个位置出发的四个方向是否能走到前缀树的叶子节点，若能走到，则覆盖该单词，记录已遍历
3. 对于从一个位置出发后，则将该位置标记为'#'，避免重复遍历，四个方向遍历完之后，则将该位置复原

```go
type Trie struct {
	subTrie []*Trie
	word    string
}

func contract() *Trie {
	return &Trie{
		subTrie: make([]*Trie, 26),
	}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.subTrie[ch] == nil {
			node.subTrie[ch] = contract()
		}
		node = node.subTrie[ch]
	}
	node.word = word
}

func findWords(board [][]byte, words []string) []string {
	if words == nil || len(words) == 0 || board == nil || len(board) == 0 {
		return nil
	}
	result := make([]string, 0)
	root := contract()
	directs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, word := range words {
		root.Insert(word)
	}
	m, n := len(board), len(board[0])
	seen := map[string]bool{}
	var dfs func(x int, y int, t *Trie)
	dfs = func(x int, y int, trie *Trie) {
		ch := board[x][y]
		trie = trie.subTrie[ch - 'a']
		if trie == nil {
			return
		}
		if trie.word != "" {
			seen[trie.word] = true
		}
		board[x][y] = '#'
		for _, direct := range directs {
			tmpx, tmpy := x+direct[0], y+direct[1]
			if tmpx >= 0 && tmpx < m && tmpy >= 0 && tmpy < n && board[tmpx][tmpy] != '#' {
				dfs(tmpx, tmpy, trie)
			}
		}
		board[x][y] = ch
	}
	for i, row := range board {
		for j := range row {
			dfs(i, j, root)
		}
	}

	for key, _ := range seen {
		result = append(result, key)
	}
	return result
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(m*n)$$
- **空间复杂度：** 空间复杂度为$$O(m*n)$$
