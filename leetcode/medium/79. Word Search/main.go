// https://leetcode.com/problems/word-search/description/
package main

import "fmt"

func main() {
	fmt.Println(exist(newBoard("ABCE", "SFCS", "ADEE"), "ABCCED")) // true
	fmt.Println(exist(newBoard("ABCE", "SFCS", "ADEE"), "SEE"))    // true
	fmt.Println(exist(newBoard("ABCE", "SFCS", "ADEE"), "ABCB"))   // false
}

func exist(board [][]byte, word string) bool {
	// TODO: решить
	return false
}

// --- Вспомогательный код для запуска примеров ---

// newBoard строит доску из строк: newBoard("ab", "cd") — это [["a","b"],["c","d"]].
func newBoard(rows ...string) [][]byte {
	board := make([][]byte, len(rows))
	for i, row := range rows {
		board[i] = []byte(row)
	}
	return board
}
