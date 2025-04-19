// https://leetcode.com/problems/contains-duplicate/
package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}

// Time complexity O(n) space complexity O(n)
func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool, len(nums))

	for _, num := range nums {
		if _, exists := seen[num]; exists {
			return true
		}

		seen[num] = true
	}

	return false
}
