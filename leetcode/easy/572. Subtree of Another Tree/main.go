// https://leetcode.com/problems/subtree-of-another-tree/description/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isSubtree(newTree(3, 4, 5, 1, 2), newTree(4, 1, 2)))                            // true
	fmt.Println(isSubtree(newTree(3, 4, 5, 1, 2, null, null, null, null, 0), newTree(4, 1, 2))) // false
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// TODO: решить
	return false
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
