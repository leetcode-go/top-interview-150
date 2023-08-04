# 同构字符串

题目链接: [https://leetcode.cn/problems/isomorphic-strings](https://leetcode.cn/problems/isomorphic-strings)

## 解题思路：

1. 同构的前提是s中每个出现的字符都可以映射到t中的字符，且同一个字符不能映射到不通的字符身上
2. 因此只需要遍历两个字符串，建立s与t的每个字符的映射表以及t与s的每个字符的映射表
3. 当出现t中的某个字符被s中的两个字符映射或s中的某个字符映射到了t中的两个字符时，则判定不为同构字符串

```go
func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
		return false
	}
	sToT := map[byte]byte{}
	tToS := map[byte]byte{}
	for idx := range s {
        x,y:=s[idx],t[idx]
        if (sToT[x]>0&&sToT[x]!=y) ||(tToS[y]>0&&tToS[y]!=x){
            return false
        }
		sToT[x] = y
		tToS[y] = x
	}
	return true
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度是 $$O(n)$$
- **空间复杂度：** 空间复杂度是 $$O(n)$$
