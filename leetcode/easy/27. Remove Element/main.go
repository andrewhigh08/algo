// https://leetcode.com/problems/remove-element/description/
package main

import "fmt"

func main() {
	// Тестовый пример 1: [3, 2, 2, 3] с удалением элемента 3
	// Test case 1: [3, 2, 2, 3] with element 3 to be removed
	n1 := []int{3, 2, 2, 3}
	fmt.Println(removeElement(n1, 3), "nums = ", n1)

	// Тестовый пример 2: [0, 1, 2, 2, 3, 0, 4, 2] с удалением элемента 2
	// Test case 2: [0, 1, 2, 2, 3, 0, 4, 2] with element 2 to be removed
	n2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(n2, 2), "nums = ", n2)
}

// Time complexity: O(n)
// Space complexity: O(1)
func removeElement(nums []int, val int) int {
	// l - указатель для позиции, куда нужно поместить следующий элемент, не равный val
	// l - pointer for the position where to put the next element not equal to val
	l := 0

	// Проходим по всему массиву
	// Iterate through the entire array
	for i := 0; i < len(nums); i++ {
		// Если текущий элемент не равен значению, которое нужно удалить
		// If the current element is not equal to the value to be removed
		if nums[i] != val {
			// Помещаем его в позицию l и увеличиваем l
			// Place it at position l and increment l
			nums[l] = nums[i]
			l++
		}
		// Если элемент равен val, просто пропускаем его (не увеличиваем l)
		// If the element equals val, we just skip it (don't increment l)
	}

	// Возвращаем количество элементов, которые не равны val
	// Return the number of elements that are not equal to val
	return l
}
