# 两数之和

题目链接: [https://leetcode.cn/problems/two-sum](https://leetcode.cn/problems/two-sum)

## 解题思路：

1. 变量数组，对每一个元素，在 `map` 中找能组合给定值的另一半数字，如果找到了，直接返回两个元素的下标即可
2. 如果找不到，就把这个数字存入 map 中，等待可以组合的数字的时候，再取出来返回结果
3. 如果数组遍历结束，返回 `nil`

```go
func twoSum(nums []int, target int) []int {
	// 创建一个 map，用于存储数组元素的值和索引
	m := make(map[int]int)

	// 遍历数组 nums
	for i, num := range nums {
		// 检查 map 中是否存在满足条件的目标元素
		if j, ok := m[target-num]; ok {
			// 如果存在，直接返回结果
			return []int{j, i}
		}

		// 如果不存在，将当前元素的值和索引存入 map
		m[num] = i
	}

	// 如果遍历结束，仍然没有找到满足条件的目标元素，则返回一个空的数组
	return []int{}
}
```

## 复杂度分析

- 时间复杂度：$$O(n)$$，其中 $$n$$ 是数组中的元素数量
- 空间复杂度：$$O(n)$$，其中 $$n$$ 是数组中的元素数量，`map` 开销
