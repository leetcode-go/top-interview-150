# 最长递增子序列


题目链接: [https://leetcode.cn/problems/longest-increasing-subsequence/](https://leetcode.cn/problems/longest-increasing-subsequence/)

## 解题思路：


1. 设定一维数组dp，dp[i]就是到nums[i]时，以nums[i]结尾的最长子序列长度
2. j<i if nums[i] > nums[j] dp[i]=max(dp[i],dp[j]+1)
3. nums的最长递增子序列的长度就是遍历dp[i]取最大值
```go
func lengthOfLIS(nums []int) int {
    maxLen := 0
    n := len(nums)
    if n<=1{
        return n
    }
    dp := make([]int,n+1)
    dp[0]=1
    for i:= 1;i<n;i++{
        dp[i]=1
        for j:=0;j<i;j++{
            if nums[i]>nums[j]{
                dp[i]=max(dp[i],dp[j]+1)
            }
        }
        maxLen = max(maxLen,dp[i])
    }
    return maxLen
}
func max(a,b int) int {
    if a>b{
        return a
    }
    return b
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度是 $$O(n²)$$
- **空间复杂度：** 空间复杂度是 $$O(n)$$
