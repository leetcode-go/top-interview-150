# 移除元素

题目链接: [https://leetcode.cn/problems/remove-element/](https://leetcode.cn/problems/remove-element/)

解题思路：

1. 循环遍历数组
2. 遇到与val相等的不处理，不相等的从头按顺序写入数组，依次覆盖原数组内容

````go
func removeElement(nums []int, val int) int {
    if len(nums)==0{
        return 0
    }
    n:=0
    for i:=0;i<len(nums);i++{
        if nums[i]!=val{
            nums[n]=nums[i]
            n++
        }
    }
    return n
}
````
