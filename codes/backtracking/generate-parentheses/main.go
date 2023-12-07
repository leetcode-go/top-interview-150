package main

import "fmt"

// https://leetcode.cn/problems/generate-parentheses
func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	var res []string
	m := n * 2
	path := make([]byte, m)
	var backtrack func(i, open int)
	backtrack = func(i, open int) {
		if i == m {
			res = append(res, string(path))
		}
		if open < n {
			path[i] = '('
			backtrack(i+1, open+1)
		}
		if i-open < open {
			path[i] = ')'
			backtrack(i+1, open)
		}
	}
	backtrack(0, 0)
	return res
}
