# 最小覆盖子串

题目链接: [https://leetcode.cn/problems/minimum-window-substring](https://leetcode.cn/problems/minimum-window-substring)

## 解题思路：

1. 用右指针遍历每个字符，并统计覆盖`t`中的字符的数量
2. 判断子串是否完全覆盖了字符串`t`
3. 若覆盖了，做右移左指针，缩短子串，判断子串是否仍然覆盖字符串`t`
4. 记录最小子串

```go
func minWindow(s string, t string) string {
	sLength, tLength := len(s), len(t)
	if sLength < tLength {
		return ""
	}
	resLen := math.MaxInt32
	resL, resR := -1, -1
	tMap, sMap := map[byte]int{}, map[byte]int{}
	for i := 0; i < tLength; i++ {
		tMap[t[i]]++
	}
	for l, r := 0, 0; r < sLength; r++ {
		if tMap[s[r]] > 0 {
			sMap[s[r]]++
		}
		for l <= r {
			match := true
			for k, v := range tMap {
				if sMap[k] < v {
					match = false
					break
				}
			}
			if !match {
				break
			}
			if r-l+1 < resLen {
				resLen = r - l + 1
				resL, resR = l, r+1
			}
			if _, ok := tMap[s[l]]; ok {
				sMap[s[l]]--
			}
			l++
		}
	}
	if resL == -1 {
		return ""
	}
	return s[resL:resR]
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n^2)$$
- **空间复杂度：** 空间复杂度为$$O(1)$$
