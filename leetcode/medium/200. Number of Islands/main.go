// https://leetcode.com/problems/number-of-islands/description/
package main

import "fmt"

func main() {
	fmt.Println(numIslands(newGrid("11110", "11010", "11000", "00000"))) // 1
	fmt.Println(numIslands(newGrid("11000", "11000", "00100", "00011"))) // 3
}

func numIslands(grid [][]byte) int {
	// TODO: решить
	return 0
}

// --- Вспомогательный код для запуска примеров ---

// newGrid строит сетку из строк: newGrid("10", "01") — это [["1","0"],["0","1"]].
func newGrid(rows ...string) [][]byte {
	grid := make([][]byte, len(rows))
	for i, row := range rows {
		grid[i] = []byte(row)
	}
	return grid
}
