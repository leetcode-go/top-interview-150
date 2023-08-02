# 删除有序数组中的重复项II

题目链接: [80. 删除有序数组中的重复项 II](https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/description/)

解题思路:

1. 思路与[26. 删除有序数组中的重复项](https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/)类似
2. 针对本题采取了快慢指针的方式，快指针用来遍历数组，慢指针用来更新数组以及确保新数组元素个数

```go
func removeDuplicates(nums []int) int {
    length := len(nums)
    if length <= 2 {
        return length
    }
    slow,fast := 2,2
    for fast <length {
        if nums[slow-2]!=nums[fast]{
            nums[slow]=nums[fast]
            slow++
        }
        fast++
    }
    return slow
}

```


