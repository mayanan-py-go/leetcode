package main

/*

*/

func main() {

}
type ListNode struct {
	Val int
	Next *ListNode
}
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for slow != nil && fast != nil {
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		}else {
			return false
		}
		if slow == fast {
			return true
		}
	}
	return false
}










