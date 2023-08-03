# 跳跃游戏


题目链接: [860. 柠檬水找零](https://leetcode.cn/problems/lemonade-change/description/)

## 解题思路：

1. 简单模拟，因为没有比20更大的面值，所以20无需用于找零，只需要保存5元和10元的个数即可
2. 优先使用10元找零，防止后续出现10元无法找零的情况



```Go []

func lemonadeChange(bills []int) bool {
    five,ten := 0,0
    for _, v := range bills {
        if v == 5 {
            five++
        } else if v == 10 {
            if five == 0 {
                return false
            }
            five--
            ten++
        }else{
            if five>0&&ten>0{
                five--
                ten--
            }else if five >=3{
                five -= 3
            }else{
                return false
            }
        }
    }
    return true
}

```

## 复杂度分析

- **时间复杂度：** 只遍历了一遍数组 $$t$$，因此时间复杂度为 $$O(n)$$，其中 $$n$$ 是数组 $$bills$$ 的长度
- **空间复杂度：** 只使用了常数个变量，因此空间复杂度为 $$O(1)$$，无论输入的数组 $$bills$$ 的长度如何，函数使用的空间都是固定的
