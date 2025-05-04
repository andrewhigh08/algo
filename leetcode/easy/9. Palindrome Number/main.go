// https://leetcode.com/problems/palindrome-number/
package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}

// Time complexity: O(log₁₀(x))
// Space complexity: O(1)
func isPalindrome(x int) bool {
	// Отрицательные числа не являются палиндромами
	if x < 0 {
		return false
	}

	// Найдем делитель, чтобы получить самую левую цифру
	div := 1
	// Используем (x / div) >= 10 чтобы избежать потенциального переполнения при div *= 10
	// Эта проверка эквивалентна x >= 10 * div для положительных x, но безопаснее
	for x/div >= 10 {
		div *= 10
	}

	// Сравниваем цифры с краев и "сдвигаемся" к центру
	for x > 0 {
		right := x % 10 // Самая правая цифра
		left := x / div // Самая левая цифра

		if left != right {
			return false // Если цифры не совпадают, это не палиндром
		}

		// Убираем проверенные крайние цифры
		x = (x % div) / 10
		// Уменьшаем делитель, так как мы убрали две цифры (слева и справа)
		div /= 100
	}

	// Если цикл завершился, значит все пары цифр совпали
	return true
}
