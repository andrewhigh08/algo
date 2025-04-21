// https://leetcode.com/problems/two-sum/description/
package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{15, 2, 11, 7}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

// Time complexity O(n) space complexity O(n)
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		if idx, found := numMap[num]; found {
			return []int{idx, i}
		}
		numMap[target-num] = i
	}

	return []int{}
}
