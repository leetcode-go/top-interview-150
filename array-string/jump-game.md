# 跳跃游戏


题目链接: [55. 跳跃游戏](https://leetcode.cn/problems/jump-game/description/)

## 解题思路：

1. 正向遍历:从头开始遍历，用一个遍历记录我们能够跳到的最远距离，逐步更新最远距离，若能到达目标位置，即为成功
2. 反向遍历:若一个点能到达目标节点，将其更新为目标节点，若最终目标节点为起点，说明可以从起点跳到终点

```go
//正向遍历
func canJump1(nums []int) bool {
    maxStep := 0
    n := len(nums)
    if n == 1{
        return true
    }
    for k,v := range nums {
        if maxStep >= k && k+v>maxStep {
            maxStep = k+v
        }
    }
    return maxStep >=n-1
}
//反向遍历
func canJump(nums []int) bool {
    n := len(nums)
    if n == 1{
        return true
    }
    maxStep := n-1
        
    for i:=n-2;i>=0;i-- {
        if nums[i]+i >= maxStep {
            maxStep =i  
        }
    }
    return maxStep == 0
}
```

## 复杂度分析

- **时间复杂度：** 只遍历了一遍数组 $$t$$，因此时间复杂度为 $$O(n)$$，其中 $$n$$ 是数组 $$nums$$ 的长度
- **空间复杂度：** 只使用了常数个变量，因此空间复杂度为 $$O(1)$$，无论输入的数组 $$nums$$ 的长度如何，函数使用的空间都是固定的
