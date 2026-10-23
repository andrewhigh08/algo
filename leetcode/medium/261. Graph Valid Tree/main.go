// https://leetcode.com/problems/graph-valid-tree/description/ (Premium)
// https://www.lintcode.com/problem/178/
package main

import "fmt"

func main() {
	fmt.Println(validTree(5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}))         // true
	fmt.Println(validTree(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}})) // false
}

func validTree(n int, edges [][]int) bool {
	// TODO: решить
	return false
}
