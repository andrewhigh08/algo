// https://leetcode.com/problems/evaluate-reverse-polish-notation/description/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))                                             // 9
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))                                            // 6
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})) // 22
}

// Time complexity O(n)
// Space complexity O(n)
func evalRPN(tokens []string) int {
	// operators: A map to quickly check if a token is an operator.
	// Using an empty struct{} as the value is memory-efficient for sets.
	// operators: Карта для быстрой проверки, является ли токен оператором.
	// Использование пустой структуры struct{} в качестве значения экономит память при создании множества.
	operators := map[string]struct{}{
		"*": {},
		"/": {},
		"+": {},
		"-": {},
	}

	// stack: A slice used as a stack to store integer operands.
	// stack: Срез, используемый как стек для хранения целочисленных операндов.
	stack := make([]int, 0)

	// Iterate through each token in the input slice.
	// Итерация по каждому токену во входном срезе.
	for _, t := range tokens {
		// Check if the current token 't' exists as a key in the operators map.
		// Проверяем, существует ли текущий токен 't' как ключ в карте операторов.
		if _, exists := operators[t]; !exists {
			// If the token is not an operator, it must be a number.
			// Convert the token string to an integer.
			// The error is ignored based on the problem's constraints (input is always valid).
			// Если токен не оператор, значит это число.
			// Конвертируем строку токена в целое число.
			// Ошибка игнорируется согласно ограничениям задачи (входные данные всегда валидны).
			num, _ := strconv.Atoi(t)
			// Push the number onto the stack.
			// Помещаем число в стек.
			stack = append(stack, num)
		} else {
			// If the token is an operator.
			// Get the current length of the stack.
			// Если токен - оператор.
			// Получаем текущую длину стека.
			n := len(stack)
			// Pop the top two operands from the stack.
			// Note: operand2 is popped first (the last element).
			// Извлекаем два верхних операнда из стека.
			// Примечание: operand2 извлекается первым (последний элемент).
			operand2 := stack[n-1]
			operand1 := stack[n-2]
			// Remove the top two elements from the stack by slicing.
			// Удаляем два верхних элемента из стека с помощью среза.
			stack = stack[:n-2]

			// Variable to store the result of the operation.
			// Переменная для хранения результата операции.
			var result int

			// Perform the operation based on the operator token.
			// Выполняем операцию в зависимости от токена-оператора.
			switch t {
			case "*":
				result = operand1 * operand2
			case "/":
				// Integer division, truncates towards zero.
				// Целочисленное деление, отбрасывает дробную часть.
				result = operand1 / operand2
			case "+":
				result = operand1 + operand2
			case "-":
				result = operand1 - operand2
			}

			// Push the result back onto the stack.
			// Помещаем результат обратно в стек.
			stack = append(stack, result)
		}
	}

	// After processing all tokens, the final result is the only element left on the stack.
	// Return the top (and only) element of the stack.
	// После обработки всех токенов окончательный результат - это единственный элемент, оставшийся в стеке.
	// Возвращаем верхний (и единственный) элемент стека.
	return stack[0]
}
