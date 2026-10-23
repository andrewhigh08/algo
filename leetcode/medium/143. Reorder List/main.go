// https://leetcode.com/problems/reorder-list/description/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	head := newList(1, 2, 3, 4)
	reorderList(head)
	fmt.Println(head) // [1 4 2 3]

	head = newList(1, 2, 3, 4, 5)
	reorderList(head)
	fmt.Println(head) // [1 5 2 4 3]
}

func reorderList(head *ListNode) {
	// TODO: решить
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
