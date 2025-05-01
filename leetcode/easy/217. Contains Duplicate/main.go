// https://leetcode.com/problems/contains-duplicate/
package main

import "fmt"

func main() {
	// Тестовые примеры для функции containsDuplicate
	// Test cases for the containsDuplicate function
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))                   // true, так как 1 повторяется / true, as 1 is repeated
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))                   // false, все элементы уникальны / false, all elements are unique
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})) // true, несколько повторений / true, multiple repetitions
}

// Временная сложность: O(n), где n - длина массива
// Time complexity: O(n), where n is the length of the array
//
// Пространственная сложность: O(n) - в худшем случае все элементы уникальны и будут добавлены в карту
// Space complexity: O(n) - worst case when all elements are unique and will be added to the map
func containsDuplicate(nums []int) bool {
	// Создаём карту для отслеживания уже просмотренных чисел
	// Create a map to track numbers we've already seen
	seen := make(map[int]bool, len(nums))

	// Перебираем все числа в массиве
	// Iterate through all numbers in the array
	for _, num := range nums {
		// Проверяем, встречалось ли это число ранее
		// Check if we have seen this number before
		if _, exists := seen[num]; exists {
			return true // Нашли дубликат / Found a duplicate
		}

		// Добавляем число в карту просмотренных
		// Add the number to the map of seen numbers
		seen[num] = true
	}

	// Дубликатов не найдено
	// No duplicates found
	return false
}
