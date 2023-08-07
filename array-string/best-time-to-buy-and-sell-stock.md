# 买卖股票的最佳时机

题目链接: [https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/)

# 思路一
> 贪心

# 解题方法
 获取最左最小值，在每个阶段比较当前值与当前最小值的差值，若是比局部解大即可更新

# 复杂度

- **时间复杂度:** $$O(n)$$
- **空间复杂度:** $$O(1)$$

```go
func maxProfit(prices []int) int {
    Max,Min := 0,math.MaxInt
    for _,v := range prices {
        Max = max(Max,v-Min)
        Min = min(Min,v)
    }
    return Max
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

# 思路二
> 动态规划

# 解题方法
> `dp[i]` 代表前 $$i$$ 天获取的最大利润，则 `dp[i]=max(dp[i-1],prices[i]-min)`

# 复杂度

- **时间复杂度:** $$O(n)$$
- **空间复杂度:** $$O(n)$$

```go
func maxProfit(prices []int) int {
    length := len(prices)
    if length == 0 {
        return 0
    }
    dp := make([]int,length)
    Min := prices[0]
    for i:=1;i<length;i++{
        Min = min(Min,prices[i])
        dp[i]=max(dp[i-1],prices[i]-Min)
    }
    return dp[length-1]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

# 思路三

# 解题方法

双重遍历，遍历计算买个买入日与其他卖出日的获利，并判断是否最大
> 对于最新的用例中，该解题方法会出现解答超时

# 复杂度

- **时间复杂度:** $$O(1)$$
- **空间复杂度:** $$O(n^2)$$

```go
func maxProfit(prices []int) int {
    max:=0
    for i:=0;i<len(prices);i++{
        for j:=i+1;j<len(prices);j++{
            tmpMax:=prices[j]-prices[i]
            if tmpMax>max{
                max=tmpMax
            }
        }
    }
    return max
}
```
