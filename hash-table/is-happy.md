# 快乐数

题目链接: [https://leetcode.cn/problems/happy-number](https://leetcode.cn/problems/happy-number)

## 解题思路：

1. 对于任何一个大于 $$4$$ 的数字，最终都会得到 $$1$$ 或 $$4$$，因此可以在循环中判断当前数字是否为 $$4$$，如果是，则直接返回 `false`
2. 当 $$n$ 不等于 $$1$$ 或 $$4$$ 时，反复计算各个位上数字的平方和

```go
func isHappy(n int) bool {
	for n != 1 { // 当n不等于1时，继续循环
		if n == 4 { // 如果n等于4，返回false
			return false
		}
		n = next(n) // 计算n的下一个数
	}

	return n == 1 // 如果n最终等于1，返回true
}

// next 返回 n 的各个位上数字的平方和
func next(n int) int {
	sum := 0
	for n > 0 { // 当n大于0时，继续循环
		digit := n % 10      // 取n的个位数字
		sum += digit * digit // 将个位数字的平方加到sum上
		n /= 10              // 将n的个位数字去掉
	}
	return sum // 返回各个位上数字的平方和
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度是 $$O(k)$$，其中 $$k$$ 是最终得到 $$1$$ 或者 $$4$$ 的迭代次数。在最坏的情况下，迭代次数是一个循环链，其中每个数字都会被访问一次，因此迭代次数是有限的
- **空间复杂度：** 空间复杂度是 $$O(1)$$，函数只使用了常数个变量来存储中间结果，没有使用额外的空间

`map` 解法：

```go
func isHappy(n int) bool {
	m := make(map[int]bool)
	
	for n != 1 {
		if m[n] {
			return false
		}
		m[n] = true
		n = next(n)
	}

	return true
}
```
