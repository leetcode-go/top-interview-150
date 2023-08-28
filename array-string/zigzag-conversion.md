# N字形变换

题目链接: [https://leetcode.cn/problems/zigzag-conversion](https://leetcode.cn/problems/zigzag-conversion)

# 解题思路

1. 矩阵模拟实现，每轮覆盖numRows-1列，一共numRows层，逐层进行字符追加即可

```go
func convert(s string, numRows int) string {
    r:=numRows
    if r==1||r>=len(s){
        return s
    }
    matrix:=make([][]byte,r)
    t,x:=r*2-2,0
    for i,char:=range s{
        matrix[x]=append(matrix[x],byte(char))
        if i%t<r-1{
            x++
        }else{x--}
    }
    return string(bytes.Join(matrix,nil))
}
```

## 复杂度分析

- **时间复杂度:** $$O(n)$$
- **空间复杂度:** $$O(n)$$
