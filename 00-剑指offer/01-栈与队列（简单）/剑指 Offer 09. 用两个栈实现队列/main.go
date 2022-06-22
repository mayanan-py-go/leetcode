package main

import (
	"container/list"
	"fmt"
)

/*
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead
分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead操作返回 -1 )
https://leetcode.cn/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/
*/

func main() {
	cq := Constructor()
	cq.AppendTail(11)
	cq.AppendTail(22)
	cq.AppendTail(33)
	fmt.Println(cq.DeleteHead())
	fmt.Println(cq.DeleteHead())
	fmt.Println(cq.DeleteHead())
	fmt.Println(cq.DeleteHead())
}

type CQueue struct {
	stack1, stack2 *list.List
}

func Constructor() CQueue {
	return CQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

func (c *CQueue) AppendTail(value int) {
	// 栈1添加元素
	c.stack1.PushBack(value)
}
func (c *CQueue) DeleteHead() int {
	// 栈2删除元素
	// 如果栈2位空，那么将栈1中的元素挨个取出来，然后放入到栈2中去, 也就相当于反转栈了
	if c.stack2.Len() == 0 {
		for c.stack1.Len() > 0 {
			c.stack2.PushBack(c.stack1.Remove(c.stack1.Back()))
		}
	}
	// 如果栈2中有元素那么就返回元素，没有元素就返回-1
	if c.stack2.Len() > 0 {
		ele := c.stack2.Back()
		c.stack2.Remove(ele)
		return ele.Value.(int)
	}
	return -1
}
