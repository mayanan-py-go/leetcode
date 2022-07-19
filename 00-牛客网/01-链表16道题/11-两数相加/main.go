package main

/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，
并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
*/

func main() {

}
type ListNode struct {
	Val int
	Next *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummyNode = new(ListNode)
	cur := dummyNode
	var value int
	for l1 != nil || l2 != nil || value > 0 {
		cur.Next = new(ListNode)
		cur = cur.Next
		if l1 != nil {
			value += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			value += l2.Val
			l2 = l2.Next
		}
		cur.Val = value % 10
		value /= 10
	}
	return dummyNode.Next
}












