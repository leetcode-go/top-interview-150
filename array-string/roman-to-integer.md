# 罗马数字转整数

题目链接: [https://leetcode.cn/problems/roman-to-integer](https://leetcode.cn/problems/roman-to-integer)

## 解题思路：

1. 当小值在大值的左边，则减小值，如 `IV=5-1=4`
2. 当小值在大值的右边，则加小值，如 `VI=5+1=6`
3. 由上可知，右值永远为正，因此最后一位必然为正
4. 从左往右遍历，如果当前值小于右边的值，则减去当前值，否则加上当前值
5. 由于最后一位必然为正，因此最后一位必然会加上

```go
var m = map[rune]int{
	'I': 1,    // 1
	'V': 5,    // 5
	'X': 10,   // 10
	'L': 50,   // 50
	'C': 100,  // 100
	'D': 500,  // 500
	'M': 1000, // 1000
}

func romanToInt(s string) int {
	var ans int
	var prev int

	for _, r := range s {
		curr := m[r]
		if curr > prev {
			ans += curr - 2*prev // 如果当前值比前一个值大，则减去前一个值的两倍
		} else {
			ans += curr // 否则加上当前值
		}
		prev = curr // 更新前一个值
	}

	return ans
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为 $$O(n)$$，其中`n`是字符串`s`的长度
- **空间复杂度：** 空间复杂度为 $$O(1)$$，因为只需要常数级别的额外空间来存储映射表和一些变量
