# 课程表 II

题目链接: [https://leetcode.cn/problems/course-schedule-ii](https://leetcode.cn/problems/course-schedule-ii)

## 解题思路：

1. 先遍历统计每门课程依赖其他课程的数量，因为课程编号均是从0到`numCourses-1`，因此，用数组就可以进行存储
2. 同时构建每门课程被哪些课程依赖的关系拓扑
3. 从没有依赖的课程开始处理逐个进入结果数组，并根据关系拓扑对依赖该课程的课程的依赖数量进行扣减，课程依赖数量为0的课程进入队列
4. 重复3，直到队列为空

```go
func findOrder(numCourses int, prerequisites [][]int) []int {
	courseDeps := make([]int, numCourses)
	courseEdgeList := make([][]int, numCourses)
	for _, item := range prerequisites {
		courseDeps[item[0]]++
		courseEdgeList[item[1]] = append(courseEdgeList[item[1]], item[0])
	}

	queue := make([]int, 0)
	for idx, item := range courseDeps {
		if item == 0 {
			queue = append(queue, idx)
		}
	}
	var res []int
	for len(queue) > 0 {
		item := queue[0]
		res = append(res, item)
		queue = queue[1:]

		for _, v := range courseEdgeList[item] {
			courseDeps[v]--
			if courseDeps[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
    if len(res) != numCourses {
		return []int{}
	}
	return res
}
```

## 复杂度分析

- **时间复杂度：** 时间复杂度为$$O(n)$$,$$n$$课程数与先修课程数只和
- **空间复杂度：** 空间复杂度为$$O(n)$$,$$n$$为课程数与先修课程数只和
