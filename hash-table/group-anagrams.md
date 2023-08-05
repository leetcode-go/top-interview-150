# 字母异位词分组

题目链接: [https://leetcode.cn/problems/group-anagrams](https://leetcode.cn/problems/group-anagrams)

## 解题思路：

1. 对字符串数组中每个字符串都拆解成byte数组，并排序重组成新的字符串
2. 所有异位字符串排序后行成的新的字符串都是相同的字符串，可根据这个特新构建map，map的每个元素都是异位字符

```go
func groupAnagrams(strs []string) [][]string {
    mp:=map[string][]string{}
    for _,item:=range strs{
        s:=[]byte(item)
        sort.Slice(s,func(i,j int)bool{return s[i]>s[j]})
        sStr:=string(s)
        mp[sStr]=append(mp[sStr],item)
    }
    res:=make([][]string,0,len(mp))
    for _,item:=range mp{
        res=append(res,item)
    }
    return res
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度是 $$O(n^2)$$
- **空间复杂度：** 空间复杂度是 $$O(n)$$
