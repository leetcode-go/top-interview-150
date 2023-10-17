# 寻找峰值

题目链接: [https://leetcode.cn/problems/find-peak-element](https://leetcode.cn/problems/find-peak-element)

## 解题思路：

1. 首先计算数组的长度 $$n$$，并初始化左指针 `left` 和右指针 `right` 分别指向数组的第一个元素和最后一个元素
2. 使用 for 循环进行二分查找，直到左指针和右指针相等
3. 在每次循环中，计算中间位置 `mid`，并比较中间元素和右侧元素的大小
4. 如果中间元素大于右侧元素，则峰值元素在左半边，将右指针更新为 `mid`
5. 否则，峰值元素在右半边，将左指针更新为 $$mid + 1$$。
6. 循环结束后，返回左指针的值，即峰值元素的索引


```go
func findPeakElement(nums []int) int {
	n := len(nums)

	left, right := 0, n-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[mid+1] { // 如果中间元素大于右侧元素，则峰值元素在左半边
			right = mid // mid 可能是峰值
		} else { // 如果中间元素小于右侧元素，则峰值元素在右半边
			left = mid + 1
		}
	}

	return left
}
```

## 复杂度分析

- **时间复杂度：** 间复杂度是 $$O(\log n)$$，其中 $$n$$ 是数组的长度，因为每次查找都将数组的长度减半
- **空间复杂度：** 空间复杂度是 $$O(1)$$，因为只需要使用常数级别的额外空间来存储变量
