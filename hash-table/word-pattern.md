# 单词规律

题目链接: [https://leetcode.cn/problems/word-pattern](https://leetcode.cn/problems/word-pattern)

## 解题思路：

1. 遵循的前提是pattern中每个出现的字符都可以映射到s中的单词，且同一个字符不能映射到不同的单词身上
2. 因此需要先对s按空格拆分字符串数组得到sList
2. 再遍历字符串pattern，建立pattern每个字符字符与sList的每个单词的映射表以及sList每个单词与s的每个字符的映射表
3. 当出现sList中的某个单词被pattern中的两个字符映射或pattern中的某个字符映射到了sList中的两个单词时，则判定为不遵循

```go
func wordPattern(pattern string, s string) bool {
	sList := strings.Split(s, " ")
	if len(pattern)!=len(sList){
		return false
	}
	p2s := map[byte]string{}
	s2p := map[string]byte{}
	for i := range pattern {
		x, y := pattern[i], sList[i]
		if (p2s[x] != "" && p2s[x] != y) || (s2p[y] > 0 && s2p[y] != x) {
			return false
		}
		p2s[x] = y
		s2p[y] = x
	}
	return true
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度是 $$O(n)$$
- **空间复杂度：** 空间复杂度是 $$O(n)$$
