# 只出现一次的数字

题目链接: [https://leetcode.cn/problems/single-number](https://leetcode.cn/problems/single-number)

## 解题思路：

1. 异或运算：同性无结果

```go
func singleNumber(nums []int) int {
	var ans int
	for _, num := range nums {
		ans ^= num
	}
	return ans
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度：$$O(n)$$，其中 $$n$$ 是数组长度，只需要对数组遍历一次
- **空间复杂度：** 空间复杂度是 $$O(1)$$

