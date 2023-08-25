# 汇总区间

题目链接: [https://leetcode.cn/problems/summary-ranges](https://leetcode.cn/problems/summary-ranges)

## 题目描述

给定一个 无重复元素 的 有序 整数数组 nums 。

返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表 。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个范围但不属于 nums 的数字 x 。

列表中的每个区间范围 \[a,b] 应该按如下格式输出：

* "a->b" ，如果 a != b
* "a" ，如果 a == b

示例 1：

> 输入：nums = \[0,1,2,4,5,7]
>
> 输出：\["0->2","4->5","7"]
>
> 解释：区间范围是：
>
> \[0,2] --> "0->2"
>
> \[4,5] --> "4->5"
>
> \[7,7] --> "7"&#x20;

示例 2：&#x20;

> 输入：nums = \[0,2,3,4,6,8,9]
>
> 输出：\["0","2->4","6","8->9"]
>
> 解释：区间范围是：
>
> \[0,0] --> "0"
>
> \[2,4] --> "2->4"
>
> \[6,6] --> "6"
>
> \[8,9] --> "8->9"

提示：

* 0 <= nums.length <= 20
* \-231 <= nums\[i] <= 231 - 1
* nums 中的所有值都 互不相同
* nums 按升序排列

## 解题方法：

1. 分别计算每个元素的前缀乘积以及后缀乘积，再将每个数的前缀乘积与后缀乘积相乘

```go
func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	if len(nums) == 1 {
		return []string{fmt.Sprintf("%d", nums[0])}
	}
	m := make(map[int]bool, len(nums))
	for _, v := range nums {
		m[v] = true
	}
	res := make([]string, 0)
	start := nums[0]
	needAppend := false
	for _, v := range nums {
		if !m[v-1] {
			start = v
		}
		if m[v+1] {
			needAppend = true
		} else {
			if needAppend {
				res = append(res, fmt.Sprintf("%d->%d", start, v))
			} else {
				res = append(res, fmt.Sprintf("%d", start))
			}
			needAppend = false
		}
	}
	return res
}
```

## 复杂度分析：

* **时间复杂度:** $$O(n)$$，$$n$$为$$nums$$中元素的个数
* **空间复杂度:** $$O(n)$$
