# 组合

题目链接: [https://leetcode.cn/problems/combinations](https://leetcode.cn/problems/combinations)

## 解题思路：

1. 递归调用`backtrack`，每次从上一次使用的数字开始向后遍历，逐个尝试放入组合列表中
2. 若遍历到第`k`个数字，则将组合列表写入到结果列表中

```go
func combine(n, k int) [][]int {
	var res [][]int
	var backtrack func(int, int, []int)
	backtrack = func(last, idx int, tmp []int) {
		tmpList := make([]int, len(tmp))
		copy(tmpList, tmp)
		tmpList = append(tmpList, last)
		if idx == k {
			res = append(res, tmpList)
			return
		}
		for i := last + 1; i <= n; i++ {
			backtrack(i, idx+1, tmpList)
		}
	}
	for i := 1; i <= n; i++ {
		backtrack(i, 1, []int{})
	}
	return res
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n*(n-1)/2)$$
- **空间复杂度：** 空间复杂度为$$O(n*(n-1)/2)$$
