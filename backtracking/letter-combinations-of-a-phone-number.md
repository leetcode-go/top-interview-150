# 电话号码的字母组合

题目链接: [https://leetcode.cn/problems/letter-combinations-of-a-phone-number](https://leetcode.cn/problems/letter-combinations-of-a-phone-number)

## 解题思路：

1. 构建一个`map`存储每个数字对应的字母列表
2. 遍历`digits`, 找出每个数字对应的字母列表, 并逐个尝试拼接到字符串上,并传递回调函数继续下一个`digit`的处理
3. 如果`digits`遍历完成,则将结果填入`res`中返回

```go
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	byteMap := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	result := make([]string, 0)
	var backtrack func(idx int, last string)
	backtrack = func(idx int, last string) {
		if idx == len(digits) {
			result = append(result, last)
			return
		}
		byteItemList := byteMap[digits[idx]]
		for _, item := range byteItemList {
			backtrack(idx+1, last+string(item))
		}
	}
	backtrack(0, "")
	return result
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(m*n*8)$$,$$m$$为矩阵的行数,$$n$$为矩阵的列数
- **空间复杂度：** 空间复杂度为$$O(m*n*8)$$,$$m$$为矩阵的行数,$$n$$为矩阵的列数
