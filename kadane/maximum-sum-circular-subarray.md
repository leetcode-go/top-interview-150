# 环形子数组的最大和

题目链接: [https://leetcode.cn/problems/maximum-sum-circular-subarray](https://leetcode.cn/problems/maximum-sum-circular-subarray)

## 解题思路：

将环形数组拆成两个普通数组，然后分别求出两个数组的最大子序和，最后取两个最大子序和的较大值作为环形子数组的最大和

**具体实现：**

1. 遍历数组，计算当前的最大和、最小和、最大子序和、最小子序和和数组的和
2. 如果数组的和等于最小子序和，则说明数组中所有元素都是负数，此时最大子序和即为最大值
3. 否则，最大子序和为 `max(maxSum, sum-minSum)`

```go
func maxSubarraySumCircular(nums []int) int {
	maxSum, minSum := nums[0], nums[0]
	curMax, curMin := nums[0], nums[0]
	sum := nums[0]

	// 遍历数组，计算当前的最大和、最小和、最大子序和、最小子序和和数组的和
	for i := 1; i < len(nums); i++ {
		sum += nums[i]                        // 计算数组的和
		curMax = max(curMax+nums[i], nums[i]) // 计算当前的最大和
		maxSum = max(maxSum, curMax)          // 计算最大子序和
		curMin = min(curMin+nums[i], nums[i]) // 计算当前的最小和
		minSum = min(minSum, curMin)          // 计算最小子序和
	}

	// 如果数组的和等于最小子序和，则说明数组中所有元素都是负数，此时最大子序和即为最大值
	if sum == minSum {
		return maxSum
	}

	// 否则，最大子序和为 max(maxSum, sum-minSum)
	return max(maxSum, sum-minSum)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度是 $$O(n)$$
- **空间复杂度：** 空间复杂度是 $$O(1)$$
