# 跳跃游戏II


题目链接: [https://leetcode.cn/problems/jump-game-ii](https://leetcode.cn/problems/jump-game-ii)

## 解题思路一：

1. 反向遍历数组，从最后一个元素出发，然后从前往后查找哪个元素能跳到这个元素，一但查到，记录该元素的下标，跳跃次数+1，跳出本次查询
2. 重复步骤1，不同的是每次从前往后找那个元素能跳到上一轮查找到的那个元素的下标，直至查找到第0位元素为止


```go
func jump(nums []int) bool {
    step := 0
	idx := len(nums) - 1
	for idx > 0 {
		for i := 0; i < idx; i++ {
			if i+nums[i] >= idx {
				idx = i
				step++
				break
			}
		}
	}
	return step
}
```

## 复杂度分析一

- **时间复杂度：** 对数组进行了双重遍历，因此时间复杂度为 $$O(n^2)$$，其中 $$n$$ 是数组 $$nums$$ 的长度
- **空间复杂度：** 空间复杂度为 $$O(1)$$

## 解题思路二：

1. 正向遍历数组，计算查找出每次能跳的最远的距离，直至覆盖到最后一个元素，统计跳跃次数


```go
func jump(nums []int) int {
    step, end, max, length := 0, 0, 0, len(nums)-1
	for i := 0; i < length; i++ {
		if max < i+nums[i] {
			max = i + nums[i]
		}
        // 判断出是否已经遍历完当前这个end到对应的i之间的所有元素能达到的最远距离，开始下一轮最远跳跃距离的查询
		if i == end {
			end = max
			step++
		}
	}
	return step
}
```

## 复杂度分析一

- **时间复杂度：** 对数组进行了一次遍历，因此时间复杂度为 $$O(n)$$，其中 $$n$$ 是数组 $$nums$$ 的长度
- **空间复杂度：** 空间复杂度为 $$O(1)$$