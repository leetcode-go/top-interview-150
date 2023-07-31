# 判断子序列

题目链接: [https://leetcode.cn/problems/is-subsequence](https://leetcode.cn/problems/is-subsequence)

## 解题思路：

1. 双指针遍历
2. 如果 $$s$$ 是一个空字符串，它被认为是任何字符串的子序列
3. 如果 $$s$$ 的当前字符与 $$t$$ 的当前字符匹配，$$i$$ 将加 $$1$$
4. 如果 $$i$$ 到达 $$s$$ 的末尾，$$s$$ 是 $$t$$ 的子序列，函数返回`true`
5. 如果 $$j$$ 到达了 $$t$$ 的末尾，而我没有到达 $$s$$ 的末尾，那么 $$s$$ 不是 $$t$$ 的子序列，函数返回`false`

```go
func isSubsequence(s, t string) bool {
	// 如果 s 为空字符串，则它是任何字符串的子序列，返回 true
	if len(s) == 0 {
		return true
	}

	// 用两个指针 i 和 j 分别指向字符串 s 和 t 的开头
	i := 0
	for j := 0; j < len(t); j++ {
		// 如果当前字符与 s 中的字符相同，则将 i 向后移动一位
		if s[i] == t[j] {
			i++
			// 如果 i 移动到了 s 的末尾，则说明 s 是 t 的子序列，返回 true
			if i == len(s) {
				return true
			}
		}
	}

	// 如果遍历完 t 后 i 还没有移动到 s 的末尾，则说明 s 不是 t 的子序列，返回 false
	return false
}
```

## 复杂度分析

- **时间复杂度：** 只遍历了一遍字符串 $$t$$，因此时间复杂度为 $$O(n)$$，其中 $$n$$ 是字符串 $$t$$ 的长度
- **空间复杂度：** 只使用了常数个变量，因此空间复杂度为 $$O(1)$$，无论输入的字符串 $$s$$ 和 $$t$$ 的长度如何，函数使用的空间都是固定的
