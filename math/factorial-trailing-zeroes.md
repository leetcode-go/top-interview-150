# 阶乘后的零

题目链接: [https://leetcode.cn/problems/factorial-trailing-zeroes](https://leetcode.cn/problems/factorial-trailing-zeroes)

## 解题思路：

1. 首先判断 $$n$$ 是否小于 $$5$$，如果是，则返回 $$0$$，因为小于 $$5$$ 的数的阶乘后不会有零
2. 否则，将 $$n$$ 除以 $$5$$，得到商和余数。商表示 $$n$$ 中包含 $$5$$ 的个数，余数表示 $$n$$ 中包含 $$25$$ 的个数，以此类推
3. 然后，将商和余数分别传递给递归函数 `trailingZeroes`，计算商和余数的阶乘后的零的个数
4. 最后，将商和余数的阶乘后的零的个数相加，得到 $$n$$ 的阶乘后的零的个数

```go
func trailingZeroes(n int) int {
	if n < 5 {
		return 0
	}

	return n/5 + trailingZeroes(n/5)
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度是 $$O(\log_{10} n)$$，其中 $$n$$ 是输入的整数
- **空间复杂度：** 空间复杂度是 $$O(1)$$
