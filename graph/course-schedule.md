# 课程表

题目链接: [https://leetcode.cn/problems/course-schedule](https://leetcode.cn/problems/course-schedule)

## 解题思路：

1. 先翻转`preprequisites`数组的每个元素，计算每个课程完成后后续可执行的课程列表，一个类似图的数组
2. 遍历`tmpGraph`数组，对每个课程进行深度优先遍历，如果遍历过程中遇到已遍历的课程，则说明存在环，返回`false`
3. 遍历结束后，返回`true`

```go
func canFinish(numCourses int, prerequisites [][]int) bool {
	tmpGraph := make([][]int, numCourses)
	visited := make([]int, numCourses)
	res := true

	// 翻转每条边的起始，构建出一个图，图的每条边表示完成这个课程后，后续可继续完成的课程
	for _, item := range prerequisites {
		tmpGraph[item[1]] = append(tmpGraph[item[1]], item[0])
	}

	var dfs func(i int)

	dfs = func(i int) {
		visited[i] = 1
		for _, j := range tmpGraph[i] {
			if visited[j] == 0 {
				dfs(j)
				if !res {
					return
				}
			} else if visited[j] == 1 {
				res = false
				return
			}
		}
		visited[i] = 2
		return
	}

	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
		if !res {
			break
		}
	}
	return res
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n)$$,$$n$$为`prerequisites`的长度
- **空间复杂度：** 空间复杂度为$$O(n)$$,$$n$$为`prerequisites`的长度
