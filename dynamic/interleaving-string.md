# 交错字符串


题目链接: [https://leetcode.cn/problems/interleaving-string/](https://leetcode.cn/problems/interleaving-string/)

## 解题思路：
1. 由题意可知，s3是由s1和s2交错组成的,则可得s3 的前i+j个字符是由s1[:i]+s2[:j]交错组成的
2. 若s3的长度不是s1和s2的长度之和可以直接返回false
3. 若有一项为空，则判断另一项和最终字符串是否一致即可
4. 设定二维数组dp[i][j],代表s1的前i个字符和s2的前j个字符能否构成s3的前i+j个字符
5. dp[i][0]表示s1的前i个字符能否构成s3的前i个字符
6. dp[0][j]表示s3的前j个字符能否构成s3的前j个字符
7. 注意下标及边界，前i个字符串的下标为[0,i-1]
8. dp[i][j]=((dp[i][j-1]&&s2[j-1]==s3[i+j-1])||(dp[i-1][j]&&s1[i-1]==s3[i+j-1]))


```go
func isInterleave(s1 string, s2 string, s3 string) bool {
    l1, l2, l3 := len(s1), len(s2), len(s3)
    if l1==0 && l2==0 &&l3==0 {
        return true
    }
    if (l1 == 0 && s2!=s3) ||(l2==0 && s1!=s3){
        return false
    }
    if l3 != l1+l2  {
        return false
    }
    dp := make([][]bool, l1+1)
    for i := 0; i < l1+1; i++ {
        dp[i] = make([]bool, l2+1)
    }
    for i := 1; i < l1+1; i++ {
        if s1[:i] == s3[:i] {
            dp[i][0] = true
        }else{
            break
        }
    }
    
    for j := 1; j < l2+1; j++ {
        if s2[:j] == s3[:j] {
            dp[0][j] = true
        }else{
            break
        }
    }
    for i:=1;i<l1+1;i++{
        for j:=1;j<l2+1;j++{
            if (dp[i][j-1]&&s2[j-1]==s3[i+j-1])||(dp[i-1][j]&&s1[i-1]==s3[i+j-1]){
                dp[i][j]=true
            }
        }
        
    }
    
    return dp[l1][l2]
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度是 $$O(l1*l2)$$，其中 $$l1$$ 是字符串 `s1` 的长度,其中 $$l2$$ 是字符串 `s2` 的长度
- **空间复杂度：** 空间复杂度是 $$O(l1*l2)$$，我们需要一个l1*l2二维数组存储每个位置的状态。
