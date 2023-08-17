# 有效的括号

题目链接: [https://leetcode.cn/problems/valid-parentheses](https://leetcode.cn/problems/valid-parentheses)

## 解题思路：

1. 如果字符串长度为奇数，直接返回 `false`
2. 使用一个栈来存储左括号，遍历字符串 `s` 中的每个字符
3. 如果当前字符是左括号，则将其压入栈中，如果当前字符是右括号，则从栈中弹出一个左括号进行匹配
4. 如果栈为空，则说明没有左括号与当前右括号匹配，返回 `false`，如果栈顶的左括号与当前右括号不匹配，则返回 `false`
5. 如果遍历完成后栈不为空，则说明还有左括号没有匹配，返回 `false`。否则，返回 `true`


```go
func isValid(s string) bool {
	// 如果字符串长度为奇数，直接返回 false
	if len(s)&1 == 1 {
		return false
	}

	// 创建一个栈，用于存储左括号
	stack := make([]byte, len(s))
	top := 0

	// 遍历字符串中的每一个字符
	for i := 0; i < len(s); i++ {
		// 如果是左括号，将其入栈
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack[top] = s[i]
			top++
		} else {
			// 如果是右括号，判断栈是否为空
			if top == 0 {
				return false
			}
			// 如果栈不为空，判断栈顶元素是否与当前右括号匹配
			if s[i] == ')' && stack[top-1] != '(' {
				return false
			}
			if s[i] == '}' && stack[top-1] != '{' {
				return false
			}
			if s[i] == ']' && stack[top-1] != '[' {
				return false
			}
			// 如果匹配成功，将栈顶元素出栈
			top--
		}
	}

	// 如果栈为空，说明所有括号都匹配成功，返回 true
	return top == 0
}
```

## 复杂度分析

- **时间复杂度：** 只遍历了一遍字符串，因此时间复杂度为 $$O(n)$$，其中 $$n$$ 是字符串 `s` 的长度
- **空间复杂度：** 空间复杂度为 $$O(n)$$，函数使用了一个长度为 $$n$$ 的切片来存储栈。在最坏的情况下，栈的深度可能达到 $$n$$，所以空间复杂度为 $$O(n)$$
