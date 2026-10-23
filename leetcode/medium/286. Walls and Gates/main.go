// https://leetcode.com/problems/walls-and-gates/description/ (Premium)
// https://www.lintcode.com/problem/663/
package main

import (
	"fmt"
	"math"
)

func main() {
	rooms := [][]int{
		{INF, -1, 0, INF},
		{INF, INF, INF, -1},
		{INF, -1, INF, -1},
		{0, -1, INF, INF},
	}
	wallsAndGates(rooms)
	fmt.Println(rooms) // [[3 -1 0 1] [2 2 1 -1] [1 -1 2 -1] [0 -1 3 4]]

	rooms = [][]int{{-1}}
	wallsAndGates(rooms)
	fmt.Println(rooms) // [[-1]]
}

func wallsAndGates(rooms [][]int) {
	// TODO: решить
}

// --- Вспомогательный код для запуска примеров ---

// INF — пустая комната: 2³¹ - 1 = 2147483647.
const INF = math.MaxInt32
