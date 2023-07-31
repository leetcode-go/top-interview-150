# 合并两个有序数组

题目链接: [https://leetcode.cn/problems/merge-sorted-array/](https://leetcode.cn/problems/merge-sorted-array/)

解题思路：

1. 题目要求将nums2合并入nums1，时间复杂为O(n+m)，因此必然不能是从前往后合并，从前往后合并需要插入后将后续的元素逐个后移，时间复杂度遍不会是O(n+m)
2. nums1中在数组后预留了足够的空间用于存放nums2，因此在合并数组时可以不考虑将nums2数据写入nums1中时覆盖未处理数据
3. 此处仅需从后往前逐个对比nums1以及nums2的数据，写入大的数据，并下标减一即可，具体实现如下

```go
func merge(nums1 []int, m int, nums2 []int, n int)  {
    // 从nums1末尾开始写入
    length:=len(nums1)-1
    m=m-1
    n=n-1
    for ;n>=0&&m>=0;{
        if nums1[m]>nums2[n]{
            nums1[length]=nums1[m]
            m--
        }else{
            nums1[length]=nums2[n]
            n--
        }
        length--
    }
    // 若nums1中数据已经完全处理完，则直接灌入nums2的数据
    for ;n>=0;n--{
        nums1[length]=nums2[n]
        length--
    }
}
```

也可以使用一些函数

```go
func merge(nums1 []int, m int, nums2 []int, n int)  {
    copy(nums1[m:], nums2)
    sort.Ints(nums1)
}
```
