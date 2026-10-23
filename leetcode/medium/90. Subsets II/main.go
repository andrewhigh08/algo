// https://leetcode.com/problems/subsets-ii/description/
package main

import "fmt"

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2})) // [[] [1] [1 2] [1 2 2] [2] [2 2]] (в любом порядке)
	fmt.Println(subsetsWithDup([]int{0}))       // [[] [0]] (в любом порядке)
}

func subsetsWithDup(nums []int) [][]int {
	// TODO: решить
	return nil
}
