package main

/*
给定两个 非空链表 l1和 l2 来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

可以假设除了数字 0 之外，这两个数字都不会以零开头。

*/

func main() {

}
type ListNode struct {
	Val int
	Next *ListNode
}
func addInList(head1 *ListNode, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	n1 := reverseList(head1)
	n2 := reverseList(head2)

	var dummyNode = new(ListNode)
	cur := dummyNode
	var value int
	for n1 != nil || n2 != nil || value > 0 {
		cur.Next = new(ListNode)
		cur = cur.Next
		if n1 != nil {
			value += n1.Val
			n1 = n1.Next
		}
		if n2 != nil {
			value += n2.Val
			n2 = n2.Next
		}
		cur.Val = value%10
		value /= 10
	}
	cur = dummyNode.Next
	return reverseList(cur)
}
func reverseList(node *ListNode) *ListNode {
	var prev *ListNode
	for node != nil {
		tmp := node.Next
		node.Next = prev
		prev = node
		node = tmp
	}
	return prev
}












