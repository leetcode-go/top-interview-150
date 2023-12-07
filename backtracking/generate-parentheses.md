# 括号生成

题目链接: [https://leetcode.cn/problems/generate-parentheses](https://leetcode.cn/problems/generate-parentheses)

## 解题思路：

1. 要生成`n`对括号，那么就需要`2n`个位置，每个位置都可以放置一个左括号或者右括号，那么可以分为两种情况：
    - 放置左括号，那么接下来还可以放置左括号或者右括号
    - 放置右括号，那么接下来只能放置右括号
2. 通过`i`及`open`记录一共写入了多少位，以及有多少个右括号
3. 当`i == 2n`时，说明已经写完了，将结果添加到结果集
4. 当`open < n`时，说明还可以继续放置左括号，那么就放置一个左括号，然后递归调用，
5. 当`i - open < open`时，说明此时可以放置右括号了，那么就放置一个右括号，然后递归调用

```go
func generateParenthesis(n int) []string {
	var res []string
	m := n * 2
	path := make([]byte, m)
	var backtrack func(i, open int)
	backtrack = func(i, open int) {
		if i == m {
			res = append(res, string(path))
		}
		if open < n {
			path[i] = '('
			backtrack(i+1, open+1)
		}
		if i-open < open {
			path[i] = ')'
			backtrack(i+1, open)
		}
	}
	backtrack(0, 0)
	return res
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(2^n)$$
- **空间复杂度：** 空间复杂度为$$O(n)$$
