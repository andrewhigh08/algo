// https://leetcode.com/problems/pacific-atlantic-water-flow/description/
package main

import "fmt"

func main() {
	fmt.Println(pacificAtlantic([][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	})) // [[0 4] [1 3] [1 4] [2 2] [3 0] [3 1] [4 0]] (в любом порядке)
	fmt.Println(pacificAtlantic([][]int{{1}})) // [[0 0]]
}

func pacificAtlantic(heights [][]int) [][]int {
	// TODO: решить
	return nil
}
