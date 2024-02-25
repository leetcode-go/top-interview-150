package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}

var (
	dict  map[string]int
	graph [][]int
)

func ladderLength(beginWord, endWord string, wordList []string) int {
	dict = map[string]int{}
	graph = make([][]int, 0)
	for _, item := range wordList {
		insertEdge(item)
	}
	endIdx, has := dict[endWord]
	if !has {
		return 0
	}
	beginIdx := insertEdge(beginWord)
	dist := make([]int, len(dict))
	// 初始化数据为int最大值
	for idx := range dist {
		dist[idx] = math.MaxInt
	}
	dist[beginIdx] = 0
	queue := []int{beginIdx}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur == endIdx {
			// 因为整条路线是通过虚拟字符串拼接起来的,所以实际路线长度应该是长度/2 + 1,+1是要把开始节点算上
			return dist[endIdx]/2 + 1
		}
		for _, next := range graph[cur] {
			// 如果有具体数值，则不在重复增加长度
			if dist[next] == math.MaxInt {
				dist[next] = dist[cur] + 1
				queue = append(queue, next)
			}
		}
	}
	return 0
}

func insertWord(word string) int {
	idx, has := dict[word]
	if !has {
		idx = len(dict)
		dict[word] = idx
		graph = append(graph, []int{})
	}
	return idx
}

func insertEdge(word string) int {
	idx := insertWord(word)
	s := []byte(word)
	for i, item := range s {
		// 替换某一字符为*，构建虚拟字符串，两个字符串若只相差一个字符，那么可以通过这个虚拟字符串进行连接
		s[i] = '*'
		subIdx := insertWord(string(s))
		// 构建虚拟字符串到原字符串之间的连线，因为是单向图，所以需要建立两条连线
		graph[idx] = append(graph[idx], subIdx)
		graph[subIdx] = append(graph[subIdx], idx)
		s[i] = item
	}
	return idx
}
