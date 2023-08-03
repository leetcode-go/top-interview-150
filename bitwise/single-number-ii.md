# 只出现一次的数字 II

题目链接: [https://leetcode.cn/problems/single-number-ii](https://leetcode.cn/problems/single-number-ii)

## 解题思路：

1. 对于每个数字 `num`，分别更新变量 `a` 和 `b` 的值，具体来说，更新 `a` 的值为 `(a ^ num) & ^b`，更新 `b` 的值为 `(b ^ num) & ^a`
2. 这两个式子的含义是将 `num` 与 `a` 进行异或运算，然后将结果与 `b` 取反的结果进行按位与运算，得到新的 `a` 的值
3. 将 `num` 与 `b` 进行异或运算，然后将结果与 `a` 取反的结果进行按位与运算，得到新的 `b` 的值
4. 这样可以保证 `a` 和 `b` 中只存储出现一次的数字，而不存储出现多次的数字

```go
func singleNumber(nums []int) int {
	a, b := 0, 0

	for _, num := range nums {
		a = (a ^ num) & ^b
		b = (b ^ num) & ^a
	}

	return a
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度：$$O(n)$$，其中 $$n$$ 是数组长度
- **空间复杂度：** 空间复杂度是 $$O(1)$$
