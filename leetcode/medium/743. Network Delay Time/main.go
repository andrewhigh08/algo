// https://leetcode.com/problems/network-delay-time/description/
package main

import "fmt"

func main() {
	fmt.Println(networkDelayTime([][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2)) // 2
	fmt.Println(networkDelayTime([][]int{{1, 2, 1}}, 2, 1))                       // 1
	fmt.Println(networkDelayTime([][]int{{1, 2, 1}}, 2, 2))                       // -1
}

func networkDelayTime(times [][]int, n int, k int) int {
	// TODO: решить
	return 0
}
