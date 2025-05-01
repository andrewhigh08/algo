// https://leetcode.com/problems/min-stack/description/
package main

import "fmt"

func main() {
	// Пример из описания задачи
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println("getMin():", minStack.GetMin()) // возвращает -3
	minStack.Pop()
	fmt.Println("top():", minStack.Top())       // возвращает 0
	fmt.Println("getMin():", minStack.GetMin()) // возвращает -2
}

type MinStack struct {
}

func Constructor() MinStack {

}

func (this *MinStack) Push(val int) {

}

func (this *MinStack) Pop() {

}

func (this *MinStack) Top() int {

}

func (this *MinStack) GetMin() int {

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
