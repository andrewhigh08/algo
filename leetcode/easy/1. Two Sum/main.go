// https://leetcode.com/problems/two-sum/description/
package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{15, 2, 11, 7}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		if j, ok := m[v]; ok {
			return []int{j, i}
		}
		m[target-v] = i
	}

	return []int{}
}
