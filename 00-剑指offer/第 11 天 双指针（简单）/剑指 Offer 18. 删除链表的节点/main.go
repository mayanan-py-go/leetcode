package main

import "fmt"

/*
给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
返回删除后的链表的头节点。
注意：此题对比原题有改动
https://leetcode.cn/problems/shan-chu-lian-biao-de-jie-dian-lcof/
示例：
输入: head = [4,5,1,9], val = 5
输出: [4,1,9]
解释: 给定你链表中值为5的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
 */

func main() {
	l1 := &ListNode{11, &ListNode{22, &ListNode{33, &ListNode{44, nil}}}}
	fmt.Println(deleteNode(l1, 22).Next.Next.Next)
}
type ListNode struct {
	Val int
	Next *ListNode
}
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	prev := head
	cur := head.Next
	for cur != nil {
		if cur.Val != val {
			cur = cur.Next
			prev = prev.Next
		} else {
			prev.Next = cur.Next
			break
		}
	}
	return head
}










