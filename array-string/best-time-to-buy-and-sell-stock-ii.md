# 买卖股票的最佳时机II

题目链接: (https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/description/)


# 思路
> 动态规划

# 解题方法
  1.  看到题目的时候发现手持股票是有一个状态的，在这边应该使用二维dp数组 
  2. dp[i][j],i代表第i天，j有两个状态为0的时候代表手上不持有股票，为1的时候代表手上持有股票
  3. 我们需要的是最终的dp[n][0],因为这个时候手上不持有股票代表手上的现金肯定是比持有股票的时候多的
  4. 从题目中我们可以知道，手上只能持有一只股票，代表在同一个时间段内我们不能重复购入，要等我们的股票抛出才可以购入
    由此可以推出动态转移方程为：
  >  dp[i][0]=max(dp[i-1][0],dp[i-1][1]+prices[i])
  >  dp[i][1]=max(dp[i-1][1],dp[i-1[0]-prices[i]])


# 复杂度

- **时间复杂度:** $$O(n)$$
- **空间复杂度:** $$O(n)$$

```go
func maxProfit(prices []int) int {
    length := len(prices)
    if length == 0 {
        return 0
    }
    dp := make([][]int, length)
    for k, _ := range dp {
        dp[k] = make([]int, 2)
    }
    dp[0][0], dp[0][1] = 0, -prices[0]
    for i := 1;i<length; i++{   
		dp[i][0] = max(dp[i-1][0],prices[i]+dp[i-1][1])
        dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
    }
    return dp[length-1][0]
}
func max(a, b int)int {
    if a>b{
        return a
    }
    return b
}

```


