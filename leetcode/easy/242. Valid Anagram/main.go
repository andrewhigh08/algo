// https://leetcode.com/problems/valid-anagram/
package main

import "fmt"

func main() {
	// Тестовый пример 1: "anagram" и "nagaram" - это анаграммы
	// Test case 1: "anagram" and "nagaram" are anagrams
	s, t := "anagram", "nagaram"

	// Тестовый пример 2: "rat" и "car" - не анаграммы
	// Test case 2: "rat" and "car" are not anagrams
	s2, t2 := "rat", "car"

	// Вывод результатов проверки
	// Output the check results
	fmt.Println(isAnagram(s, t))   // true
	fmt.Println(isAnagram(s2, t2)) // false
}

// Временная сложность: O(n), где n - длина строки
// Time complexity: O(n), where n is the length of the string
//
// Пространственная сложность: O(1), так как используется массив фиксированного размера (26 букв)
// Space complexity: O(1), as we use a fixed-size array (26 letters)
func isAnagram(s string, t string) bool {
	// Если длины строк различаются, они не могут быть анаграммами
	// If string lengths differ, they cannot be anagrams
	if len(s) != len(t) {
		return false
	}

	// Массив для подсчёта частоты появления каждой буквы
	// Array to count frequency of each letter
	var freq [26]int

	// Проходим по обеим строкам одновременно
	// Iterate through both strings simultaneously
	for i := 0; i < len(s); i++ {
		// Увеличиваем счётчик для буквы из первой строки
		// Increment counter for letter from the first string
		freq[s[i]-'a']++

		// Уменьшаем счётчик для буквы из второй строки
		// Decrement counter for letter from the second string
		freq[t[i]-'a']--
	}

	// Проверяем, что все счётчики равны нулю
	// Check that all counters are equal to zero
	for i := 0; i < len(freq); i++ {
		// Если счётчик не равен нулю, строки не являются анаграммами
		// If counter is not zero, strings are not anagrams
		if freq[i] != 0 {
			return false
		}
	}

	// Если все счётчики равны нулю, строки являются анаграммами
	// If all counters are zero, strings are anagrams
	return true
}
