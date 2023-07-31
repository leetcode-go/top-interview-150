# 合并两个有序数组
[TOC]

# 思路一
> 贪心

# 解题方法
> 获取最左最小值，在每个阶段比较当前值与当前最小值的差值，若是比局部解大即可更新

# 复杂度
- 时间复杂度:
>  O(n)

- 空间复杂度:
>  O(1)

# Code
```Go []
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
> dp[i]代表前i天获取的最大利润，则dp[i]=max(dp[i-1],prices[i]-Min)

# 复杂度
- 时间复杂度:
>  $O(n)$

- 空间复杂度:
>  $O(n)$

# Code
```Go []
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