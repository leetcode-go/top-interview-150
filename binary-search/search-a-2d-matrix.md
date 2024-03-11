# 搜索二维矩阵

题目链接: [https://leetcode.cn/problems/search-a-2d-matrix](https://leetcode.cn/problems/search-a-2d-matrix)

## 解题思路：

1. 遍历每个单词，并将每个单词的某一个字符变为`*`，构建虚拟单词，并构建原单词与虚拟单词的连线
2. 若两个单词之间只差一个字符，那么必然能通过虚拟单词将两个单词连通
3. 因此只需要从`beginWord`开始，通过广度优先遍历，找到与`endWord`相同的虚拟单词，即可找到最短路径

```go
func searchMatrix(matrix [][]int, target int) bool {
	row := searchRow(matrix, target)
	left, right := 0, len(matrix[row])-1
	for left <= right {
		mid := (right-left)/2 + left
		if target > matrix[row][mid] {
			left = mid + 1
		} else if target < matrix[row][mid] {
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}

func searchRow(data [][]int, target int) int {
	left, right := 0, len(data)-1
	for left < right {
		rowMid := (right-left)/2 + left
		if data[rowMid][0] <= target {
			left = rowMid
		} else {
			right = rowMid - 1
		}
	}
	return left
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(log mn)$$, $$m$$为矩阵的行数，$$n$$为举证的列数。
- **空间复杂度：** 空间复杂度为$$O(1)$$。
