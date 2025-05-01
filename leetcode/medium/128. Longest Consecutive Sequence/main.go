package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))         // 4
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})) // 9
	fmt.Println(longestConsecutive([]int{1, 0, 1, 2}))                   // 3
}

// Time Complexity: O(N)
// Space Complexity: O(N)
func longestConsecutive(nums []int) int {
	// Создаем множество из элементов массива для быстрого поиска O(1)
	// Create a set from array elements for fast O(1) lookup
	numsSet := make(map[int]struct{})
	for _, n := range nums {
		numsSet[n] = struct{}{}
	}

	// Здесь храним максимальную длину последовательности
	// Store the maximum sequence length here
	maxSeq := 0

	// Перебираем все числа из множества
	// Iterate through all numbers in the set
	for n := range numsSet {
		// Проверяем, есть ли в множестве число n-1
		// Если есть, то n не является началом последовательности,
		// и мы сразу переходим к следующему числу
		//
		// Check if n-1 exists in the set
		// If it does, then n is not the beginning of a sequence,
		// so we immediately move to the next number
		if _, exist := numsSet[n-1]; !exist {
			// Если n-1 отсутствует, то n - начало последовательности
			// Начинаем проверять наличие последующих чисел
			//
			// If n-1 is missing, then n is the beginning of a sequence
			// Start checking for subsequent numbers
			seqLen := 1 // Начальная длина последовательности | Initial sequence length

			for {
				// Проверяем, существует ли следующее число в последовательности
				// Check if the next number in the sequence exists
				if _, exist = numsSet[n+seqLen]; exist {
					seqLen++ // Увеличиваем длину последовательности | Increase sequence length
					continue
				}

				// Когда последовательность закончилась, сравниваем её длину с максимальной найденной
				// When the sequence ends, compare its length with the maximum found so far
				maxSeq = max(seqLen, maxSeq)
				break
			}
		}
	}

	return maxSeq
}

// Вспомогательная функция для нахождения максимума из двух чисел
// Helper function to find the maximum of two numbers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
