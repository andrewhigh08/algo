// https://leetcode.com/problems/insert-interval/description/
package main

import "fmt"

func main() {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))                            // [[1 5] [6 9]]
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8})) // [[1 2] [3 10] [12 16]]
}

func insert(intervals [][]int, newInterval []int) [][]int {
	// TODO: решить
	return nil
}
