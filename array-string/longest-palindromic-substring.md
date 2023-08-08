# 最长回文子串

题目链接: [https://leetcode.cn/problems/longest-palindromic-substring](https://leetcode.cn/problems/longest-palindromic-substring)

## 解题方法

1. 遍历每个元素，假设从这个元素为回文中心，向左向右扩展查找以这个元素为中心的回文子串，并计算长度
2. 需要查找两次，因为回文串的格式可以是`aba`或`abba`


```go
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	start,end:=0,0
	for i := 0; i < len(s); i++ {
		left1,right1:=checkMaxLeftAndRight(s,i,i)
		left2,right2:=checkMaxLeftAndRight(s,i,i+1)
		if right1-left1>end-start{
			end,start=right1,left1
		}
		if right2-left2>end-start{
			end,start=right2,left2
		}
	}
	return s[start:end+1]
}

func checkMaxLeftAndRight(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left+1, right-1
}
```

## 复杂度分析

- **时间复杂度:** $$O(n^2)$$
- **空间复杂度:** $$O(1)$$
