# 爬楼梯


题目链接: [https://leetcode.cn/problems/climbing-stairs/](https://leetcode.cn/problems/climbing-stairs/)

## 解题思路：


1. 斐波那契数列 dp[i]=dp[i-1]+dp[i-2]

```go
func climbStairs(n int) int {
    dp := make([]int,n+1)
    dp[0]=1
    dp[1]=1
    for i:=2;i<=n;i++{
        dp[i]=dp[i-1]+dp[i-2]
    }
    return dp[n]
}

```

## 复杂度分析

- **时间复杂度：** 时间复杂度是 $$O(n)$$
- **空间复杂度：** 空间复杂度是 $$O(n)$$
