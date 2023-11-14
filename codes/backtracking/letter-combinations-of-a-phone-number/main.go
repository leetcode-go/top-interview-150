package main

import "fmt"

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number
func main() {
	fmt.Println(letterCombinations("23"))
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	byteMap := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	result := make([]string, 0)
	var backtrack func(idx int, last string)
	backtrack = func(idx int, last string) {
		if idx == len(digits) {
			result = append(result, last)
			return
		}
		byteItemList := byteMap[digits[idx]]
		for _, item := range byteItemList {
			backtrack(idx+1, last+string(item))
		}
	}
	backtrack(0, "")
	return result
}
