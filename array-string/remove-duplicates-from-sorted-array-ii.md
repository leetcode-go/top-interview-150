# 删除有序数组中的重复项 II

题目链接: [https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii](https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii)

## 解题思路

1. 本题的关键就是如何判断一个元素出现的次数是否大于2
2. 联想上一题，那么我们的思路就可以是判断单前元素上上个元素是否相等，但是数组处理过程中可能会出现上一个元素刚好被我们要留下来的数据替换，导致错误的识别当前这个元素是第三次出现而抛弃
3. 我们可以通过一个变量将上一个保存的元素的定位记录下来，通过这个元素查找上上个保存下来的元素进行判断
4. 需要注意的是本次需要从数组的第3个元素开始进行处理，因此应先判断数组长度是否小于3

## 答案

```go
func removeDuplicates(nums []int) int {
    if len(nums)<3{
        return len(nums)
    }
    n,last:=2,2
    for last<len(nums){
        if nums[n-2]!=nums[last]{
            nums[n]=nums[last]
            n++
        }
        last++
    }
    return n
}
```
