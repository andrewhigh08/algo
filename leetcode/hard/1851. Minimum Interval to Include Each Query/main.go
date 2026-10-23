// https://leetcode.com/problems/minimum-interval-to-include-each-query/description/
package main

import "fmt"

func main() {
	fmt.Println(minInterval([][]int{{1, 4}, {2, 4}, {3, 6}, {4, 4}}, []int{2, 3, 4, 5}))     // [3 3 1 4]
	fmt.Println(minInterval([][]int{{2, 3}, {2, 5}, {1, 8}, {20, 25}}, []int{2, 19, 5, 22})) // [2 -1 4 6]
}

func minInterval(intervals [][]int, queries []int) []int {
	// TODO: решить
	return nil
}
