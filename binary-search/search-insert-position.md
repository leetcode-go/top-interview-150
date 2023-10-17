# 搜索插入位置

题目链接: [https://leetcode.cn/problems/search-insert-position](https://leetcode.cn/problems/search-insert-position)

## 解题思路：

1. 如果目标值大于数组中的最后一个元素，则插入到数组末尾
2. 如果目标值小于数组中的第一个元素，则插入到数组开头
3. 否则，使用二分查找法在数组中查找目标值
4. 如果找到目标值，则返回其在数组中的索引
5. 如果没有找到目标值，则返回其在数组中的插入位置

```go
func searchInsert(nums []int, target int) int {
	length := len(nums)

	//如果目标值大于数组中的最后一个元素，则插入到数组末尾
	if target > nums[length-1] {
		return length
	}
	//如果目标值小于数组中的第一个元素，则插入到数组开头
	if target < nums[0] {
		return 0
	}

	left, right := 0, length-1
	// 二分查找
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// 如果没有找到目标值，则返回插入位置
	return left
}
```

## 复杂度分析

- **时间复杂度：** 间复杂度是 $$O(\log n)$$，其中 $$n$$ 是数组的长度，因为每次查找都将数组的长度减半
- **空间复杂度：** 空间复杂度是 $$O(1)$$，因为只需要使用常数级别的额外空间来存储左右指针和中间位置的值
