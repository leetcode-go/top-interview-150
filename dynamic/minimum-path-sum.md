# 最小路径和


题目链接: [https://leetcode.cn/problems/minimum-path-sum/](https://leetcode.cn/problems/minimum-path-sum/)

## 解题思路：

1. 设定dp[i][j],代表自顶至下到[i][j]位置最小路径和
2. dp[0][0]=grid[0][0]
3. 由说明可知，每次只能向下或者向右移动一步,得状态转移方程为dp[i][j]=min(dp[i-1][j],dp[i][j-1])+v[i][j]
4. dp[i][0]和dp[0][j]需要特殊处理
5. dp[i][0]=dp[i-1][0]+v[i][j]
6. dp[0][j]=dp[0][i-1]+v[i][j]
7. 提示中告诉我们n,m>0

```go
func minPathSum(grid [][]int) int {
    n,m := len(grid),len(grid[0]) 
    dp := make([][]int,n)
    for i:=0;i<n;i++ {
        dp[i] =make([]int,m)
    }
    dp[0][0]=grid[0][0]
    for i:=1;i<n;i++{
        dp[i][0]=dp[i-1][0]+grid[i][0]
    }
    for j:=1;j<m;j++{
        dp[0][j]=dp[0][j-1]+grid[0][j]
    }
    for i:=1;i<n;i++{
        for j:=1;j<m;j++{
            dp[i][j]=min(dp[i-1][j],dp[i][j-1])+grid[i][j]
        }
    }
    return dp[n-1][m-1]
}
func min(a,b int) int {
    if a>b {
        return b
    }
    return a
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度是 $$O(n^2)$$，其中 $$n$$ 是数组 `grid` 的长度
- **空间复杂度：** 空间复杂度是 $$O(n^2)$$，我们需要一个n*n二维数组存储每个位置的状态。
