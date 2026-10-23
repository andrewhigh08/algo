// https://leetcode.com/problems/add-two-numbers/description/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(addTwoNumbers(newList(2, 4, 3), newList(5, 6, 4)))                // [7 0 8]
	fmt.Println(addTwoNumbers(newList(0), newList(0)))                            // [0]
	fmt.Println(addTwoNumbers(newList(9, 9, 9, 9, 9, 9, 9), newList(9, 9, 9, 9))) // [8 9 9 9 0 0 0 1]
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// TODO: решить
	return nil
}

// --- Вспомогательный код для запуска примеров ---

// ListNode — узел односвязного списка, как на LeetCode.
type ListNode struct {
	Val  int
	Next *ListNode
}

// newList строит список из значений: newList(1, 2, 3) — это 1 -> 2 -> 3.
func newList(vals ...int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, v := range vals {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}
	return dummy.Next
}

// String выводит список в виде [1 2 3].
func (l *ListNode) String() string {
	var vals []string
	for node := l; node != nil; node = node.Next {
		vals = append(vals, strconv.Itoa(node.Val))
	}
	return "[" + strings.Join(vals, " ") + "]"
}
