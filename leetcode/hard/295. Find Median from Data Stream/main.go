// https://leetcode.com/problems/find-median-from-data-stream/description/
package main

import "fmt"

func main() {
	medianFinder := Constructor()
	medianFinder.AddNum(1)                          // arr = [1]
	medianFinder.AddNum(2)                          // arr = [1, 2]
	fmt.Printf("%.5f\n", medianFinder.FindMedian()) // 1.50000
	medianFinder.AddNum(3)                          // arr = [1, 2, 3]
	fmt.Printf("%.5f\n", medianFinder.FindMedian()) // 2.00000
}

type MedianFinder struct {
	// TODO: решить
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
}

func (this *MedianFinder) FindMedian() float64 {
	return 0
}
