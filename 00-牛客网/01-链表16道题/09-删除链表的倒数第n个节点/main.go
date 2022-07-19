package main

/*
给定一个链表，删除链表的倒数第 n 个节点并返回链表的头指针
例如，
给出的链表为: 1\to 2\to 3\to 4\to 51→2→3→4→5, n= 2n=2.
删除了链表的倒数第 nn 个节点之后,链表变为1\to 2\to 3\to 51→2→3→5.

数据范围： 链表长度 0\le n \le 10000≤n≤1000，链表中任意节点的值满足 0 \le val \le 1000≤val≤100
要求：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)
备注：
题目保证 nn 一定是有效的
*/

func main() {

}
type ListNode struct {
	Val int
	Next *ListNode
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 为了防止删除的倒数第N个节点正好是正数第一个节点，所以我们创建一个虚拟头节点
	dummyNode := new(ListNode)
	dummyNode.Next = head

	// 先获取链表到数第n-1个节点，然后再删除倒数第n个节点
	fast, slow := dummyNode, dummyNode
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummyNode.Next
}












