// https://leetcode.com/problems/design-add-and-search-words-data-structure/description/
package main

import "fmt"

func main() {
	dict := Constructor()
	dict.AddWord("bad")
	dict.AddWord("dad")
	dict.AddWord("mad")
	fmt.Println(dict.Search("pad")) // false
	fmt.Println(dict.Search("bad")) // true
	fmt.Println(dict.Search(".ad")) // true
	fmt.Println(dict.Search("b..")) // true
}

type WordDictionary struct {
	// TODO: решить
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
}

func (this *WordDictionary) Search(word string) bool {
	return false
}
