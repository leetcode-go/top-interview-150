# 多数元素

### 题目地址

题目链接: [https://leetcode.cn/problems/majority-element/](https://leetcode.cn/problems/majority-element/)

## 解题思路

### 思路一(排序)
1. 多数元素是指出现的次数大于n/2的元素，因此排序后，第n/2位置的元素必然是多数

### 答案一
```go
func majorityElement(nums []int) int {
    sort.Ints(nums)
    return nums[len(nums)/2]
}
```

### 思路二(统计)

1. 循环遍历数组，统计每个元素的个数
2. 遍历统计结果，判断元素个数是否大于n/2
3. 若大于则为多数

### 答案二

```go
func majorityElement(nums []int) int {
    countMap := map[int]int{}
    for _,item:=range nums{
        countMap[item]+=1
    }
    for key,value:=range countMap{
        if value>len(nums)/2{
            return key
        }
    }
    return 0
}
```