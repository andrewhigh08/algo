// https://leetcode.com/problems/surrounded-regions/description/
package main

import "fmt"

func main() {
	board := newBoard("XXXX", "XOOX", "XXOX", "XOXX")
	solve(board)
	fmt.Printf("%q\n", board) // ["XXXX" "XXXX" "XXXX" "XOXX"]

	board = newBoard("X")
	solve(board)
	fmt.Printf("%q\n", board) // ["X"]
}

func solve(board [][]byte) {
	// TODO: решить
}

// --- Вспомогательный код для запуска примеров ---

// newBoard строит доску из строк: newBoard("XO", "OX") — это [["X","O"],["O","X"]].
func newBoard(rows ...string) [][]byte {
	board := make([][]byte, len(rows))
	for i, row := range rows {
		board[i] = []byte(row)
	}
	return board
}
