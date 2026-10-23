// https://leetcode.com/problems/word-search-ii/description/
package main

import "fmt"

func main() {
	fmt.Printf("%q\n", findWords(newBoard("oaan", "etae", "ihkr", "iflv"), []string{"oath", "pea", "eat", "rain"})) // ["eat" "oath"] (в любом порядке)
	fmt.Printf("%q\n", findWords(newBoard("ab", "cd"), []string{"abcb"}))                                           // []
}

func findWords(board [][]byte, words []string) []string {
	// TODO: решить
	return nil
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
