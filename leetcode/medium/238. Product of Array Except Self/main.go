// https://leetcode.com/problems/product-of-array-except-self/description/
package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))      //[24,12,8,6]
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3})) //[0,0,9,0,0]
}

// Time Complexity: O(N)
// Space Complexity: O(N)
func productExceptSelf(nums []int) []int {
	// Получаем длину входного массива
	// Get the length of the input array
	l := len(nums)

	// Создаем массив для результата
	// Create an array for the result
	res := make([]int, l)

	// Первый элемент результата устанавливаем в 1
	// Set the first element of the result to 1
	res[0] = 1

	// Первый проход: вычисляем произведение всех элементов слева от текущего
	// First pass: calculate product of all elements to the left of current element
	for i := 1; i < l; i++ {
		// res[i] содержит произведение всех чисел до nums[i-1]
		// res[i] contains the product of all numbers up to nums[i-1]
		res[i] = res[i-1] * nums[i-1]
	}

	// Переменная для хранения произведения всех элементов справа от текущего
	// Variable to store the product of all elements to the right of current element
	rp := 1

	// Второй проход: умножаем на произведение всех элементов справа
	// Second pass: multiply by the product of all elements to the right
	for i := l - 1; i >= 0; i-- {
		// Умножаем текущий результат на произведение справа
		// Multiply current result by the product from the right
		res[i] = res[i] * rp

		// Обновляем произведение справа
		// Update the product from the right
		rp = nums[i] * rp
	}

	// Возвращаем результат
	// Return the result
	return res
}
