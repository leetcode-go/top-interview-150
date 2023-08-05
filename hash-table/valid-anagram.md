# 有效的字母异位词

题目链接: [https://leetcode.cn/problems/valid-anagram](https://leetcode.cn/problems/valid-anagram)

## 解题思路一：

1. 异位词的要求是s与t两个字符串里面所有字符出现的次数一致
2. 只需要统计并对比s与t内各个字符的数量是否相等即可

```go
func isAnagram(s string, t string) bool {
    if len(s)!=len(t){
        return false
    }
    sMap:=map[rune]int{}
    // 统计s里面每个字符出现的次数
    for _,item:=range s{
        sMap[item]+=1
    }
    // 遍历t，逐次减去sMap中的次数，如果小于0则说明t中某个字符出现的次数大于t
    for _,item:=range t{
        sMap[item]-=1
        if sMap[item]<0{
            return false
        }
    }
    return true
}

```

## 复杂度分析一

- **时间复杂度：** 时间复杂度是 $$O(n)$$
- **空间复杂度：** 空间复杂度是 $$O(n)$$

## 解题思路二：

1. 先将每个字符串转成byte数组，并进行排序
2. 将排序后的byte数组转成string，若异位词，则排序后应该形成相同的字符串，所以只用对比排序后转换成的字符串是否相等即可


```go
func isAnagram(s string, t string) bool {
    s1, s2 := []byte(s), []byte(t)
    sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
    sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
    return string(s1) == string(s2)
}
```

## 复杂度分析二

- **时间复杂度：** 时间复杂度是 $$O(n)$$，将byte数组转成字符串或字符串转成byte数组需要的时间服务度均为$$O(n)$$，排序的时间复杂度根据golang底层实现判断
- **空间复杂度：** 空间复杂度是 $$O(n)$$
