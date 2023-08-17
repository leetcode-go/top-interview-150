# 逆波兰表达式求职

题目链接: [https://leetcode.cn/problems/evaluate-reverse-polish-notation](https://leetcode.cn/problems/evaluate-reverse-polish-notation)

## 解题思路：

1. 遍历tokens，判断是否是算符，是则从栈顶出栈两个元素进行计算
2. 否则，转换成int，入栈


```go
func evalRPN(tokens []string) int {
	queue := make([]int, 0)
	for _, item := range tokens {
		last := len(queue) - 1
		switch item {
		case "+":
			val1, val2 := queue[last], queue[last-1]
			queue = append(queue[:last-1], val2+val1)
		case "-":
			val1, val2 := queue[last], queue[last-1]
			queue = append(queue[:last-1], val2-val1)
		case "*":
			val1, val2 := queue[last], queue[last-1]
			queue = append(queue[:last-1], val2*val1)
		case "/":
			val1, val2 := queue[last], queue[last-1]
			queue = append(queue[:last-1], val2/val1)
		default:
			data, _ := strconv.Atoi(item)
			queue = append(queue, data)
		}
	}
	return queue[0]
}
```

## 复杂度分析

- **时间复杂度：** 只遍历了一遍tokens，因此时间复杂度为 $$O(n)$$，其中 $$n$$ 是字符串数组 `tokens` 的长度
- **空间复杂度：** 只使用了一个栈，因此空间复杂度为 $$O(n)$$，
