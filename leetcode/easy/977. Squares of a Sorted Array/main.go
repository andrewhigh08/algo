package main

import (
	"fmt"
)

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	l, r := 0, len(nums)-1
	c := r

	for l <= r {
		if nums[l]*nums[l] < nums[r]*nums[r] {
			res[c] = nums[r] * nums[r]
			r--
		} else {
			res[c] = nums[l] * nums[l]
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
