// https://leetcode.com/problems/valid-sudoku/description/
package main

import (
	"fmt"
)

func main() {
	var board1 = [][]byte{ // true
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	var board2 = [][]byte{ // false
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println(isValidSudoku(board1))
	fmt.Println(isValidSudoku(board2))
}

// Time & space complexity O(1)
func isValidSudoku(board [][]byte) bool {
	// Используем map[byte]struct{} для эмуляции сетов (byte, так как цифры - это byte)
	rows := make([]map[byte]struct{}, 9)    // Слайс сетов для строк
	cols := make([]map[byte]struct{}, 9)    // Слайс сетов для колонок
	squares := make([]map[byte]struct{}, 9) // Слайс сетов для 3x3 квадратов

	// Инициализируем все сеты
	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]struct{})
		cols[i] = make(map[byte]struct{})
		squares[i] = make(map[byte]struct{})
	}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			digit := board[r][c]

			if digit == '.' {
				continue // если '.', пропускаем
			}

			// Проверяем строку
			if _, exists := rows[r][digit]; exists {
				return false // Уже есть в строке
			}

			// Проверяем колонку
			if _, exists := cols[c][digit]; exists {
				return false // Уже есть в колонке
			}

			// Проверяем 3x3 квадрат
			// Вычисляем индекс квадрата (от 0 до 8)
			squareIndex := (r/3)*3 + (c / 3)
			if _, exists := squares[squareIndex][digit]; exists {
				return false // Уже есть в квадрате
			}

			rows[r][digit] = struct{}{}
			cols[c][digit] = struct{}{}
			squares[squareIndex][digit] = struct{}{}
		}
	}

	// Если прошли всю доску без нарушений
	return true
}
