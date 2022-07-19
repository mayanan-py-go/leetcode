package main

import "fmt"

/*
给定一个单链表，请设定一个函数，将链表的奇数位节点和偶数位节点分别放在一起，重排后输出。
注意是节点的编号而非节点的数值。

数据范围：节点数量满足 0 \le n \le 10^50≤n≤10
5
 ，节点中的值都满足 0 \le val \le 10000≤val≤1000
要求：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
*/

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}
	fmt.Println(oddEvenList(head))
}
type ListNode struct {
	Val int
	Next *ListNode
}
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 经典双指针法
	first, second, even := head, head.Next, head.Next
	for second != nil && second.Next != nil {
		first.Next = second.Next
		first = first.Next

		second.Next = first.Next
		second = second.Next
	}
	first.Next = even
	return head
}












