// https://leetcode.com/problems/encode-and-decode-strings/description/
// https://www.lintcode.com/problem/659/
package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	encodedString := encodeString([]string{"Вася", "code", "love", "you"})
	fmt.Println(encodedString)
	fmt.Println(decodeString(encodedString))

	encodedString2 := encodeString([]string{"we", "say", ":", "yes"})
	fmt.Println(encodedString2)
	fmt.Println(decodeString(encodedString2))
}

// Time & space complexity O(n)
func encodeString(inString []string) string {
	var builder strings.Builder

	for _, str := range inString {
		runeCount := utf8.RuneCountInString(str)
		builder.WriteString(strconv.Itoa(runeCount))
		builder.WriteString(string('#'))
		builder.WriteString(str)
	}

	return builder.String()
}

// Time & space complexity O(n)
func decodeString(inString string) []string {
	var res []string
	block, sizeBlock := "", 0

	for _, char := range inString {
		if char != 35 && sizeBlock == 0 { // '#' in UTF8
			block += string(char)
			continue
		} else if char == 35 {
			sizeBlock, _ = strconv.Atoi(block)
			block = ""
			continue
		}
		block += string(char)
		sizeBlock--
		if sizeBlock == 0 {
			res = append(res, block)
			block = ""
		}
	}

	return res
}
