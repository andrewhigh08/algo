// https://leetcode.com/problems/invert-binary-tree/description/
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(invertTree(newTree(4, 2, 7, 1, 3, 6, 9))) // [4 7 2 9 6 3 1]
	fmt.Println(invertTree(newTree(2, 1, 3)))             // [2 3 1]
	fmt.Println(invertTree(newTree()))                    // []
}

func invertTree(root *TreeNode) *TreeNode {
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

// String выводит дерево по уровням, как на LeetCode: [3 9 20 null null 15 7].
func (t *TreeNode) String() string {
	var vals []string
	queue := []*TreeNode{t}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			vals = append(vals, "null")
			continue
		}
		vals = append(vals, strconv.Itoa(node.Val))
		queue = append(queue, node.Left, node.Right)
	}
	for len(vals) > 0 && vals[len(vals)-1] == "null" {
		vals = vals[:len(vals)-1]
	}
	return "[" + strings.Join(vals, " ") + "]"
}
