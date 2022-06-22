package main

import (
	"fmt"
)

/*
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
https://leetcode.cn/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
*/

func main() {
	head := &ListNode{Val: 11}
	head.Next = &ListNode{Val: 22}
	head.Next.Next = &ListNode{Val: 15}

	fmt.Println(reversePrint(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var ret = make([]int, 0)
	if head == nil {
		return ret
	}
	for ; head != nil; head = head.Next {
		var tmp = make([]int, len(ret)+1)
		tmp[0] = head.Val
		copy(tmp[1:], ret)
		ret = tmp
	}
	return ret
}
