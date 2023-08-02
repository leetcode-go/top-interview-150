# 轮转数组

题目链接: [https://leetcode.cn/problems/rotate-array/](https://leetcode.cn/problems/rotate-array/)

# 思路
> 拼接后裁剪

# 解题方法
1. 轮转无非是首尾相连，窗口滑动，这个最简单的思路就是将原数组与原数组拼接后获取对应的结果数组
2. 当然，需要注意边界问题

# 复杂度

- **时间复杂度:** $$O(n)$$
- **空间复杂度:** $$O(n)$$

```go
func rotate(nums []int, k int) {
    length := len(nums)
    k = k%length
    if k<length && k != 0 {
        copy(nums, append(nums[length-k:],nums[:length-k+1]...))
    }
}
```
