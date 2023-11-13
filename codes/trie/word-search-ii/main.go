package main

import "fmt"

type Trie struct {
	subTrie []*Trie
	word    string
}

func contract() *Trie {
	return &Trie{
		subTrie: make([]*Trie, 26),
	}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.subTrie[ch] == nil {
			node.subTrie[ch] = contract()
		}
		node = node.subTrie[ch]
	}
	node.word = word
}

// https://leetcode.cn/problems/word-search-ii
func main() {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'}}
	words := []string{"oath", "pea", "eat", "rain"}
	fmt.Println(findWords(board, words))
}

func findWords(board [][]byte, words []string) []string {
	if words == nil || len(words) == 0 || board == nil || len(board) == 0 {
		return nil
	}
	result := make([]string, 0)
	root := contract()
	directs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, word := range words {
		root.Insert(word)
	}
	m, n := len(board), len(board[0])
	seen := map[string]bool{}
	var dfs func(x int, y int, t *Trie)
	dfs = func(x int, y int, trie *Trie) {
		ch := board[x][y] - 'a'
		fmt.Println(board[x][y])
		trie = trie.subTrie[ch]
		if trie == nil {
			return
		}
		if trie.word != "" {
			seen[trie.word] = true
		}
		board[x][y] = '#'
		for _, direct := range directs {
			tmpx, tmpy := x+direct[0], y+direct[1]
			if tmpx >= 0 && tmpx < m && tmpy >= 0 && tmpy < n && board[tmpx][tmpy] != '#' {
				//fmt.Printf("data : %s\n", string(board[tmpx][tmpy]))
				dfs(tmpx, tmpy, trie)
			}
		}
		board[x][y] = ch + 'a'
	}
	for i, row := range board {
		for j := range row {
			dfs(i, j, root)
		}
	}

	for key := range seen {
		result = append(result, key)
	}
	return result
}
