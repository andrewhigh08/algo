// https://leetcode.com/problems/non-overlapping-intervals/description/
package main

import "fmt"

func main() {
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})) // 1
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}}))         // 2
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}}))                 // 0
}

func eraseOverlapIntervals(intervals [][]int) int {
	// TODO: решить
	return 0
}
