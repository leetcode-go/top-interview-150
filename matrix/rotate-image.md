# 旋转矩阵

题目链接: [https://leetcode.cn/problems/rotate-image/](https://leetcode.cn/problems/rotate-image/)

## 解题思路：

1. 矩阵旋转时，原有的`(row,coll)`位置的元素经过旋转，会变换到`(coll,n-row-1)`的位置
2. 所以我们可以逐行遍历矩阵，将每行元素旋转复制到一个临时矩阵中，在把最终结果复制回原矩阵

```go
func rotate(matrix [][]int) {
	if matrix == nil || len(matrix) == 0 || len(matrix) == 1 {
		return
	}
	n := len(matrix)
	tmp := make([][]int, n)
	for idx := range tmp {
		tmp[idx] = make([]int, n)
	}
	for i, item := range matrix {
		for j, cell := range item {
			tmp[j][n-i-1] = cell
		}
	}
	copy(matrix, tmp)
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n^2)$$,$$n$$为矩阵的行数
- **空间复杂度：** 空间复杂度为$$O(n^2)$$,$$n$$为矩阵的行数
