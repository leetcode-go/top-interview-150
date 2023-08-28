# 反转字符串重的单词

题目链接: [https://leetcode.cn/problems/reverse-words-in-a-string](https://leetcode.cn/problems/reverse-words-in-a-string)

# 解题思路

1. 从字符串s末尾开始遍历，遇到非空格就追加到tmp前面，遇到空格就将tmp拼接到res后面

```go
func reverseWords(s string) string {
	res,tmp := "",""
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			tmp = string(s[i]) + tmp
		} else if tmp != "" {
			res = getStr(res, tmp)
			tmp = ""
		}
	}
	if tmp != "" {
		res = getStr(res, tmp)
	}
	return res
}

func getStr(res, tmp string) string {
	if len(res) != 0 {
		res += " "
	}
	return res + tmp
}

```

## 复杂度分析

- **时间复杂度:** $$O(n)$$
- **空间复杂度:** $$O(n)$$
