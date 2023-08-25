# 除自身以外数组的乘积

题目链接: [https://leetcode.cn/problems/product-of-array-except-self](https://leetcode.cn/problems/product-of-array-except-self)


## 解题方法：
1. 分别计算每个元素的前缀乘积以及后缀乘积，再将每个数的前缀乘积与后缀乘积相乘


```go
func productExceptSelf(nums []int) []int {
    length:=len(nums)
    left,right:=make([]int,length),make([]int,length)
    left[0]=1
    right[length-1]=1
    for i:=length-2;i>=0;i--{
        right[i]=right[i+1]*nums[i+1]
    }
    res:=make([]int,length)
    for i := 0; i < length; i++ {
		if i > 0 {
            // 求结果的同时计算前缀乘积
			left[i] = left[i-1] * nums[i-1]
		}
		res[i] = left[i] * right[i]
	}
    return res
}
```

## 复杂度分析：

- **时间复杂度:** $$O(n)$$，$$n$$为$$nums$$中元素的个数
- **空间复杂度:** $$O(n)$$