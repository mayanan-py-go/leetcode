package main

import "container/list"

/*
用两个栈来实现一个队列，使用n个元素来完成 n 次在队列尾部插入整数(push)和n次在队列头部删除整数(pop)的功能。
队列中的元素为int类型。保证操作合法，即保证pop操作时队列内已有元素。
数据范围： n\le1000n≤1000
要求：存储n个元素的空间复杂度为 O(n)O(n) ，插入与删除的时间复杂度都是 O(1)O(1)
*/

func main() {

}

var stack1 = list.New()
var stack2 = list.New()

func Push(node int) {
	stack1.PushFront(node)
}
func Pop() int {
	if stack2.Len() == 0 {
		for stack1.Len() > 0 {
			stack2.PushFront(stack1.Remove(stack1.Front()))
		}
	}
	if stack2.Len() > 0 {
		return stack2.Remove(stack2.Front()).(int)
	}
	return -1
}

















