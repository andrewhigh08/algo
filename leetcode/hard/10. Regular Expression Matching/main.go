// https://leetcode.com/problems/regular-expression-matching/description/
package main

import "fmt"

func main() {
	fmt.Println(isMatch("aa", "a"))  // false
	fmt.Println(isMatch("aa", "a*")) // true
	fmt.Println(isMatch("ab", ".*")) // true
}

func isMatch(s string, p string) bool {
	// TODO: решить
	return false
}
