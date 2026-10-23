// https://leetcode.com/problems/word-ladder/description/
package main

import "fmt"

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})) // 5
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))        // 0
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// TODO: решить
	return 0
}
