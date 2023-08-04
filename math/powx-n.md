# Pow(x, n)

题目链接: [https://leetcode.cn/problems/powx-n](https://leetcode.cn/problems/powx-n)

## 解题思路：

1. 判断 $$n$$ 是否等于 $$0$$，等于 $$0$$ 直接返回 $$1$$
2. 如果 $$n$$ 小于 $$0$$，将 $$x$$ 取倒数，$$n$$ 取相反数
3. 如果 $$n$$ 是偶数，递归调用 `myPow` 函数，传入 $$x$$ 的平方和 $$n$$ 除以 $$2$$ 的结果
4. 如果 $$n$$ 是奇数，返回 $$x$$ 乘以递归调用 `myPow` 函数，传入 $$x$$ 的平方和 $$n$$ 除以 $$2$$ 的结果

```go
func myPow(x float64, n int) float64 {
	// 如果 n 等于 0，返回 1
	if n == 0 {
		return 1
	}

	// 如果 n 小于 0，将 x 取倒数，n 取相反数
	if n < 0 {
		x = 1 / x
		n = -n
	}

	// 如果 n 是偶数，递归调用 myPow 函数，传入 x 的平方和 n 除以 2 的结果
	if n%2 == 0 {
		return myPow(x*x, n/2)
	}

	// 如果 n 是奇数，返回 x 乘以递归调用 myPow 函数，传入 x 的平方和 n 除以 2 的结果
	return x * myPow(x*x, n/2)
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度是 $$O(\log_{10} n)$$
- **空间复杂度：** 空间复杂度是 $$O(\log_{10} n)$$
