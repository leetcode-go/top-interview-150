# 数字范围按位与

题目链接: [https://leetcode.cn/problems/bitwise-and-of-numbers-range](https://leetcode.cn/problems/bitwise-and-of-numbers-range)

## 解题思路：

对于每个数字，将其与 `right` 进行按位与运算，将结果存储到 `right` 中。然后，将 `right` 减 `1`，再次进行循环。这样可以保证 `right` 中存储的是给定范围内所有数字的按位与

```go
func rangeBitwiseAnd(left int, right int) int {
	for left < right {
		right &= (right - 1)
	}

	return right
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度是 $$O(\log n)$$，其中 $$n$$ 是 `left` 和 `right` 的二进制表示中的位数
- **空间复杂度：** 空间复杂度是 $$O(1)$$
