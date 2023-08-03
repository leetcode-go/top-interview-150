# 三角形最小路径和


题目链接: [https://leetcode.cn/problems/triangle/](https://leetcode.cn/problems/triangle/)

## 解题思路：

1. 设定dp[i][j],代表自顶至下到[i][j]位置最小路径和
2. 二维数组为三角形,结合题干和提示可知,j<=i
3. dp[0][0]=triangle[0][0]
4. dp[i][j]=min(dp[i-1][j-1]+triangle[i][j],dp[i][j-1]+triangle[i][j])
5. dp[i][0]和dp[i][i]的位置需要特殊处理
6. 最终遍历dp[n-1][i]获取路径最小值

# Code
```Go []

func minimumTotal(triangle [][]int) int {
    n := len(triangle)
    if n == 0 {
        return 0
    }
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, i+1)
    }
    dp[0][0] = triangle[0][0]
    for i := 1; i < n; i++ {
        dp[i][0] = dp[i-1][0] + triangle[i][0]
        for j := 1; j < i; j++ {
            dp[i][j] = min(dp[i-1][j-1]+triangle[i][j], dp[i-1][j]+triangle[i][j])
        }
        dp[i][i] = dp[i-1][i-1] + triangle[i][i]
    }
    Min := 10001
    for i:=0;i<n;i++{
        Min =min(Min,dp[n-1][i])
    }
    return Min
}
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度是 $$O(n²)$$，其中 $$n$$ 是数组 `triangle` 的长度
- **空间复杂度：** 空间复杂度是 $$O(n²)$$，我们需要一个边长为n的等腰直角三角形大小的二维数组。
