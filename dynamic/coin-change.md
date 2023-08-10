# 零钱兑换


题目链接: [https://leetcode.cn/problems/coin-change/](https://leetcode.cn/problems/coin-change/)

## 解题思路：

1. 设定一维数组dp,dp[i]代表金额为i的时候所需要的最小硬币个数
2. dp[i]=min(dp[i-coins[j]])+1
3. 最大值选择math.MaxInt会有超int问题，故再次max选择amount+1


```go
func coinChange(coins []int, amount int) int {
    if amount == 0 {
        return 0
    }
    dp := make([]int, amount+1)
    for i := 0; i <= amount; i++ {
        dp[i] = amount+1
    }
    sort.Ints(coins)
    n := len(coins)
    dp[0]=0
    for i := 1; i <= amount; i++ {
        for j := 0; j < n; j++ {
            if i < coins[j] {
                break
            }
            dp[i] = min(dp[i], dp[i-coins[j]]+1)
        }
    }
    if dp[amount] >amount {
        return -1
    }
    return dp[amount]
}

func min(a,b int)int {
    if a<b {
        return a
    }
    return b
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度是 $$O(n*amount)$$ 其中 $$n$$ 是数组 `coins` 的数量,也就是不同面额的硬币的数量
- **空间复杂度：** 空间复杂度是 $$O(amount)$$
