// https://leetcode.com/problems/redundant-connection/description/
package main

import "fmt"

func main() {
	fmt.Println(findRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}}))                 // [2 3]
	fmt.Println(findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}})) // [1 4]
}

func findRedundantConnection(edges [][]int) []int {
	// TODO: решить
	return nil
}
