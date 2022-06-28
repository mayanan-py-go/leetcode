package main

import "fmt"

/*
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。
题目数据 保证 整个链式结构中不存在环。
注意，函数返回结果后，链表必须 保持其原始结构 。
https://leetcode.cn/problems/intersection-of-two-linked-lists-lcci/
*/

func main() {
	ll := &ListNode{Val: 33}
	l1 := &ListNode{Val: 11, Next: &ListNode{Val: 22, Next: ll}}
	l2 := &ListNode{Val: 11, Next: ll}
	l2.Next.Next = &ListNode{Val: 44}
	fmt.Println(getIntersectionNode(l1, l2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var m1 = make(map[*ListNode]bool)
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		m1[tmp] = true
	}

	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if m1[tmp] {
			return tmp
		}
	}
	return nil
}










