// https://leetcode.com/problems/merge-two-sorted-lists/description/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(mergeTwoLists(newList(1, 2, 4), newList(1, 3, 4))) // [1 1 2 3 4 4]
	fmt.Println(mergeTwoLists(newList(), newList()))               // []
	fmt.Println(mergeTwoLists(newList(), newList(0)))              // [0]
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
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
