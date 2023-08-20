# 加油站


题目链接: [https://leetcode.cn/problems/gas-station](https://leetcode.cn/problems/gas-station)

## 解题思路

1. 遍历数组，计算从第一个加油站开始到下一个加油站消耗的总得汽油数及获得的汽油数，判断是否满足汽油消耗且有剩余
2. 若有则继续，若无则从下一个加油站开始重复步骤1，判断是否能走到底(步骤一已经判断了前一段会导致汽油不足的情况，因此可以直接跳过)


```go
func canCompleteCircuit(gas []int, cost []int) int {
    for i,n:=0,len(gas);i<n;{
        sumGas,sumCost,idx:=0,0,0
        for idx<n{
            j:=(i+idx)%n
            sumGas+=gas[j]
            sumCost+=cost[j]
            if sumCost>sumGas{
                break
            }
            idx++
        }
        if idx==n{
            return i
        }else{
            i+=idx+1
        }
    }
    return -1
}
```

## 复杂度分析

- **时间复杂度：** 对数组进行了双重遍历，因此时间复杂度为 $$O(n^2)$$，其中 $$n$$ 是数组 $$gas$$ 的长度
- **空间复杂度：** 空间复杂度为 $$O(1)$$