# 颠倒二进制位

题目链接: [https://leetcode.cn/problems/reverse-bits](https://leetcode.cn/problems/reverse-bits)

## 解题思路：

1. 通过遍历 `32` 位二进制位，从右往左翻转二进制位
2. 在循环中，将结果变量 `res` 左移一位，然后将 `num` 的最后一位与 `1` 进行按位或运算，将结果存储到 `res` 中
3. 最后，将 `num` 右移一位，继续进行循环

```go
func reverseBits(num uint32) uint32 {
	var res uint32

	// 遍历 32 位二进制位
	for i := 0; i < 32; i++ {
		res <<= 1      // 将结果变量 res 左移一位
		res |= num & 1 // 将 num 的最后一位与 1 进行按位或运算，将结果存储到 res 中
		num >>= 1      // 将 num 右移一位
	}

	return res
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度是 $$O(1)$$，因为函数只需要遍历 `32` 位二进制位，进行位运算
- **空间复杂度：** 空间复杂度是 $$O(1)$$，因为函数只使用了常数个变量来存储中间结果，没有使用额外的空间
