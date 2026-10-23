// https://leetcode.com/problems/word-break/description/
package main

import "fmt"

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))                       // true
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))                  // true
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})) // false
}

func wordBreak(s string, wordDict []string) bool {
	// TODO: решить
	return false
}
