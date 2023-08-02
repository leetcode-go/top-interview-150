# 三数之和

题目链接: [https://leetcode.cn/problems/3sum/](https://leetcode.cn/problems/3sum/)

## 解题思路：

排序 + 双指针

1. 判断数组 `nums` 的长度是否小于 `3`，如果是，则直接返回 `nil`
2. 对数组 `nums` 进行排序
3. 遍历数组 `nums` 中的每个数字，在循环中，函数首先判断当前数字是否大于 `0`，如果是，则说明三数之和一定大于 `0`，直接结束循环
4. 使用双指针来查找数组中所有和为当前数字的相反数的不重复的三元组，在双指针遍历的过程中，判断当前三个数字的和与 `0` 的大小关系，并根据不同的情况来移动指针
5. 去重：首先，在遍历数组 `nums` 的过程中，如果当前数字与前一个数字相同，则直接跳过，避免重复计算。其次，在找到一个和为 0 的三元组后，函数会继续移动指针，直到找到下一个不同的数字，避免重复计算

```go
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	// 排序
	quickSort(nums, 0, len(nums)-1)

	ans := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		// 如果当前数字大于0，则三数之和一定大于0，所以结束循环
		if nums[i] > 0 {
			break
		}
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 双指针
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			switch {
			case sum == 0:
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				// 去重
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// 去重
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			case sum < 0:
				left++
			case sum > 0:
				right--
			}
		}
	}

	return ans
}

func quickSort(nums []int, left, right int) {
	if left < right {
		pivot := partition(nums, left, right)
		quickSort(nums, left, pivot-1)
		quickSort(nums, pivot+1, right)
	}
}

func partition(nums []int, left, right int) int {
	pivot := nums[left]
	for left < right {
		for left < right && nums[right] >= pivot {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] <= pivot {
			left++
		}
		nums[right] = nums[left]
	}

	nums[left] = pivot

	return left
}
```

## 复杂度分析

- **时间复杂度：**时间复杂度为 $$O(n^2)$$，其中 $$n$$ 是数组 `nums` 的长度。函数中有两个循环，外层循环遍历了数组 `nums` 中的每个数字，内层循环使用双指针遍历数组中的剩余数字。在内层循环中，左指针和右指针分别从数组的两端开始向中间移动，因此内层循环的时间复杂度为 $$O(n)$$。因此，函数的总时间复杂度为 $$O(n^2)$$
- **空间复杂度：** 只使用了常数个变量，因此空间复杂度为 $$O(1)$$，函数只使用了常数个变量来存储中间结果，没有使用额外的空间
