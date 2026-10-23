// https://leetcode.com/problems/kth-largest-element-in-a-stream/description/
package main

import "fmt"

func main() {
	kthLargest := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(kthLargest.Add(3))  // 4
	fmt.Println(kthLargest.Add(5))  // 5
	fmt.Println(kthLargest.Add(10)) // 5
	fmt.Println(kthLargest.Add(9))  // 8
	fmt.Println(kthLargest.Add(4))  // 8

	kthLargest = Constructor(4, []int{7, 7, 7, 7, 8, 3})
	fmt.Println(kthLargest.Add(2))  // 7
	fmt.Println(kthLargest.Add(10)) // 7
	fmt.Println(kthLargest.Add(9))  // 7
	fmt.Println(kthLargest.Add(9))  // 8
}

type KthLargest struct {
	// TODO: решить
}

func Constructor(k int, nums []int) KthLargest {
	return KthLargest{}
}

func (this *KthLargest) Add(val int) int {
	return 0
}
