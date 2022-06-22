package main

import "fmt"

/*
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
https://leetcode.cn/problems/fan-zhuan-lian-biao-lcof/
*/

func main() {
	head := &ListNode{Val: 11}
	head.Next = &ListNode{Val: 15}
	head.Next.Next = &ListNode{Val: 18}

	fmt.Println(reverseList(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var next *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = next
		next = head
		head = tmp
	}
	return next
}
