// https://leetcode.com/problems/product-of-array-except-self/description/
package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))      //[24,12,8,6]
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3})) //[0,0,9,0,0]
}

// Time Complexity: O(N)
// Space Complexity: O(N)
func productExceptSelf(nums []int) []int {
	l := len(nums)
	res := make([]int, l)
	res[0] = 1
	for i := 1; i < l; i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	rp := 1
	for i := l - 1; i >= 0; i-- {
		res[i] = res[i] * rp
		rp = nums[i] * rp
	}

	return res
}

// Time Complexity: O(N)
// Space Complexity: O(N)
func productExceptSelf_(nums []int) []int {
	l := len(nums)
	res := make([]int, l)
	left := make([]int, l)
	right := make([]int, l)
	left[0], right[l-1] = 1, 1

	for l, r := 1, l-2; l < len(nums); l, r = l+1, r-1 {
		left[l] = left[l-1] * nums[l-1]
		right[r] = right[r+1] * nums[r+1]
	}

	for i := 0; i < l; i++ {
		res[i] = left[i] * right[i]
	}

	return res
}
