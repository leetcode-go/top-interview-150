# 多数元素

题目链接: [https://leetcode.cn/problems/majority-element/](https://leetcode.cn/problems/majority-element/)

解题思路：

1. 多数元素是指出现的次数大于n/2的元素，因此排序后，第n/2位置的元素必然是多数

```go
func majorityElement(nums []int) int {
    sort.Ints(nums)
    return nums[len(nums)/2]
}
```