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
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	var dummyNode = new(ListNode)
	cur := dummyNode
	fast := head
	for fast != nil {
		// 此处是核心，不要直接等于fast, 而要根据fast的值进行重建，因为如果直接等于fast就会把fast链表后续所有的链都引用过来了，
		// 而这样的话只是增加了一个值
		cur.Next = &ListNode{Val: fast.Val}
		cur = cur.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		}else{
			break
		}
	}
	slow := head.Next
	for slow != nil {
		cur.Next = &ListNode{Val: slow.Val}
		cur = cur.Next
		if slow.Next != nil {
			slow = slow.Next.Next
		}else {
			break
		}
	}
	return dummyNode.Next
}












