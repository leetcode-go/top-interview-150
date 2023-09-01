# 串联所有单词的子串

题目链接: [https://leetcode.cn/problems/substring-with-concatenation-of-all-words/description](https://leetcode.cn/problems/substring-with-concatenation-of-all-words/description)

## 解题思路：

1. 以字符串`s`的每个字符为开头，截取`words`所有字符串长度之和的子串
2. 遍历`words`，构建哈希表
3. 遍历子串，将子串拆分楚跟`words`等量的字符串，判断字符串是否都出现在哈希表中
4. 若存在未在哈希表的字符串，则该子串非串联子串

```go
func findSubstring(s string, words []string) []int {
	res := make([]int, 0)
	length, m, n := len(s), len(words), len(words[0])
	totalLength := m * n
	for i := 0; i+totalLength <= length; i++ {
		wordsMap := map[string]int{}
		for _, item := range words {
			wordsMap[item] += 1
		}
		subStr := s[i : i+totalLength]
		match := true
		for j := 0; j < totalLength; j = j + n {
			if wordsMap[subStr[j:j+n]] == 0 {
				match = false
				break
			} else {
				wordsMap[subStr[j:j+n]] -= 1
			}
		}
		if match {
			res = append(res, i)
		}
	}
	return res

}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(length*n)$$,`length`为`s`的长度，`n`为`words`中每个单词的长度
- **空间复杂度：** 空间复杂度为$$O(n*m)$$，`m`为`words`的长度,`n`为`words`中每个单词的长度
