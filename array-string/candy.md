# 分发糖果


题目链接: [https://leetcode.cn/problems/candy](https://leetcode.cn/problems/candy)

## 解题思路(与官解一致)：

1. 题目要求每个小孩必须至少一颗糖果，且相邻的两个小孩评分高的拿的更多
2. 因此分糖果的结果需要满足一个条件，无论从左到右遍历还是从右到左遍历，都满足一个条件，当前这个小孩与前一个小孩相比如果评分更高则拿的多，评分低则拿的少
3. 为了能得到最小值，则，当前小孩比前一个小孩评分低时直接拿最少1颗
4. 在从左往右分配及从右往左分配两个方向中每个小孩能拿到的最多的糖果数量即可满足上述条件


```go
func candy(ratings []int) int {
	length := len(ratings)
	candyList := make([]int, length)
	for idx, item := range ratings {
		if idx > 0 && item > ratings[idx-1] {
			candyList[idx] = candyList[idx-1] + 1
		} else {
			candyList[idx] = 1
		}
	}
	last, count := 0, 0
	for i := length - 1; i >= 0; i-- {
		if i < length-1 && ratings[i] > ratings[i+1] {
			last++
		} else {
			last = 1
		}
		if last > candyList[i] {
			count += last
		} else {
			count += candyList[i]
		}
	}
	return count
}
```

## 复杂度分析

- **时间复杂度：** 对数组进行了单轮遍历，因此时间复杂度为 $$O(n)$$，其中 $$n$$ 是数组 $$ratings$$ 的长度
- **空间复杂度：** 空间复杂度为 $$O(n)$$