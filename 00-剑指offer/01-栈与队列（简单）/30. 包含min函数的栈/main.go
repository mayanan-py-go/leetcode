package main

import (
	"container/list"
	"fmt"
)

/*
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中
调用 min、push 及 pop 的时间复杂度都是 O(1)。
*/

func main() {
	stack := Constructor()
	stack.Push(85)
	stack.Push(3)
	stack.Push(6)
	stack.Push(3)
	fmt.Println(stack.Min()) // 输出3
	stack.Pop()
	fmt.Println(stack.Min()) // 输出3
	stack.Pop()
}

type MinStack struct {
	stack, min *list.List
}

func Constructor() MinStack {
	return MinStack{
		stack: list.New(),
		min:   list.New(),
	}
}
func (m *MinStack) Push(value int) {
	m.stack.PushBack(value)
	if m.min.Len() == 0 || value <= m.min.Back().Value.(int) {
		m.min.PushBack(value)
	}
}
func (m *MinStack) Pop() {
	if m.stack.Len() > 0 {
		if v := m.stack.Remove(m.stack.Back()); v == m.min.Back().Value {
			m.min.Remove(m.min.Back())
		}
	}
}
func (m *MinStack) Top() int {
	if m.stack.Len() > 0 {
		return m.stack.Back().Value.(int)
	}
	return 0
}
func (m *MinStack) Min() int {
	if m.min.Len() > 0 {
		return m.min.Back().Value.(int)
	}
	return 0
}
