# 无重复字符的最长子串

题目链接: [https://leetcode.cn/problems/longest-substring-without-repeating-characters](https://leetcode.cn/problems/longest-substring-without-repeating-characters)

## 题目描述

给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

> 输入: s = "abcabcbb"
>
> 输出: 3 
>
>解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:

> 输入: s = "bbbbb"
>
> 输出: 1
>
> 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:

> 输入: s = "pwwkew"
>
> 输出: 3
>
> 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
>
>      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 

提示：

- 0 <= s.length <= 5 * 104
- s 由英文字母、数字、符号和空格组成

## 解题方法：

1. 从头遍历字符串，以每个字符为子串的起点，用right遍历字符串，直到找到任一字符在子串中已出现为子串结尾后的字符
2. 计算子串长度以及与记录的子串长度进行对比保留最大值
3. 以每个字符作为子串起点前，都需要将这个字符的前一个字符从map中移除，避免对后续的查找右边界的干扰

```go
func lengthOfLongestSubstring(s string) int {
    charMap:=map[byte]int{}
    right,res:=-1,0
    length:=len(s)
    for i:=0;i<length;i++{
        if i!=0{
            delete(charMap,s[i-1])
        }
        for right+1<length&&charMap[s[right+1]]==0{
            charMap[s[right+1]]++
            right++
        }
        if res<right-i+1{
            res=right-i+1
        }
    }
    return res
}
```

## 复杂度分析：

* **时间复杂度:** $$O(n^2)$$，$$n$$为$$s$$中字符的个数
* **空间复杂度:** $$O(n)$$
