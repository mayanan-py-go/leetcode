package main

import (
	"container/list"
	"fmt"
)

/*
定义栈的数据结构，请在该类型中实现一个能够得到栈中所含最小元素的 min 函数，
输入操作时保证 pop、top 和 min 函数操作时，栈中一定有元素。
此栈包含的方法有：
push(value):将value压入栈中
pop():弹出栈顶元素
top():获取栈顶元素
min():获取栈中最小元素
数据范围：操作数量满足 0 \le n \le 300 \0≤n≤300  ，输入的元素满足 |val| \le 10000 \∣val∣≤10000
进阶：栈的各个操作的时间复杂度是 O(1)\O(1)  ，空间复杂度是 O(n)\O(n)
*/

func main() {
	Push(2)
	Push(3)
	Push(6)
	Push(2)
	Push(4)
	Push(5)
	Pop()
	Pop()
	Pop()
	fmt.Println(Min())
}
type MinStack struct {
	Stack *list.List
	Min *list.List
}
var minStack = MinStack{Stack: list.New(), Min: list.New()}
func Push(node int) {
	minStack.Stack.PushBack(node)
	if minStack.Min.Len() == 0 || node <= minStack.Min.Back().Value.(int) {
		minStack.Min.PushBack(node)
	}
}
func Pop() {
	if minStack.Stack.Len() > 0 {
		pop := minStack.Stack.Remove(minStack.Stack.Back()).(int)
		if pop == Min() {
			minStack.Min.Remove(minStack.Min.Back())
		}
	}
}
func Top() int {
	if minStack.Stack.Len() > 0 {
		return minStack.Stack.Back().Value.(int)
	}
	return 0
}
func Min() int {
	if minStack.Min.Len() > 0 {
		return minStack.Min.Back().Value.(int)
	}
	return 0
}















