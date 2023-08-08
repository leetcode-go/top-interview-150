# x 的平方根

题目链接: [https://leetcode.cn/problems/sqrtx](https://leetcode.cn/problems/sqrtx)

## 解题思路：

**牛顿迭代法**

1. 首先判断输入的整数是否小于 $$2$$，如果是，则返回该整数
2. 否则，使用牛顿迭代法来计算该整数的平方根。在牛顿迭代法中，使用变量 `guess` 表示当前的猜测值
3. 在每次迭代中，计算 `guess` 的平方与该整数的差值，如果差值小于等于一个给定的精度，则返回 `guess`
4. 否则，将 `guess` 更新为 `(guess + x/guess) / 2`，然后继续迭代
5. 最终，当差值小于等于给定的精度时，返回 `guess` 即可

```go
func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	guess := x / 2        //初始化猜测值为 x 的一半
	for guess*guess > x { // 当猜测值的平方大于 x 时，执行循环
		guess = (guess + x/guess) / 2 // 更新猜测值为 (猜测值加上 x 除以猜测值的商)除以 2
	}

	return guess
}
```

**二分查找**

1. 首先判断输入的整数是否小于 $$2$$，如果是，则返回该整数
2. 否则，使用二分查找算法来计算该整数的平方根。在二分查找中，使用变量 `left` 和 `right` 分别表示搜索区间的左右端点
3. 在每次循环中，计算搜索区间的中间点 `mid`，如果中间点的平方等于该整数，则返回中间点。如果中间点的平方小于该整数，则将搜索区间的左端点移动到中间点的右侧，否则将搜索区间的右端点移动到中间点的左侧。最终，当搜索区间的左端点大于右端点时，返回右端点即可

```go
func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	left, right := 1, x/2 // 初始化搜索区间的左右端点
	for left <= right {   // 当搜索区间的左端点小于等于右端点时，执行循环
		mid := left + (right-left)/2 // 计算搜索区间的中间点
		if mid == x/mid {            // 如果中间点的平方等于 x，返回中间点
			return mid
		} else if mid < x/mid { // 如果中间点的平方小于 x，将搜索区间的左端点移动到中间点的右侧，继续搜索
			left = mid + 1
		} else {
			right = mid - 1 // 如果中间点的平方大于 x，将搜索区间的右端点移动到中间点的左侧，继续搜索
		}
	}

	return right
}
```

## 复杂度分析

> 两种思路复杂度一致

- **时间复杂度：** 时间复杂度为 $$O(\log x)$$
- **空间复杂度：** 空间复杂度为 $$O(1)$$