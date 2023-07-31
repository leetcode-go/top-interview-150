# 验证回文串

题目链接: [https://leetcode.cn/problems/valid-palindrome](https://leetcode.cn/problems/valid-palindrome)

## 解题思路：

**回文串**

1. 将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样
2. 字母和数字都属于字母数字字符

**解题思路**

1. 左右双指针，跳过非字母数字的字符
2. 大小写字母 ASCII 值差距是 32
3. 详细的思路看代码块中的注释

```go
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	// 左右指针相向移动，直到相遇
	for left < right {
		// 如果左指针指向的字符不是字母或数字，则左指针右移
		if !isAlphanumeric(s[left]) {
			left++
			continue
		}
		// 如果右指针指向的字符不是字母或数字，则右指针左移
		if !isAlphanumeric(s[right]) {
			right--
			continue
		}
		// 如果左右指针指向的字符不相等，则不是回文字符串
		if s[left]|32 != s[right]|32 {
			return false
		}
		// 左右指针同时向中间移动
		left++
		right--
	}

	return true
}

// 判断一个字符是否为字母或数字
func isAlphanumeric(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || '0' <= c && c <= '9'
}
```

## 复杂度分析

- **时间复杂度：**只遍历了一遍字符串，因此时间复杂度为 $O(n)$，其中 $n$ 是字符串的长度
- **空间复杂度：**只使用了常数个变量，因此空间复杂度为 $O(1)$
