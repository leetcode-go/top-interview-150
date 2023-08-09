# 单词拆分



题目链接: [https://leetcode.cn/problems/word-break/](https://leetcode.cn/problems/word-break/)

## 解题思路：

1. 若s可以利用字典拼出，则s去掉某个字典内的单词剩下的也可以用字典内出现的单词拼出
2. 可以利用哈希表存储字典单词
3. dp[i]=dp[j]&&dict(s[j:i])
```go
func wordBreak(s string, wordDict []string) bool {
    wordDictMap := make(map[string]bool)
    for _,v := range wordDict {
        wordDictMap[v]=true
    }
    n := len(s)
    dp := make([]bool,n+1)
    dp[0]=true
    for i := 1;i<=n;i++{
        for j:=0;j<i;j++{
            if dp[j] && wordDictMap[s[j:i]]{
                dp[i]=true
                break
            }
        }
    }
    return dp[n]
}

```

## 复杂度分析

- **时间复杂度：** 时间复杂度是 $$O(n²)$$
- **空间复杂度：** 空间复杂度是 $$O(n)$$
