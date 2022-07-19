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
	m1 := make(map[*ListNode]bool)
	for head != nil {
		if _, ok := m1[head]; ok {
			return true
		}else {
			m1[head] = true
		}
		head = head.Next
	}
	return false
}










