package main

/*

*/

func main() {
	//head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}
}
type ListNode struct {
	Val int
	Next *ListNode
}
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		}else{
			cur = cur.Next
		}
	}
	return head
}













