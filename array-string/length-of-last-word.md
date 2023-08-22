# 最后一个单词的长度


题目链接: [https://leetcode.cn/problems/length-of-last-word](https://leetcode.cn/problems/length-of-last-word)

## 解题思路：

1. 遍历字符串，查找两段空格之间的长度即可

```go
func lengthOfLastWord(s string) int {
	length := 0
	tmpLength := 0
	for _, item := range s {
		if item == ' ' {
			if tmpLength != 0 {
				length = tmpLength
			}
			tmpLength = 0
			continue
		} else {
			tmpLength++
		}
	}
	if tmpLength != 0 {
		length = tmpLength
	}
	return length
}
```

## 复杂度分析

- **时间复杂度：** 只遍历了一遍数组 $$s$$，因此时间复杂度为 $$O(n)$$
- **空间复杂度：** 只使用了常数个变量，因此空间复杂度为 $$O(1)$$
