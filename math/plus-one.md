# 加一

题目链接: [https://leetcode.cn/problems/plus-one](https://leetcode.cn/problems/plus-one)

## 解题思路：

1. 从最后一个元素开始，依次向前遍历，对于每个元素，如果其小于 `9`，则将其加一并返回 `digits` 数组
2. 否则，将其置为 `0`，继续向前遍历。如果遍历完整个数组后仍未返回，则说明需要在数组的最前面添加一个元素 `1`，表示进位

```go
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	return append([]int{1}, digits...)
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度是 $$O(n)$$，其中 $$n$$ 是数组 `digits` 的长度。在函数中，需要遍历数组 `digits`，对于每个元素，需要执行一次比较和一次赋值操作
- **空间复杂度：** 空间复杂度是 $$O(1)$$，函数只使用了常数个变量来存储中间结果，没有使用额外的空间
