// https://leetcode.com/problems/clone-graph/description/
package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(cloneGraph(newGraph([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}))) // [[2 4] [1 3] [2 4] [1 3]]
	fmt.Println(cloneGraph(newGraph([][]int{{}})))                             // [[]]
	fmt.Println(cloneGraph(newGraph([][]int{})))                               // []
}

func cloneGraph(node *Node) *Node {
	// TODO: решить
	return nil
}

// --- Вспомогательный код для запуска примеров ---

// Node — узел неориентированного графа, как на LeetCode.
type Node struct {
	Val       int
	Neighbors []*Node
}

// newGraph строит граф из списка смежности, как в примерах LeetCode, и возвращает узел с Val = 1.
func newGraph(adjList [][]int) *Node {
	if len(adjList) == 0 {
		return nil
	}
	nodes := make([]*Node, len(adjList)+1)
	for val := 1; val <= len(adjList); val++ {
		nodes[val] = &Node{Val: val}
	}
	for i, neighbors := range adjList {
		for _, val := range neighbors {
			nodes[i+1].Neighbors = append(nodes[i+1].Neighbors, nodes[val])
		}
	}
	return nodes[1]
}

// String выводит граф, достижимый из узла, списком смежности, как на LeetCode: [[2 4] [1 3] [2 4] [1 3]].
func (node *Node) String() string {
	if node == nil {
		return "[]"
	}
	nodes := []*Node{node}
	seen := map[*Node]bool{node: true}
	for i := 0; i < len(nodes); i++ {
		for _, nb := range nodes[i].Neighbors {
			if !seen[nb] {
				seen[nb] = true
				nodes = append(nodes, nb)
			}
		}
	}
	sort.Slice(nodes, func(i, j int) bool { return nodes[i].Val < nodes[j].Val })
	lists := make([]string, len(nodes))
	for i, n := range nodes {
		vals := make([]string, len(n.Neighbors))
		for j, nb := range n.Neighbors {
			vals[j] = strconv.Itoa(nb.Val)
		}
		lists[i] = "[" + strings.Join(vals, " ") + "]"
	}
	return "[" + strings.Join(lists, " ") + "]"
}
