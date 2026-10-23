// https://leetcode.com/problems/merge-intervals/description/
package main

import "fmt"

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})) // [[1 6] [8 10] [15 18]]
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))                    // [[1 5]]
	fmt.Println(merge([][]int{{4, 7}, {1, 4}}))                    // [[1 7]]
}

func merge(intervals [][]int) [][]int {
	// TODO: решить
	return nil
}
