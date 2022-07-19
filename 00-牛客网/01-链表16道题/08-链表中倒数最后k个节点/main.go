package main

/*
输入一个长度为 n 的链表，设链表中的元素的值为 ai ，返回该链表中倒数第k个节点。
如果该链表长度小于k，请返回一个长度为 0 的链表。
*/

func main() {

}
type ListNode struct {
	Val int
	Next *ListNode
}
func FindKthToTail(pHead *ListNode, k int) *ListNode {
	if pHead == nil || k == 0 {
		return nil
	}
	fast, slow := pHead, pHead
	for i := 0; i < k-1; i++ {
		fast = fast.Next
		if fast == nil {
			return nil
		}
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}












