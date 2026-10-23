// https://leetcode.com/problems/merge-k-sorted-lists/description/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(mergeKLists([]*ListNode{newList(1, 4, 5), newList(1, 3, 4), newList(2, 6)})) // [1 1 2 3 4 4 5 6]
	fmt.Println(mergeKLists([]*ListNode{}))                                                  // []
	fmt.Println(mergeKLists([]*ListNode{newList()}))                                         // []
}

func mergeKLists(lists []*ListNode) *ListNode {
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
