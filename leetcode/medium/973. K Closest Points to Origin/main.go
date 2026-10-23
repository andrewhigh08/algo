// https://leetcode.com/problems/k-closest-points-to-origin/description/
package main

import "fmt"

func main() {
	fmt.Println(kClosest([][]int{{1, 3}, {-2, 2}}, 1))          // [[-2 2]]
	fmt.Println(kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2)) // [[3 3] [-2 4]] (в любом порядке)
}

func kClosest(points [][]int, k int) [][]int {
	// TODO: решить
	return nil
}
