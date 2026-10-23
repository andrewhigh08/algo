// https://leetcode.com/problems/rotting-oranges/description/
package main

import "fmt"

func main() {
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}})) // 4
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}})) // -1
	fmt.Println(orangesRotting([][]int{{0, 2}}))                          // 0
}

func orangesRotting(grid [][]int) int {
	// TODO: решить
	return 0
}
