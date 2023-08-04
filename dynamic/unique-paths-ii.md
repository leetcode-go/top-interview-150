# 不同路径 II

题目链接: [https://leetcode.cn/problems/unique-paths-ii/](https://leetcode.cn/problems/unique-paths-ii/)

## 解题思路：

1. 设定dp[i][j],代表自[0][0]到[i][j]位置路径数量和
2. dp[0][0]=1
3. 机器人只能往右或往下走,在无障碍时，dp[i][j]=dp[i-1][j]+dp[i][j-1],在有障碍时，dp[i][j]=0
4. dp[i][0]和dp[0][j]需要特殊处理
5. 在无障碍时dp[i][0]=dp[i-1][0],dp[0][j]=[0][j-1]
6. 最终遍历dp[n-1][m-1]为最终结果

```go
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    n, m := len(obstacleGrid), len(obstacleGrid[0])
    if obstacleGrid[0][0] == 1 {
        return 0
    }
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, m)
    }
    dp[0][0] = 1
    for i := 1; i < n; i++ {
        if obstacleGrid[i][0] == 1 {
            break
        }
        dp[i][0] = dp[i-1][0]
    }
    for j := 1; j < m; j++ {
        if obstacleGrid[0][j] == 1 {
            break
        }
        dp[0][j] = dp[0][j-1]
    }
    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            if obstacleGrid[i][j] == 0 {
                dp[i][j] = dp[i-1][j] + dp[i][j-1]
            }
        }
    }
    return dp[n-1][m-1]
}

```

## 复杂度分析

- **时间复杂度：** 时间复杂度是 $$O(n*m)$$，其中 $$n$$ 是数组 `obstacleGrid` 的长度 ,$$m$$ 是数组`obstacleGrid[0]` 的长度
- **空间复杂度：** 空间复杂度是 $$O(n*m)$$，我们需要一个大小为n*m的数组来保存所有的状态
