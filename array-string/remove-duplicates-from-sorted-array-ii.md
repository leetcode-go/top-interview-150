# 删除有序数组中的重复项II

题目链接: [80. 删除有序数组中的重复项 II](https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/description/)

解题思路:

1. 思路与[26. 删除有序数组中的重复项](https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/)类似
2. 仅添加一个临时参数记录当前元素的个数，当大于两个时不做操作，元素更新时临时变量更新

```go
func removeDuplicates(nums []int) int {
    length, cnt, tmp := len(nums), 0, 1
    if length == 0 {
        return 0
    }
    for i := 1; i < length; i++ {
        if nums[i] == nums[cnt] {
            tmp++
            if tmp<3{
                cnt++
                nums[cnt] = nums[i]
            }
        } else {
            cnt++
            nums[cnt] = nums[i]
            tmp = 1
        }
    
    }
    return cnt + 1
}

```


