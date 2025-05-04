// https://leetcode.com/problems/generate-parentheses/description/
package main

import (
	"fmt"
	"strings"
)

// Время (Time): O(C_n * n) или O(4^n / sqrt(n))
// Память (Space): O(C_n * n)
func main() {
	fmt.Println(generateParenthesis(3)) // ["((()))","(()())","(())()","()(())","()()()"]
	fmt.Println(generateParenthesis(1)) // ["()"]
}

// generateParenthesis генерирует все корректные комбинации скобок в количестве n пар.
// Использует подход с бэктрекингом и стеком.
// generateParenthesis generates all valid combinations of n pairs of parentheses.
// It uses a backtracking approach with a stack.
func generateParenthesis(n int) []string {
	var res []string   // Слайс для хранения результатов. / Slice to store the results.
	var stack []string // Слайс, используемый как стек для построения текущей комбинации. / Slice used as a stack to build the current combination.

	// Запускаем вспомогательную рекурсивную функцию backtrack.
	// Начинаем с 0 открывающих и 0 закрывающих скобок.
	// Start the recursive helper function backtrack.
	// Begin with 0 open and 0 close parentheses.
	backtrack(n, 0, 0, &stack, &res)

	return res // Возвращаем полученный слайс с комбинациями. / Return the resulting slice of combinations.
}

// backtrack - вспомогательная рекурсивная функция для построения комбинаций.
// n - общее требуемое количество пар скобок.
// openN - количество использованных открывающих скобок на текущем пути.
// closedN - количество использованных закрывающих скобок на текущем пути.
// stack - указатель на слайс, представляющий текущую комбинацию скобок (стек).
// result - указатель на слайс, куда добавляются готовые корректные комбинации.
//
// backtrack - a recursive helper function to build combinations.
// n - the total number of pairs of parentheses required.
// openN - the number of opening parentheses used on the current path.
// closedN - the number of closing parentheses used on the current path.
// stack - a pointer to the slice representing the current combination of parentheses (stack).
// result - a pointer to the slice where completed valid combinations are added.
func backtrack(n, openN, closedN int, stack *[]string, result *[]string) {
	// Базовый случай: если количество открывающих и закрывающих скобок равно n,
	// значит, мы сформировали полную и корректную комбинацию.
	// Base case: If the number of open and closed parentheses equals n,
	// we have formed a complete and valid combination.
	if openN == n && closedN == n {
		// Объединяем элементы стека в одну строку и добавляем в результаты.
		// Join the elements of the stack into a single string and add to the results.
		*result = append(*result, strings.Join(*stack, ""))
		return // Завершаем ветку рекурсии. / End the recursive branch.
	}

	// Условие для добавления открывающей скобки:
	// Мы можем добавить открывающую скобку, если их количество еще не достигло n.
	// Condition for adding an opening parenthesis:
	// We can add an opening parenthesis if the number used is less than n.
	if openN < n {
		// Добавляем '(' в стек (push).
		// Add '(' to the stack (push).
		*stack = append(*stack, "(")
		// Рекурсивный вызов: увеличиваем счетчик открывающих скобок.
		// Recursive call: increment the count of open parentheses.
		backtrack(n, openN+1, closedN, stack, result)
		// Бэктрекинг: удаляем '(' из стека (pop), чтобы исследовать другие варианты.
		// Backtrack: remove the '(' from the stack (pop) to explore other options.
		*stack = (*stack)[:len(*stack)-1]
	}

	// Условие для добавления закрывающей скобки:
	// Мы можем добавить закрывающую скобку, если их количество строго меньше количества открывающих.
	// Это гарантирует, что в любой момент строка является префиксом корректной последовательности.
	// Condition for adding a closing parenthesis:
	// We can add a closing parenthesis if the number of closed ones is strictly less than the number of open ones.
	// This ensures that at any point, the string is a prefix of a valid sequence.
	if closedN < openN {
		// Добавляем ')' в стек (push).
		// Add ')' to the stack (push).
		*stack = append(*stack, ")")
		// Рекурсивный вызов: увеличиваем счетчик закрывающих скобок.
		// Recursive call: increment the count of closed parentheses.
		backtrack(n, openN, closedN+1, stack, result)
		// Бэктрекинг: удаляем ')' из стека (pop).
		// Backtrack: remove the ')' from the stack (pop).
		*stack = (*stack)[:len(*stack)-1]
	}
}
