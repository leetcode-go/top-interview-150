# 打家劫舍


题目链接: [https://leetcode.cn/problems/house-robber/](https://leetcode.cn/problems/house-robber/)

## 解题思路：


1. 报警条件,连续偷了两间房
2. 设定一维dp数组,dp[i]表示到了第i间房偷窃到的最大金额,不一定偷了第i间房
3. 状态转移方程:dp[i]=max(dp[i-2]+nums[i],dp[i-1])
4. dp[i-2]+nums[i]表示不偷上一家获取到的最大金额
5. dp[i-1]表示当前房屋不偷获取的最大金额
6. dp[0]=0表示没有偷，dp[1]表示偷了第一家房子也就是nums[0]

```go
func rob(nums []int) int {
    n := len(nums)
    if n == 1{
        return nums[0]
    }else if n == 2 {
        return max(nums[0],nums[1])
    }
    dp := make([]int ,n+1)
    dp[0]=0
    dp[1]=nums[0]
    for i :=2;i<=n;i++{
        dp[i]=max(dp[i-2]+nums[i-1],dp[i-1])
    }
    return dp[n]
}
func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度是 $$O(n)$$
- **空间复杂度：** 空间复杂度是 $$O(n)$$
