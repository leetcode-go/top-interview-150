# 盛最多水的容器

题目链接: [https://leetcode.cn/problems/container-with-most-water](https://leetcode.cn/problems/container-with-most-water)

## 解题思路：

1. 当左指针小于右指针时，计算当前容器的面积，并更新最大值
2. 如果左指针对应的高度小于右指针对应的高度，则左指针向右移动一位，否则右指针向左移动一位
3. 重复上面两个步骤，直到左指针和右指针相遇为止

```go
func maxArea(height []int) int {
	max := 0                        // 用于记录最大值
	left, right := 0, len(height)-1 // 初始化左指针和右指针

	for left < right { // 当左指针小于右指针时，循环执行以下操作
		area := (right - left) * min(height[left], height[right]) // 计算当前容器的面积
		if area > max {                                           // 如果当前面积大于最大值，则更新最大值
			max = area
		}

		if height[left] > height[right] { // 如果左指针对应的高度大于右指针对应的高度，则右指针向左移动一位
			right--
		} else { // 否则左指针向右移动一位
			left++
		}
	}

	return max
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
```

## 复杂度分析

- **时间复杂度：** 只遍历了一遍字符串，因此时间复杂度为 $$O(n)$$，其中 $$n$$ 是数组的长度
- **空间复杂度：** 只使用了常数个变量，因此空间复杂度为 $$O(1)$$
