package main

import "fmt"

// https://leetcode.cn/problems/combination-sum
func main() {
	res := combinationSum([]int{2, 3, 6, 7}, 7)
	for _, item := range res {
		fmt.Println(item)
	}
}

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var comb []int
	var backTracking func(idx, tmpTarget int)
	backTracking = func(idx, tmpTarget int) {
		if idx == len(candidates) {
			return
		}
		if tmpTarget == 0 {
			res = append(res, append([]int{}, comb...))
			return
		}
		backTracking(idx+1, tmpTarget)
		if tmpTarget-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			backTracking(idx, tmpTarget-candidates[idx])
			comb = comb[:len(comb)-1]
		}
	}
	backTracking(0, target)
	return res
}
