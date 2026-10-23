// https://leetcode.com/problems/subsets/description/
package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3})) // [[] [1] [2] [1 2] [3] [1 3] [2 3] [1 2 3]] (в любом порядке)
	fmt.Println(subsets([]int{0}))       // [[] [0]] (в любом порядке)
}

func subsets(nums []int) [][]int {
	// TODO: решить
	return nil
}
