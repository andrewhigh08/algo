// https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/description/ (Premium)
// https://www.lintcode.com/problem/3651/
package main

import "fmt"

func main() {
	fmt.Println(countComponents(5, [][]int{{0, 1}, {1, 2}, {3, 4}}))         // 2
	fmt.Println(countComponents(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}})) // 1
}

func countComponents(n int, edges [][]int) int {
	// TODO: решить
	return 0
}
