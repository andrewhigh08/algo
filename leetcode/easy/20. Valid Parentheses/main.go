// https://leetcode.com/problems/valid-parentheses/description/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("([])"))   // true
}

// Time complexity: O(n)
// Space complexity: O(n)
func isValid(s string) bool {
	// Если длина строки нечетная, то строка не может быть валидной
	// If the string length is odd, it cannot be valid
	if utf8.RuneCountInString(s)%2 != 0 {
		return false
	}

	// Создаем стек для хранения открывающих скобок
	// Create a stack to store opening brackets
	var stack []rune

	// Хеш-таблица для сопоставления закрывающих скобок с открывающими
	// Hash table to map closing brackets to their corresponding opening brackets
	brackets := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	// Проходим по каждому символу в строке
	// Iterate through each character in the string
	for _, c := range s {
		// Если это закрывающая скобка
		// If this is a closing bracket
		if openBracket, exists := brackets[c]; exists {
			// Проверяем, что стек не пуст и последняя открывающая скобка соответствует текущей закрывающей
			// Check that the stack is not empty and the last opening bracket matches the current closing bracket
			if len(stack) == 0 || stack[len(stack)-1] != openBracket {
				return false
			}
			// Удаляем последнюю открывающую скобку из стека
			// Remove the last opening bracket from the stack
			stack = stack[:len(stack)-1]
		} else {
			// Если это открывающая скобка, добавляем ее в стек
			// If it's an opening bracket, add it to the stack
			stack = append(stack, c)
		}
	}

	// Если стек пуст, то все скобки были корректно закрыты
	// If the stack is empty, all brackets were properly closed
	return len(stack) == 0
}
