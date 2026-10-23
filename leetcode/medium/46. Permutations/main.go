// https://leetcode.com/problems/permutations/description/
package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3})) // [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]] (в любом порядке)
	fmt.Println(permute([]int{0, 1}))    // [[0 1] [1 0]] (в любом порядке)
	fmt.Println(permute([]int{1}))       // [[1]]
}

func permute(nums []int) [][]int {
	// TODO: решить
	return nil
}
