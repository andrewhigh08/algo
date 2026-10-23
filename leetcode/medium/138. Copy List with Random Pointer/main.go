// https://leetcode.com/problems/copy-list-with-random-pointer/description/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(copyRandomList(newRandomList([][]int{{7, null}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}))) // [[7 null] [13 0] [11 4] [10 2] [1 0]]
	fmt.Println(copyRandomList(newRandomList([][]int{{1, 1}, {2, 1}})))                               // [[1 1] [2 1]]
	fmt.Println(copyRandomList(newRandomList([][]int{{3, null}, {3, 0}, {3, null}})))                 // [[3 null] [3 0] [3 null]]
}

func copyRandomList(head *Node) *Node {
	// TODO: решить
	return nil
}

// --- Вспомогательный код для запуска примеров ---

// Node — узел списка с дополнительным указателем Random, как на LeetCode.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// null в паре [val, random_index] означает, что Random ни на что не указывает.
const null = -1

// newRandomList строит список из пар [val, random_index], как в примерах LeetCode.
func newRandomList(pairs [][]int) *Node {
	nodes := make([]*Node, len(pairs))
	for i, p := range pairs {
		nodes[i] = &Node{Val: p[0]}
	}
	for i, p := range pairs {
		if i+1 < len(nodes) {
			nodes[i].Next = nodes[i+1]
		}
		if p[1] != null {
			nodes[i].Random = nodes[p[1]]
		}
	}
	if len(nodes) == 0 {
		return nil
	}
	return nodes[0]
}

// String выводит список парами [val random_index], как на LeetCode: [[7 null] [13 0]].
func (l *Node) String() string {
	index := make(map[*Node]int)
	for node, i := l, 0; node != nil; node, i = node.Next, i+1 {
		index[node] = i
	}
	var pairs []string
	for node := l; node != nil; node = node.Next {
		random := "null"
		if node.Random != nil {
			random = strconv.Itoa(index[node.Random])
		}
		pairs = append(pairs, "["+strconv.Itoa(node.Val)+" "+random+"]")
	}
	return "[" + strings.Join(pairs, " ") + "]"
}
