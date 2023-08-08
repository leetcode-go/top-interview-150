# 最长回文子串


题目链接: [https://leetcode.cn/problems/longest-palindromic-substring/](https://leetcode.cn/problems/longest-palindromic-substring/)

## 解题思路：

1. 设定二维数组dp dp[i][j]:表示[i,j]的子串是回文子串
2. 有两种情况s[i]=s[j] dp[i][j]=true s[i]!=s[j] dp[i][j]=false
3. 当s[i]=s[j]时，有以下几种情况:
```text
a. i=j,单字符必定为回文串
b. j-i<=2 ,dp[i][j]=true
c. i和j相差大于一 如果s[i]=s[j]则dp[i][j]=dp[i+1]dp[j-1]
```
4. 注意遍历的时候先j后i，我们是需要维护dp[i][j] j>=i这一部分的状态

```go
func longestPalindrome(s string) string {
    length:=len(s)
    if length <=1 {
        return s
    }
    
    dp := make([][]bool,length)
    for i:=0;i<length;i++{
        dp[i]=make([]bool,length)
    }
    star :=0
    Maxlen:=1
    for i:=0;i<length;i++{
        dp[i][i]=true
    }
    for j:=1;j<length;j++{
        for i:=0;i<j;i++{
            if s[i]==s[j] {
                if j-i<=2{
                    dp[i][j]=true
                }else {
                    dp[i][j]= dp[i+1][j-1]
                }
            }
            if dp[i][j] && j-i+1 >Maxlen{
                Maxlen=j-i+1
                star=i
            }
        }
    }
    return s[star:star+Maxlen]
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度是 $$O(n^2)$$，其中 $$n$$ 是字符串 `s` 的长度
- **空间复杂度：** 空间复杂度是 $$O(n^2)$$，我们需要一个n*n二维数组存储每个位置的状态。
