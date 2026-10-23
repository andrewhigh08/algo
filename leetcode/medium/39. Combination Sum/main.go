// https://leetcode.com/problems/combination-sum/description/
package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7)) // [[2 2 3] [7]] (в любом порядке)
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))    // [[2 2 2 2] [2 3 3] [3 5]] (в любом порядке)
	fmt.Println(combinationSum([]int{2}, 1))          // []
}

func combinationSum(candidates []int, target int) [][]int {
	// TODO: решить
	return nil
}
