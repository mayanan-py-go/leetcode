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
	// 获取链表长度
	var length int
	tmpHead := head
	for ; tmpHead != nil; tmpHead = tmpHead.Next {
		length++
	}

	// 初始化切片，从后面挨个赋值
	s1 := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		s1[i] = head.Val
		head = head.Next
	}
	return s1
}
