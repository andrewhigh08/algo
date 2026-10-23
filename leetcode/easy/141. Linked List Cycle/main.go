// https://leetcode.com/problems/linked-list-cycle/description/
package main

import "fmt"

func main() {
	fmt.Println(hasCycle(newCycleList([]int{3, 2, 0, -4}, 1))) // true
	fmt.Println(hasCycle(newCycleList([]int{1, 2}, 0)))        // true
	fmt.Println(hasCycle(newCycleList([]int{1}, -1)))          // false
}

func hasCycle(head *ListNode) bool {
	// TODO: решить
	return false
}

// --- Вспомогательный код для запуска примеров ---

// ListNode — узел односвязного списка, как на LeetCode.
type ListNode struct {
	Val  int
	Next *ListNode
}

// newCycleList строит список из vals и замыкает хвост на узел с индексом pos (pos = -1 — без цикла).
func newCycleList(vals []int, pos int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	var cycleStart *ListNode
	for i, v := range vals {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
		if i == pos {
			cycleStart = tail
		}
	}
	tail.Next = cycleStart
	return dummy.Next
}
