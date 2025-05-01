// https://leetcode.com/problems/squares-of-a-sorted-array/
package main

import "fmt"

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10})) // [0,1,9,16,100]
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11})) // [4,9,9,49,121]
}

// Временная сложность: O(n), где n - длина массива
// Time complexity: O(n), where n is the length of the array
//
// Пространственная сложность: O(n) для хранения результата
// Space complexity: O(n) to store the result
func sortedSquares(nums []int) []int {
	// Создаем массив для результата
	// Create an array for the result
	res := make([]int, len(nums))

	// Используем два указателя: l - слева, r - справа
	// Use two pointers: l - left, r - right
	l, r := 0, len(nums)-1

	// c - указатель для заполнения результата с конца
	// c - pointer for filling the result from the end
	c := r

	// Сравниваем квадраты элементов с обоих концов массива
	// Compare squares of elements from both ends of the array
	for l <= r {
		// Вычисляем квадраты текущих элементов
		// Calculate squares of current elements
		sqL := nums[l] * nums[l]
		sqR := nums[r] * nums[r]

		// Помещаем больший квадрат в текущую позицию результата
		// Place the larger square in the current position of the result
		if sqL < sqR {
			res[c] = sqR
			r-- // Сдвигаем правый указатель влево / Move right pointer to the left
		} else {
			res[c] = sqL
			l++ // Сдвигаем левый указатель вправо / Move left pointer to the right
		}
		c-- // Переходим к предыдущей позиции в результате / Move to the previous position in the result
	}

	return res
}
