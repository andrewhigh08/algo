// https://leetcode.com/problems/top-k-frequent-elements/
package main

import "fmt"

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)) // [1, 2]
	//fmt.Println(topKFrequent([]int{1}, 1))                // [1]
	fmt.Println(topKFrequent([]int{1, 2}, 2)) // [1, 2]
}

// Time Complexity: O(N)
// Space Complexity: O(N)
func topKFrequent(nums []int, k int) (res []int) {
	// Создаем карту для подсчета частоты каждого числа
	// Create a map to count frequency of each number
	cntMp := make(map[int]int)

	// Подсчитываем частоту каждого числа в массиве
	// Count frequency of each number in the array
	for _, num := range nums {
		if cnt, exist := cntMp[num]; exist {
			cntMp[num] = cnt + 1 // Увеличиваем счетчик, если число уже встречалось
			// Increment counter if number already exists
		} else {
			cntMp[num] = 1 // Инициализируем счетчик, если число встречается впервые
			// Initialize counter if number appears for the first time
		}
	}

	// Создаем массив "корзин", где индекс - частота, а значение - список чисел с такой частотой
	// Create a "bucket" array where index is frequency and value is list of numbers with that frequency
	cntSl := make([][]int, len(nums)+1)

	// Распределяем числа по корзинам в соответствии с их частотой
	// Distribute numbers to buckets according to their frequency
	for num, cnt := range cntMp {
		cntSl[cnt] = append(cntSl[cnt], num) // Добавляем число в соответствующую корзину
		// Add number to corresponding bucket
	}

	// Проходим по корзинам от наибольшей частоты к наименьшей
	// Iterate through buckets from highest frequency to lowest
	for i := len(cntSl) - 1; i > 0; i-- {
		res = append(res, cntSl[i]...) // Добавляем все числа из текущей корзины в результат
		// Add all numbers from current bucket to result
		if len(res) == k {
			return // Возвращаем результат, когда нашли k наиболее частых элементов
			// Return result when we found k most frequent elements
		}
	}

	return // Возвращаем результат (на случай, если k больше количества уникальных чисел)
	// Return result (in case k is larger than number of unique numbers)
}
