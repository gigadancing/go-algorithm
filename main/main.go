package main

import (
	"fmt"
	"go-algorithm/tree/trie"
)

func main() {
	s := trie.NewKeyWordTree()
	s.Put(1, "ba")
	s.Put(2, "abd")
	s.Put(3, "acd")
	s.Debugout()
	fmt.Println(s.Search("a", 3))
}
