// https://leetcode.com/problems/valid-anagram/
package main

import "fmt"

func main() {
	s, t := "anagram", "nagaram"
	s2, t2 := "rat", "car"

	fmt.Println(isAnagram(s, t))
	fmt.Println(isAnagram(s2, t2))
}

// Time complexity O(n) space complexity O(1)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var freq [26]int

	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
		freq[t[i]-'a']--
	}

	for i := 0; i < len(freq); i++ {
		if freq[i] != 0 {
			return false
		}
	}

	return true
}
