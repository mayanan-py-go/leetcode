package main

import "fmt"

/*

*/

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{2, nil}}}
	fmt.Println(deleteDuplicates(head))
}
type ListNode struct {
	Val int
	Next *ListNode
}
func deleteDuplicates(head *ListNode) *ListNode {
	/*
	时间复杂度：O(n)
	空间复杂度：O(n)
	 */
	if head == nil {
		return nil
	}
	m1 := make(map[int]int)
	cur := head
	for cur != nil {
		if _, ok := m1[cur.Val]; ok {
			m1[cur.Val] = m1[cur.Val]+1
		}else{
			m1[cur.Val] = 1
		}
		cur = cur.Next
	}

	dummyNode := new(ListNode)
	dummyNode.Next = head
	cur = dummyNode
	for cur.Next != nil {
		if m1[cur.Next.Val] > 1 {
			cur.Next = cur.Next.Next
		}else{
			cur = cur.Next
		}
	}
	return dummyNode.Next
}
















