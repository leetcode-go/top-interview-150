# 整数转罗马数字


题目链接: [https://leetcode.cn/problems/integer-to-roman](https://leetcode.cn/problems/integer-to-roman)

## 解题思路一：

1. 列出从高到低列出每个罗马数字对应的阿拉伯数字，额外再加上4、9、40、90、400、900对应的罗马数字
2. 逐个取出阿拉伯数字，然后用num循环减去该数字，直至num小于阿拉伯数字再取下一个，每减一次，则在结果字符串上追加一个对应的罗马数字


```go
var (
	values = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
)

func intToRoman(num int) string {
	res := ""
	for _, value := range values {
		for num >= value {
			num -= value
			switch value {
			case 1000:
				res += "M"
			case 900:
				res += "CM"
			case 500:
				res += "D"
			case 400:
				res += "CD"
			case 100:
				res += "C"
			case 90:
				res += "XC"
			case 50:
				res += "L"
			case 40:
				res += "XL"
			case 10:
				res += "X"
			case 9:
				res += "IX"
			case 5:
				res += "V"
			case 4:
				res += "IV"
			case 1:
				res += "I"
			}
		}
		if num == 0 {
			break
		}
	}
	return res
}
```

## 复杂度分析一

- **时间复杂度：** 对 $$values$$ 进行了一次遍历，又在其中对 $$num$$ 进行了循环处理，时间复杂度为 $$O(n^2)$$
- **空间复杂度：** $$values$$ 属于预定义数据，占用空间排除在算法外,空间复杂度为 $$O(1)$$

## 解题思路二：

1. 列举出1到9,10-90,100-900，1000-3000对应的每个层级的罗马字符
2. 对num进行处理，计算出每个层级对应的数字取出相应的罗马字符进行拼接


```go
var (
	thousand = []string{"", "M", "MM", "MMM"}
	hundred  = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	ten      = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	one      = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VII", "IX"}
)

func intToRoman(num int) string {
	return thousand[num/1000] + hundred[num%1000/100] + ten[num%100/10] + one[num%10]
}
```

## 复杂度分析一

- **时间复杂度：** 时间复杂度为 $$O(1)$$
- **空间复杂度：** 空间复杂度为 $$O(1)$$