# 全排列

题目链接: [https://leetcode.cn/problems/combination-sum](https://leetcode.cn/problems/combination-sum)

## 解题思路：

1. 从0开始将`candidates`里的数据写入`comb`，同时将`target`减去相应的数值
2. 判断`target`是否为0，如果为0，说明找到了一组解，将`comb`写入结果集
3. 到了`idx == len(candidates)`时，说明已经遍历完所有候选数，返回

```go
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var comb []int
	var backTracking func(idx, tmpTarget int)
	backTracking = func(idx, tmpTarget int) {
		if idx == len(candidates) {
			return
		}
		if tmpTarget == 0 {
			res = append(res, append([]int{}, comb...))
			return
		}
		backTracking(idx+1, tmpTarget)
		if tmpTarget-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			backTracking(idx, tmpTarget-candidates[idx])
			comb = comb[:len(comb)-1]
		}
	}
	backTracking(0, target)
	return res
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n)$$，$$n$$为所有答案解的个数
- **空间复杂度：** 空间复杂度为$$O(n)$$，$$n$$为答案数组长度
