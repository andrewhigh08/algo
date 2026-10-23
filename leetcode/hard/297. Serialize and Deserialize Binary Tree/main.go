// https://leetcode.com/problems/serialize-and-deserialize-binary-tree/description/
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	ser := Constructor()
	deser := Constructor()

	data := ser.serialize(newTree(1, 2, 3, null, null, 4, 5))
	fmt.Println(deser.deserialize(data)) // [1 2 3 null null 4 5]

	data = ser.serialize(newTree())
	fmt.Println(deser.deserialize(data)) // []
}

type Codec struct {
	// TODO: решить
}

func Constructor() Codec {
	return Codec{}
}

// serialize кодирует дерево в одну строку.
func (this *Codec) serialize(root *TreeNode) string {
	return ""
}

// deserialize восстанавливает дерево из строки, полученной от serialize.
func (this *Codec) deserialize(data string) *TreeNode {
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
