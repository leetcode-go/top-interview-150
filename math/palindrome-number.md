# 回文数

题目链接: [https://leetcode.cn/problems/palindrome-number](https://leetcode.cn/problems/palindrome-number)

## 解题思路：

1. 将数字翻转一半
2. 如果 $$x$$ 等于反转的数字或者 $$x$$ 等于反转的数字除以`10`（去掉中间的数字），那么 $$x$$ 是回文数

```go
func isPalindrome(x int) bool {
	// 如果x是负数或者x的个位数是0但x不是0，那么x不是回文数
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	// 将x的数字反转，直到反转的数字大于等于x
	var reversed int
	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}

	// 如果x等于反转的数字或者x等于反转的数字除以10（去掉中间的数字），那么x是回文数
	return x == reversed || x == reversed/10
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度是 $$O(\log_{10} x)$$，其中 $$x$$ 是输入的整数
- **空间复杂度：** 空间复杂度是 $$O(1)$$，函数只使用了常数个变量来存储中间结果，没有使用额外的空间
