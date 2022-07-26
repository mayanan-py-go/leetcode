package main

/*
用两个栈来实现一个队列，使用n个元素来完成 n 次在队列尾部插入整数(push)和n次在队列头部删除整数(pop)的功能。
队列中的元素为int类型。保证操作合法，即保证pop操作时队列内已有元素。
数据范围： n\le1000n≤1000
要求：存储n个元素的空间复杂度为 O(n)O(n) ，插入与删除的时间复杂度都是 O(1)O(1)
*/

func main() {

}

var stack1 []int
var stack2 []int

func Push(node int) {
	stack1 = append(stack1, node)
}
func Pop() int {
	if len(stack2) == 0 {
		length := len(stack1)
		for i := 0; i < length; i++ {
			tmp := stack1[len(stack1)-1]
			stack2 = append(stack2, tmp)
			stack1 = stack1[:len(stack1)-1]
		}
	}
	if len(stack2) > 0 {
		tmp := stack2[len(stack2)-1]
		stack2 = stack2[:len(stack2)-1]
		return tmp
	}
	return -1
}
















