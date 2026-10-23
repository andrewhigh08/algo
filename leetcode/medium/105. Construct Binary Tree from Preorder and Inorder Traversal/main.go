// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})) // [3 9 20 null null 15 7]
	fmt.Println(buildTree([]int{-1}, []int{-1}))                           // [-1]
}

func buildTree(preorder []int, inorder []int) *TreeNode {
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
