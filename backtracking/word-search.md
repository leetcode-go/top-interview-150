# 单词搜索

题目链接: [https://leetcode.cn/problems/word-search](https://leetcode.cn/problems/word-search)

## 解题思路：

1. 遍历`board`中的每个元素，从这个元素出发是否能找到一条路线覆盖整个`word`
2. 若中间出现某个元素与对应的字符不匹配，则返回上层进行下一个方向的遍历
3. 每个方向的遍历过程中将遍历过的元素置为`-`，返回是恢复原值

```go
var directs = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

func exist(board [][]byte, word string) bool {
	m, n, wordLen := len(board), len(board[0]), len(word)
	if m*n < len(word) {
		return false
	}

	var backTrack func(i, j int, wordIndex int) bool

	backTrack = func(i, j int, wordIndex int) bool {
		if wordIndex == wordLen {
			return true
		}
		if i >= m || j >= n || i < 0 || j < 0 {
			return false
		}
		tmp := board[i][j]
		if tmp == word[wordIndex] {
			board[i][j] = '-'
			res := false
			for _, direct := range directs {
				res = res || backTrack(i+direct[0], j+direct[1], wordIndex+1)
			}
			board[i][j] = tmp
			return res
		} else {
			return false
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backTrack(i, j, 0) {
				return true
			}
		}
	}
	return false
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n*m!)$$，$$n$$为`board`行数，$$m$$为`board`列数
- **空间复杂度：** 空间复杂度为$$O(n*m!)$$
