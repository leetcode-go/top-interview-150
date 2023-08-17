# 基本计算器

题目链接: [https://leetcode.cn/problems/basic-calculator](https://leetcode.cn/problems/basic-calculator)

## 解题思路：

1. 遍历一遍字符串，因为只有加减操作，则用sign表示当前是什么计算
2. 因为存在括号，则括号内的操作计算要根据括号外的sign值来判断是否翻转操作，即括号外为-号，括号内所有操作都应该翻转

```go
func calculate(s string) int {
	queue := []int{1}
	sign, ans := 1, 0
	n := len(s)
	for i := 0; i < n; {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = queue[len(queue)-1]
			i++
		case '-':
			sign = -queue[len(queue)-1]
			i++
		case '(':
			queue = append(queue, sign)
			i++
		case ')':
			queue = queue[:len(queue)-1]
			i++
		default:
			num := 0
			for ; i < n && '0' <= s[i] && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			ans += sign * num
		}
	}
	return ans
}
```

## 复杂度分析

- **时间复杂度：** 只遍历了一遍字符串，因此时间复杂度为 $$O(n)$$，其中 $$n$$ 是字符串 `s` 的长度
- **空间复杂度：** 空间复杂度为 $$O(n)$$，函数使用了一个长度为 $$n$$ 的切片来存储栈。
