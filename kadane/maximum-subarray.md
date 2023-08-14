# 最大子数组和

题目链接: [https://leetcode.cn/problems/maximum-subarray](https://leetcode.cn/problems/maximum-subarray)

## 解题思路：

> **Kadane 算法**是一种用于求解最大子序和问题的算法，它的时间复杂度为 $$O(n)$$，其中 $$n$$ 是数组的长度。该算法的基本思想是维护两个变量：当前的最大和和当前的和。遍历数组时，如果当前的和小于 $$0$$，则将当前的和更新为当前元素的值，否则将当前的和加上当前元素的值。然后判断当前的和是否大于当前的最大和，如果是，则将当前的最大和更新为当前的和。最后返回当前的最大和即可

1. 初始化 `max` 和 `sum` 为数组的第一个元素，
2. 然后从数组的第二个元素开始遍历，如果当前的和 `sum` 小于 `0`，则将 `sum` 更新为当前元素的值，否则将 `sum` 加上当前元素的值
3. 然后判断 `sum` 是否大于 `max`，如果是，则将 max 更新为 `sum`。最后返回 `max`

```go
func maxSubArray(nums []int) int {
	max := nums[0]
	sum := 0

	for _, num := range nums {
		sum += num
		if sum > max {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}

	return max
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度是 $$O(n)$$
- **空间复杂度：** 空间复杂度是 $$O(1)$$
