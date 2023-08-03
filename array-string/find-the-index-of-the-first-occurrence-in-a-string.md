# 找出字符串中第一个匹配项的下标

题目链接: [https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string](https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string)

## 解题思路：

1. 枚举原串 `haystack` 中的每个字符作为发起点，每次从原串的发起点和匹配串的首位开始尝试匹配：
	- 匹配成功：返回本次匹配的原串发起点
	- 匹配失败：枚举原串的下一个发起点，重新尝试匹配


```go
func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i] == needle[0] && haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度是 $$O(mn)$$，其中 $$m$$ 是字符串 `haystack` 的长度，$$n$$ 是字符串 `needle` 的长度。在函数中，需要遍历字符串 `haystack` 一次，并且对于每个字符，需要比较与子串 `needle` 是否相等
- **空间复杂度：** 只使用了常数个变量，因此空间复杂度为 $$O(1)$$
