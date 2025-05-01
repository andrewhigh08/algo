// https://leetcode.com/problems/two-sum/description/
package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{15, 2, 11, 7}, 9)) // Ожидаем [1, 3] (2+7=9)
	fmt.Println(twoSum([]int{3, 2, 4}, 6))      // Ожидаем [1, 2] (2+4=6)
	fmt.Println(twoSum([]int{3, 3}, 6))         // Ожидаем [0, 1] (3+3=6)
}

// Time complexity O(n)
// Space complexity O(n)
func twoSum(nums []int, target int) []int {
	// Создаем map для хранения разницы между целевым значением и текущим числом
	// Create a map to store the difference between target and current number
	numMap := make(map[int]int)

	// Перебираем все элементы массива
	// Iterate through all elements in the array
	for i, num := range nums {
		// Проверяем, есть ли текущее число в map
		// Если есть, значит мы уже встречали число, которое в сумме с текущим дает target
		// Check if the current number exists in the map
		// If it does, we've already seen a number that sums with current to target
		if idx, found := numMap[num]; found {
			return []int{idx, i}
		}

		// Сохраняем разницу между target и текущим числом как ключ,
		// а индекс текущего числа как значение
		// Store the difference between target and current number as key,
		// and the index of current number as value
		numMap[target-num] = i
	}

	// Если не найдено двух чисел, возвращаем пустой массив
	// If no two numbers are found, return an empty array
	return []int{}
}
