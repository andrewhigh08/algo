// https://leetcode.com/problems/longest-increasing-path-in-a-matrix/description/
package main

import "fmt"

func main() {
	fmt.Println(longestIncreasingPath([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}})) // 4
	fmt.Println(longestIncreasingPath([][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}})) // 4
	fmt.Println(longestIncreasingPath([][]int{{1}}))                             // 1
}

func longestIncreasingPath(matrix [][]int) int {
	// TODO: решить
	return 0
}
