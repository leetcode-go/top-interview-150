# 插入区间

题目链接: [https://leetcode.cn/problems/insert-interval](https://leetcode.cn/problems/insert-interval)

## 解题思路

1. 遍历`intervals`中的每个区间，如果当前区间与新区间没有交集，直接将当前区间加入结果集，因为区间是按照左边界升序的，因此，只要左边界大于新区间的右边界，就说明当前区间与新区间没有交集，当前区间的右边界小于新区间的左边界，也说明当前区间与新区间没有交集
2. 如果当前区间与新区间有交集，则将当前区间与新区间合并，合并后的区间为`[min(start1, start2), max(end1, end2)]`
3. 合并后的区间与新区间有交集，则继续合并，直到当前区间与新区间没有交集
4. 合并后的区间加入结果集

```go
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	res := make([][]int, 0)
	inserted := false
	for _, item := range intervals {
		if item[1] < newInterval[0] {
			res = append(res, item)
			continue
		}
		if item[0] > newInterval[1] {
			if !inserted {
				res = append(res, newInterval)
				inserted = true
			}
			res = append(res, item)
			continue
		}
		if item[0] < newInterval[0] {
			newInterval[0] = item[0]
		}
		if item[1] > newInterval[1] {
			newInterval[1] = item[1]
		}
	}
    if !inserted{
        res=append(res,newInterval)
    }
	return res
}
```

## 复杂度分析：

* **时间复杂度:** $$O(n)$$，$$n$$为$$intervals$$中元素的个数
* **空间复杂度:** $$O(n)$$
