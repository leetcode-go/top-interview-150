# 合并区间

题目链接: [https://leetcode.cn/problems/merge-intervals](https://leetcode.cn/problems/merge-intervals)

## 题目描述

以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

示例 1：

> 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
>
> 输出：[[1,6],[8,10],[15,18]]
>
> 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

示例 2：

> 输入：intervals = [[1,4],[4,5]]
>
> 输出：[[1,5]]
>
> 解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
 

提示：

- 1 <= intervals.length <= 104
- intervals[i].length == 2
- 0 <= starti <= endi <= 104

## 解题思路

1. intervals长度小于等于1直接返回
2. 对intervals进行排序，按区间的左边界进行升序排序
3. 遍历intervals，如果res的最后一个区间的右边界小于当前区间的左边界，则将当前区间加入res
4. 否则，将当前区间与res最后一个区间进行合并，合并后的区间右边界取两个区间的最大值

```go
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if res[len(res)-1][1] < intervals[i][0] {
			res = append(res, intervals[i])
		} else {
			if res[len(res)-1][1] < intervals[i][1] {
				res[len(res)-1][1] = intervals[i][1]
			}
		}
	}
	return res
}
```

## 复杂度分析：

* **时间复杂度:** $$O(n)$$，$$n$$为$$intervals$$中元素的个数
* **空间复杂度:** $$O(n)$$
