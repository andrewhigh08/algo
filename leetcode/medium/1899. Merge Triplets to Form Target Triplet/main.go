// https://leetcode.com/problems/merge-triplets-to-form-target-triplet/description/
package main

import "fmt"

func main() {
	fmt.Println(mergeTriplets([][]int{{2, 5, 3}, {1, 8, 4}, {1, 7, 5}}, []int{2, 7, 5}))            // true
	fmt.Println(mergeTriplets([][]int{{3, 4, 5}, {4, 5, 6}}, []int{3, 2, 5}))                       // false
	fmt.Println(mergeTriplets([][]int{{2, 5, 3}, {2, 3, 4}, {1, 2, 5}, {5, 2, 3}}, []int{5, 5, 5})) // true
}

func mergeTriplets(triplets [][]int, target []int) bool {
	// TODO: решить
	return false
}
