# 用最少数量的箭箭引爆气球

题目链接: [https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons](https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons)

## 解题思路

1. 将`points`里面每个`point`按照右边界从小到大进行排序
2. 两个`point`能通过一支箭击穿的条件为两个`point`有交界，步骤一已经对`point`按照右边界进行排序，因此只要`point[i][1]>=point[i+1][0]`，则两个`point`能通过一支箭击穿
3. 从`point[0]`开始，取`point[0][1]`为初始`right`，后续遍历`points`，若`point[i][0]>right`，则`right`更新为`point[i][1]`，同时`count`加一

```go
func findMinArrowShots(points [][]int) int {
	if len(points) <= 1 {
		return len(points)
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1]<points[j][1]
	})
	right := points[0][1]
	count := 1
	for _, item := range points {
		if item[0]>right{
            right=item[1]
            count++
        }
	}
	return count
}
```

## 复杂度分析：

* **时间复杂度:** $$O(n)$$，$$n$$为$$intervals$$中元素的个数
* **空间复杂度:** $$O(1)$$
