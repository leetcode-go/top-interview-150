package main

import "fmt"

// https://leetcode.cn/problems/combinations
func main() {
	fmt.Println(combine(5, 4))
}

func combine(n, k int) [][]int {
	var res [][]int
	var backtrack func(int, int, []int)
	backtrack = func(last, idx int, tmp []int) {
		tmpList := make([]int, len(tmp))
		copy(tmpList, tmp)
		tmpList = append(tmpList, last)
		if idx == k {
			res = append(res, tmpList)
			return
		}
		for i := last + 1; i <= n; i++ {
			backtrack(i, idx+1, tmpList)
		}
	}
	for i := 1; i <= n; i++ {
		backtrack(i, 1, []int{})
	}
	return res
}
