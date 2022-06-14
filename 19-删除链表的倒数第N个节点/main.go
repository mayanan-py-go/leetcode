package main

import "fmt"

func main() {
	// 删除链表的倒数第N个节点，返回链表的头节点
	l1 := ListNode{Val: 11, Next: &ListNode{Val: 22, Next: &ListNode{Val: 33, Next: &ListNode{Val: 44}}}}
	l2 := removeNthFromEnd(&l1, 5)
	fmt.Println(l2.Val, l2.Next.Val, l2.Next.Next.Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 定义一个快节点，一个慢节点
	slow, fast := head, head

	// 将快慢节点拉开，距离相差n
	for i := 0; i < n; i++{
		fast = fast.Next
		// 如果n大于等于链表的长度，直接将头节点移除，返回第二个节点即可
		if fast == nil{
			head = head.Next
			return head
		}
	}

	// 快节点和慢节点同步往后走，直到快节点走到最后
	for fast.Next != nil{
		fast = fast.Next
		slow = slow.Next
	}

	// 需要移除的就是慢节点的下一个节点
	// 因为慢节点和快节点相差n, 快节点在最后一个, 慢节点也就在从后面数n+1的位置，慢节点-1就是需要移除的n的位置
	slow.Next = slow.Next.Next
	return head

}
