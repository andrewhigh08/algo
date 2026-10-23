// https://leetcode.com/problems/set-matrix-zeroes/description/
package main

import "fmt"

func main() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	fmt.Println(matrix) // [[1 0 1] [0 0 0] [1 0 1]]

	matrix = [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix)
	fmt.Println(matrix) // [[0 0 0 0] [0 4 5 0] [0 3 1 0]]
}

func setZeroes(matrix [][]int) {
	// TODO: решить
}
