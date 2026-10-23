// https://leetcode.com/problems/cheapest-flights-within-k-stops/description/
package main

import "fmt"

func main() {
	fmt.Println(findCheapestPrice(4, [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}, 0, 3, 1)) // 700
	fmt.Println(findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1))                           // 200
	fmt.Println(findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 0))                           // 500
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	// TODO: решить
	return 0
}
