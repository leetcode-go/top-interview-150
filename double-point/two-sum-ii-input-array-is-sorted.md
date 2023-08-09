# 两数之和 II - 输入有序数组

题目链接: [https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted](https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted)

## 解题思路一

1. 双重遍历，假设每个元素是答案之一，计算这个元素与之后的每个元素的sum，判断是否与target相同

```go
func twoSum(numbers []int, target int) []int {
    for idx,number:=range numbers{
        tmp:=target-number
        for i:=idx+1;i<len(numbers);i++{
            if numbers[i]==tmp{
                return []int{idx+1,i+1}
            }
        }
    }
    return nil
}
```

## 复杂度分析一

- **时间复杂度：** 时间复杂度为 $$O(n^2)$$，其中 $$n$$ 是数组 `numbers` 的长度。
- **空间复杂度：** 只使用了常数个变量，因此空间复杂度为 $$O(1)$$，函数只使用了常数个变量来存储中间结果，没有使用额外的空间

## 解题思路二

1. 双指针，取两个指针分别指向数组头尾，依次计算numbers[l]+numbers[r]
2. 如果和大于target，则右指针左移一位。若和小于target，则左指针右移一位。若相等返回结果

```go
func twoSum(numbers []int, target int) []int{
    l,r:=0,len(numbers)-1
    for l<r{
        tmp := numbers[l]+numbers[r]
        if tmp==target{
            return []int{l+1,r+1}
        }else if tmp>target{
            r--
        }else{
            l++
        }
    }
    return nil
}
```

## 复杂度分析二

- **时间复杂度：** 时间复杂度为 $$O(n)$$，其中 $$n$$ 是数组 `numbers` 的长度。
- **空间复杂度：** 只使用了常数个变量，因此空间复杂度为 $$O(1)$$，函数只使用了常数个变量来存储中间结果，没有使用额外的空间