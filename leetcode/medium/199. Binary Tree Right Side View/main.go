// https://leetcode.com/problems/binary-tree-right-side-view/description/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(rightSideView(newTree(1, 2, 3, null, 5, null, 4)))       // [1 3 4]
	fmt.Println(rightSideView(newTree(1, 2, 3, 4, null, null, null, 5))) // [1 3 4 5]
	fmt.Println(rightSideView(newTree(1, null, 3)))                      // [1 3]
	fmt.Println(rightSideView(newTree()))                                // []
}

func rightSideView(root *TreeNode) []int {
	// TODO: решить
	return nil
}

// --- Вспомогательный код для запуска примеров ---

// TreeNode — узел бинарного дерева, как на LeetCode.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// null обозначает отсутствующий узел в записи дерева по уровням.
const null = math.MinInt

// newTree строит дерево из записи по уровням, как в примерах LeetCode:
// newTree(3, 9, 20, null, null, 15, 7).
func newTree(vals ...int) *TreeNode {
	if len(vals) == 0 || vals[0] == null {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}
	for i := 1; i < len(vals); i += 2 {
		node := queue[0]
		queue = queue[1:]
		if vals[i] != null {
			node.Left = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Left)
		}
		if i+1 < len(vals) && vals[i+1] != null {
			node.Right = &TreeNode{Val: vals[i+1]}
			queue = append(queue, node.Right)
		}
	}
	return root
}
