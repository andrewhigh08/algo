// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 1, 2}
	//  indexes    0  1  2  3  4  5  6  7  8  9       len(nums)==10
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4} // - slice (dynamic array)
	//                            l
	//                                        r
	fmt.Println(removeDuplicates(nums1))
	fmt.Println(nums1)
	fmt.Println(removeDuplicates(nums2))
	fmt.Println(nums2)
}

// Time complexity: O(n)
// Space complexity: O(1)
func removeDuplicates(nums []int) int {
	// Если массив пуст, возвращаем 0
	// If the array is empty, return 0
	if len(nums) == 0 {
		return 0
	}

	// l - указатель на позицию, куда нужно записать следующий уникальный элемент
	// l - pointer to the position where the next unique element should be placed
	l := 1

	// r - указатель для сканирования массива
	// r - pointer for scanning the array
	for r := 1; r < len(nums); r++ {
		// Если текущий элемент отличается от предыдущего, он уникален
		// If the current element differs from the previous one, it's unique
		if nums[r] != nums[r-1] {
			// Записываем уникальный элемент в позицию l
			// Write the unique element at position l
			nums[l] = nums[r]

			// Увеличиваем l для следующего уникального элемента
			// Increment l for the next unique element
			l++
		}
		// Если текущий элемент такой же как предыдущий, просто пропускаем его
		// If the current element is the same as the previous one, we skip it
	}

	// l представляет количество уникальных элементов в массиве
	// l represents the number of unique elements in the array
	return l
}
