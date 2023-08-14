# 直线上最多的点数

题目链接: [https://leetcode.cn/problems/max-points-on-a-line](https://leetcode.cn/problems/max-points-on-a-line)

## 解题思路：

枚举每一对点，计算它们构成的直线的斜率，然后将斜率相同的点分为一组，统计每组中点的个数，最后取所有组中点的个数的最大值作为答案

使用了哈希表来存储每个斜率对应的点的个数。由于斜率可能是一个分数，我们需要使用字符串来表示斜率。在计算斜率时，我们使用了欧几里得算法来计算两个点之间的最大公约数，以避免精度误差

**具体实现：**

1. 遍历每一对点，计算它们构成的直线的斜率
2. 将斜率相同的点分为一组，统计每组中点的个数
3. 取所有组中点的个数的最大值作为答案

```go
func maxPoints(points [][]int) int {
	n := len(points)
	if n <= 2 {
		return n
	}

	res := 0
	for i := 0; i < n; i++ {
		m := make(map[string]int) // 用于存储每个斜率对应的点的个数
		dup := 1                  // 用于记录与当前点重合的点的个数

		for j := i + 1; j < n; j++ {
			x, y := points[j][0], points[j][1]
			if x == points[i][0] && y == points[i][1] { // 如果当前点与当前枚举的点重合，则将 dup 加 1
				dup++
				continue
			}

			g := gcd(x-points[i][0], y-points[i][1])         // 计算两个点之间的最大公约数
			dx, dy := (x-points[i][0])/g, (y-points[i][1])/g // 计算斜率
			m[itoa(dx, dy)]++                                // 将斜率相同的点分为一组，统计每组中点的个数
		}

		// 更新最大值
		res = max(res, dup)
		for _, v := range m {
			res = max(res, v+dup)
		}
	}

	return res
}

func itoa(a, b int) string {
	if a == 0 {
		return "0"
	}

	if b == 0 {
		return "1"
	}

	if a < 0 {
		a, b = -a, -b
	}

	return strconv.Itoa(a) + "/" + strconv.Itoa(b)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度是 $$O(n^2)$$，其中 $$n$$ 是点的个数
- **空间复杂度：** 空间复杂度是 $$O(n)$$，算法使用了哈希表来存储每个斜率对应的点的个数
