package main

import "fmt"

/*
输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。
例如，一个链表有 6 个节点，从头节点开始，它们的值依次是 1、2、3、4、5、6。这个链表的倒数第 3 个节点是值为 4 的节点。
https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/

示例：
给定一个链表: 1->2->3->4->5, 和 k = 2.
返回链表 4->5.
 */

func main() {
	l1 := &ListNode{11, &ListNode{22, &ListNode{33, &ListNode{44, nil}}}}
	fmt.Println(getKthFromEnd(l1, 2).Next)
}
type ListNode struct {
	Val int
	Next *ListNode
}
func getKthFromEnd(head *ListNode, k int) *ListNode {  // [1, 2, 3, 4, 5]
	slow, fast := head, head
	for i := 0; i < k; i++ {
		if fast == nil {
			return head
		}
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}











