package main

import "fmt"

// https://leetcode.cn/problems/permutations
func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	var res [][]int
	var backtrack func([]int, []bool)
	backtrack = func(path []int, used []bool) {
		if len(path) == len(nums) {
			res = append(res, path)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			backtrack(append(path, nums[i]), used)
			used[i] = false
		}
	}
	backtrack([]int{}, make([]bool, len(nums)))
	return res
}
