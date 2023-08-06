# 最长连续序列

题目链接: [https://leetcode.cn/problems/longest-consecutive-sequence](https://leetcode.cn/problems/longest-consecutive-sequence)

## 解题思路：

1. 对数组进行排序，得到左到右从小到大的有序数组
2. 从头开始遍历判断每个元素是否比前一个元素大1，若是判定连续
3. 若两个元素相等，则跳过
4. 若当前元素比前一个元素大超过1，则判定断开
5. 检查前一段连续数据的长度是否比记录的最大连续长度大，若是保存，否则重置tmp计数
6. 结束前再判断当前连续段的长度是否比记录的最大连续长度大，若是则覆盖
7. 返回最大长度

```go
func twoSum(nums []int, target int) []int {
	func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	max, tmp := 0, 1
	last := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]-last == 1 {
			tmp += 1
		} else if nums[i] == last {
			continue
		} else {
			if tmp >= max {
				max = tmp
			}
			tmp = 1
		}
		last = nums[i]
	}
	if tmp > max {
		max = tmp
	}
	return max
}
}
```

## 复杂度分析

* 时间复杂度：$$O(n)$$，其中 $$n$$ 是数组中的元素数量
* 空间复杂度：$$O(1)$$
