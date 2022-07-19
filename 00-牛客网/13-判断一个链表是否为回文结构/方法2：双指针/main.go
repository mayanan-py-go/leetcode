package main

import "fmt"

/*
给定一个链表，请判断该链表是否为回文结构。
回文是指该字符串正序逆序完全一致。
数据范围： 链表节点数 0 \le n \le 10^50≤n≤10
5
 ，链表中每个节点的值满足 |val| \le 10^7∣val∣≤10
7
*/

func main() {
	// -401261,-449050,-456674,-456674,-449050,-401261
	//l1 := &ListNode{-401261, &ListNode{-449050, &ListNode{-456674,
	//	&ListNode{-456674, &ListNode{-449050, &ListNode{-401261, nil},
	//}}}}}
	l1 := &ListNode{-129, &ListNode{-129, nil}}
	fmt.Println(isPail(l1))
}
type ListNode struct {
	Val int
	Next *ListNode
}
func isPail(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	var reverse func(*ListNode) *ListNode
	reverse = func(slow *ListNode) *ListNode {
		// 千万要注意：这里千万不能初始化链表比如：prev := new(ListNode)
		var prev *ListNode
		cur := slow
		for cur != nil {
			tmp := cur.Next
			cur.Next = prev
			prev = cur
			cur = tmp
		}
		return prev
	}

	// 快慢指针寻找中点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 反转链表
	slow = reverse(slow)
	fast = head
	for slow != nil {
		if slow.Val != fast.Val {
			return false
		}
		slow = slow.Next
		fast = fast.Next
	}
	return true
}












