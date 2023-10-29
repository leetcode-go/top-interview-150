package main

import "fmt"

// https://leetcode.cn/problems/implement-trie-prefix-tree
func main() {
	oper := []string{"Trie", "insert", "search", "search", "startsWith", "insert", "search"}
	value := [][]string{nil, {"apple"}, {"apple"}, {"app"}, {"app"}, {"app"}, {"app"}}
	var handler Trie
	for idx, item := range oper {
		switch item {
		case "Trie":
			handler = Constructor()
			fmt.Print(" null")
		case "insert":
			handler.Insert(value[idx][0])
			fmt.Print(" null")
		case "search":
			fmt.Printf(" %v", handler.Search(value[idx][0]))
		case "startWith":
			fmt.Printf(" %v", handler.StartsWith(value[idx][0]))
		}
	}
}

type Trie struct {
	subChar []*Trie
	end     bool
}

func Constructor() Trie {
	return Trie{
		subChar: make([]*Trie, 26),
	}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, char := range word {
		idx := char - 'a'
		if node.subChar[idx] == nil {
			sub := Constructor()
			node.subChar[idx] = &sub
		}
		node = node.subChar[idx]
	}
	node.end = true
}

func (this *Trie) search(word string) *Trie {
	node := this
	for _, char := range word {
		idx := char - 'a'
		if node.subChar[idx] == nil {
			return nil
		}
		node = node.subChar[idx]
	}
	return node
}
func (this *Trie) Search(word string) bool {
	node := this.search(word)
	if node == nil || !node.end {
		return false
	}
	return true
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.search(prefix) != nil
}
