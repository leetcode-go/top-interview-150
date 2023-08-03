# 赎金信

题目链接: [https://leetcode.cn/problems/ransom-note/](https://leetcode.cn/problems/ransom-note/)

## 解题思路：

1. 统计magazine里面每个字符出现的次数,统一用map进行存储
2. 遍历ransomNote,对于ransomNote中出现的每个字符，从步骤1的map中扣减相应的次数，如果map中某个字符的数量扣减至负数，则说明个数不足，magazine无法构成ransomNote

## 答案

```go
func canConstruct(ransomNote, magazine string) bool {
    if len(ransomNote) > len(magazine) {
        return false
    }
    magazineMap := [26]int{}
    for _, item := range magazine {
        magazineMap[item-'a']++
    }
    for _, item := range ransomNote {
        magazineMap[item-'a']--
        if magazineMap[item-'a'] < 0 {
            return false
        }
    }
    return true
}
```

## 复杂度分析

* 时间复杂度：$$O(n)$$，其中 $$n$$ 是数组中的元素数量
* 空间复杂度：$$O(n)$$，其中 $$n$$ 是数组中的元素数量，`map` 开销
