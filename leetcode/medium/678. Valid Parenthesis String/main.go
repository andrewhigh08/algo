// https://leetcode.com/problems/valid-parenthesis-string/description/
package main

import "fmt"

func main() {
	fmt.Println(checkValidString("()"))   // true
	fmt.Println(checkValidString("(*)"))  // true
	fmt.Println(checkValidString("(*))")) // true
	fmt.Println(checkValidString("("))    // false
}

func checkValidString(s string) bool {
	// TODO: решить
	return false
}
