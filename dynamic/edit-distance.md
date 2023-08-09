# 编辑距离


题目链接: [https://leetcode.cn/problems/edit-distance/](https://leetcode.cn/problems/edit-distance/)

## 解题思路：


1. 设定二维数组dp[i][j],表示word1[i]转变到word2[j]需要的最少步数
2. word1[i]==word2[j] dp[i][j]=dp[i-1][j-1]
3. word1[i]!=word2[j] dp[i][j]=min(dp[i-1][j-1],dp[i-1][j],dp[i][j-1])+1
4. dp[i-1][j-1]表示替换，dp[i-1][j]表示删除，dp[i]j-1]表示插入
5. dp[0][j]需要特殊处理，代表word1是"",变成word2的最少步骤，插入步骤，插入word2长度
6. dp[i][0]需要特殊处理，代表把word1从原字符串变成word2为空的最少步骤，删除步骤，删除word1长度


```go
func minDistance(word1 string, word2 string) int {
    l1,l2 := len(word1),len(word2)
    dp := make([][]int,l1+1)
    for i:=0;i<l1+1;i++{
        dp[i]=make([]int,l2+1)
    }
    for i:=0;i<=l1;i++{
       dp[i][0]=i
    }
    for j:=0;j<=l2;j++{
        dp[0][j]=j
    }
    for i:=1;i<=l1;i++{
        for j:=1;j<=l2;j++{
            if word1[i-1]==word2[j-1]{
                dp[i][j]=dp[i-1][j-1]
            }else{
                dp[i][j]=min(dp[i-1][j-1],min(dp[i-1][j],dp[i][j-1]))+1
            }
           
        }
    }
    return dp[l1][l2]
}
func min(a,b int) int {
    if a<b{
       return a
    }
    return b
}

```

## 复杂度分析

- **时间复杂度：** 时间复杂度是 $$O(l1*l2)$$，其中 $$l1$$ 是字符串 `word1` 的长度, $$l2$$ 是字符串 `word2` 的长度
- **空间复杂度：** 空间复杂度是 $$O(l1*l2)$$，我们需要一个l1*l2二维数组存储每个位置的状态。
