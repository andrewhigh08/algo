// https://leetcode.com/problems/implement-trie-prefix-tree/description/
package main

import "fmt"

func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))   // true
	fmt.Println(trie.Search("app"))     // false
	fmt.Println(trie.StartsWith("app")) // true
	trie.Insert("app")
	fmt.Println(trie.Search("app")) // true
}

type Trie struct {
	// TODO: решить
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
}

func (this *Trie) Search(word string) bool {
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	return false
}
