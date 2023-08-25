# 最长公共前缀

题目链接: [https://leetcode.cn/problems/longest-common-prefix/](https://leetcode.cn/problems/longest-common-prefix)

## 题目描述
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1：

> 输入：strs = ["flower","flow","flight"]
>
> 输出："fl"

示例 2：

> 输入：strs = ["dog","racecar","car"]
>
> 输出：""
>
> 解释：输入不存在公共前缀。
 

提示：

- 1 <= strs.length <= 200
- 0 <= strs[i].length <= 200
- strs[i] 仅由小写英文字母组成

## 解题思路：

1. 以第一个元素为默认前缀，遍历所有元素，对每个元素的字符从前往后进行遍历，并与前缀进行对比
2. 若字符与前缀中相同位置的字符不想等，则缩减前缀，并退出循环，对下一个元素进行遍历

```go
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		str := strs[i]
		for j := 0; j < len(str) && j < len(prefix); j++ {
			if str[j] != prefix[j] {
				prefix = prefix[:j]
				break
			}
		}
		if len(str) < len(prefix) {
			prefix = str
		}
	}
	return prefix
}
```

## 复杂度分析

- **时间复杂度:** $$O(n)$$
- **空间复杂度:** $$O(1)$$