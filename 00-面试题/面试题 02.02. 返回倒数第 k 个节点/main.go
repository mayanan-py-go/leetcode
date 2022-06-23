package main

import "fmt"

/*
实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。
https://leetcode.cn/problems/kth-node-from-end-of-list-lcci/
*/

func main() {
	l := &ListNode{Val: 11, Next: &ListNode{Val: 22, Next: &ListNode{Val: 45}}}
	fmt.Println(kthToLast(l, 2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func kthToLast(head *ListNode, k int) int {
	// 双指针法获取链表倒数第k个节点
	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow.Val
}
