# 长度最小的子数组

题目链接: [https://leetcode.cn/problems/minimum-size-subarray-sum](https://leetcode.cn/problems/minimum-size-subarray-sum)

## 题目描述

给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

示例 1：

> 输入：target = 7, nums = [2,3,1,2,4,3]
>
> 输出：2
>
> 解释：子数组 [4,3] 是该条件下的长度最小的子数组。&#x20;

示例 2：&#x20;

> 输入：target = 4, nums = [1,4,4]
>
> 输出：1

示例 3：&#x20;

> 输入：target = 11, nums = [1,1,1,1,1,1,1,1]
>
> 输出：0

提示：

* 1 <= target <= 109
* 1 <= nums.length <= 105
* 1 <= nums[i] <= 105

## 解题方法：

1. 遍历数组，记录子数组的start以及end，若子数组的合大于等于target，则从start进行收缩数组，直至小于target
2. 同时若每次收缩后的长度小于记录的子数组的长度，则更新最小子数组长度

```go
func minSubArrayLen(target int, nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	res := math.MaxInt32
	start, end, sum := 0, 0, 0
	for end < length {
		sum += nums[end]
		for sum >= target {
			if res > end-start+1 {
				res = end - start + 1
			}
			sum -= nums[start]
			start++
		}
		end++
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}
```

## 复杂度分析：

* **时间复杂度:** $$O(n^2)$$，$$n$$为$$nums$$中元素的个数
* **空间复杂度:** $$O(1)$$
