// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/
package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	root := newTree(6, 2, 8, 0, 4, 7, 9, null, null, 3, 5)
	fmt.Println(valOf(lowestCommonAncestor(root, find(root, 2), find(root, 8)))) // 6
	fmt.Println(valOf(lowestCommonAncestor(root, find(root, 2), find(root, 4)))) // 2

	root = newTree(2, 1)
	fmt.Println(valOf(lowestCommonAncestor(root, find(root, 2), find(root, 1)))) // 2
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
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

// find ищет в BST узел со значением val (p и q в примерах заданы значениями).
func find(root *TreeNode, val int) *TreeNode {
	for root != nil && root.Val != val {
		if val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}

// valOf возвращает значение узла, как его выводит LeetCode, или "null" для nil.
func valOf(node *TreeNode) string {
	if node == nil {
		return "null"
	}
	return strconv.Itoa(node.Val)
}
