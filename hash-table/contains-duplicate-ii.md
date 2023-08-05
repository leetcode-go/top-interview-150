# 存在重复元素 II

题目链接: [https://leetcode.cn/problems/contains-duplicate-ii](https://leetcode.cn/problems/contains-duplicate-ii)

## 解题思路一：

1. 判断数组中每个元素与其之后的k个元素是否有相同数字
2. 若有，则返回true

```go
func containsNearbyDuplicate(nums []int, k int) bool {
    for i:=0;i<len(nums);i++{
        for j:=i+1;j<=i+k&&j<len(nums);j++{
            if nums[i]==nums[j]{
                return true
            }
        }
    }
    return false
}
```

## 复杂度分析一

- **时间复杂度：** 时间复杂度是 $$O(n^2)$$
- **空间复杂度：** 空间复杂度是 $$O(1)$$

## 解题思路二：

1. 遍历每个元素，用map记录每个元素出现的位置
2. 当第二次出现元素时，判断当前元素与之前最后一次出现的元素的位置距离是否小于k，若是则返回true

```go
func containsNearbyDuplicate(nums []int, k int) bool {
    pos := map[int]int{}
    for i, num := range nums {
        if p, ok := pos[num]; ok && i-p <= k {
            return true
        }
        pos[num] = i
    }
    return false
}
```

## 复杂度分析二

- **时间复杂度：** 时间复杂度是 $$O(n)$$
- **空间复杂度：** 空间复杂度是 $$O(n)$$