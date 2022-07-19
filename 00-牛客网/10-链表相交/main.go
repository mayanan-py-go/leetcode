package main

/*

*/

func main() {

}
type ListNode struct {
	Val int
	Next *ListNode
}
func getIntersectionNode(headA *ListNode, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a1, a2 := headA, headB
	for a1 != a2 {
		if a1 == nil {
			a1 = headB
		}else {
			a1 = a1.Next
		}
		if a2 == nil {
			a2 = headA
		}else{
			a2 = a2.Next
		}
	}
	return a1
}













