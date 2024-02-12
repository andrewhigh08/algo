// https://leetcode.com/problems/squares-of-a-sorted-array/
package main

import (
	"fmt"
)

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	l, r := 0, len(nums)-1
	c := r

	for l <= r {
		sqL := nums[l] * nums[l]
		sqR := nums[r] * nums[r]

		if sqL < sqR {
			res[c] = sqR
			r--
		} else {
			res[c] = sqL
			l++
		}
		c--
	}

	return res
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10})) // [0,1,9,16,100]
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11})) // [4,9,9,49,121]
}
