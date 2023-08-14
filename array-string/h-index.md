# H指数


题目链接: [https://leetcode.cn/problems/h-index](https://leetcode.cn/problems/h-index)

## 解题思路：

1. h指数的条件为`引用数大于等于某个值，且满足引用数大于等于某个值的文章数量也大于等于这个值`
2. 所以我们只需要对数组进行排序，对排序后的数组进行遍历
3. 判断每个引用数到数组末尾一共多少个引用数，这些引用数一定是大于等于当前这个引用数的，找出最大的那个值就可以了


```go
func hIndex(citations []int) int {
	sort.Ints(citations)
	max := 0
	length := len(citations)
	for idx, item := range citations {
		count := length - idx
		if item >= count && count > max {
			max = count
		}
	}
	return max
}
```

## 复杂度分析一

- **时间复杂度：** 对数组进行了一次遍历，因此时间复杂度为 $$O(n)$$，其中 $$n$$ 是数组 $$citations$$ 的长度
- **空间复杂度：** 空间复杂度为 $$O(1)$$