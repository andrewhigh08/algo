// https://leetcode.com/problems/detect-squares/description/
package main

import "fmt"

func main() {
	detectSquares := Constructor()
	detectSquares.Add([]int{3, 10})
	detectSquares.Add([]int{11, 2})
	detectSquares.Add([]int{3, 2})
	fmt.Println(detectSquares.Count([]int{11, 10})) // 1
	fmt.Println(detectSquares.Count([]int{14, 8}))  // 0
	detectSquares.Add([]int{11, 2})                 // добавлять дублирующиеся точки разрешено
	fmt.Println(detectSquares.Count([]int{11, 10})) // 2
}

type DetectSquares struct {
	// TODO: решить
}

func Constructor() DetectSquares {
	return DetectSquares{}
}

func (this *DetectSquares) Add(point []int) {
}

func (this *DetectSquares) Count(point []int) int {
	return 0
}
