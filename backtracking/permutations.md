# 全排列

题目链接: [https://leetcode.cn/problems/permutations](https://leetcode.cn/problems/permutations)

## 解题思路：

1. 遍历`nums`数组，将每个数字写入组合队列中，并在`used`中标记该数字在本次组合中已经被使用
2. 递归调用，尝试写入下一个数据
3. 递归返回时，将当前数字从组合队列中移除，并从`used`中移除该数字的使用标记
4. 递归终止条件：当组合队列中数字个数等于`nums`数组长度时，说明已经生成本次组合，记录到`res`数组中

```go
func permute(nums []int) [][]int {
	var res [][]int
	var backtrack func([]int, []bool)
	backtrack = func(path []int, used []bool) {
		if len(path) == len(nums) {
			res = append(res, path)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			backtrack(append(path, nums[i]), used)
			used[i] = false
		}
	}
	backtrack([]int{}, make([]bool, len(nums)))
	return res
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n*n!)$$，$$n$$为`nums`数组长度
- **空间复杂度：** 空间复杂度为$$O(n*n!)$$，$$n$$为`nums`数组长度
