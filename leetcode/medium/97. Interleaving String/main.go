// https://leetcode.com/problems/interleaving-string/description/
package main

import "fmt"

func main() {
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac")) // true
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbbaccc")) // false
	fmt.Println(isInterleave("", "", ""))                     // true
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	// TODO: решить
	return false
}
