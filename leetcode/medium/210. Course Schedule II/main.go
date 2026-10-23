// https://leetcode.com/problems/course-schedule-ii/description/
package main

import "fmt"

func main() {
	fmt.Println(findOrder(2, [][]int{{1, 0}}))                         // [0 1]
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})) // [0 2 1 3] (или [0 1 2 3])
	fmt.Println(findOrder(1, [][]int{}))                               // [0]
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	// TODO: решить
	return nil
}
