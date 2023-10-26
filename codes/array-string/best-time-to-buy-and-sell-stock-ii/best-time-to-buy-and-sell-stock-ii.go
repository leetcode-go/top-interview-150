package main

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