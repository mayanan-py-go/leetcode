package main

import "sort"

/*
给定一个节点数为n的无序单链表，对其按升序排序。
数据范围：0 < n \le 1000000<n≤100000
要求：时间复杂度 O(nlogn)O(nlogn)
*/

func main() {

}
type ListNode struct {
	Val int
	Next *ListNode
}
func sortInList(head *ListNode) *ListNode {
	var l1 = make([]int, 0)
	for head != nil {
		l1 = append(l1, head.Val)
		head = head.Next
	}
	sort.Ints(l1)
	var dummyNode = new(ListNode)
	cur := dummyNode
	for _, v := range l1 {
		cur.Next = new(ListNode)
		cur = cur.Next
		cur.Val = v
	}
	return dummyNode.Next
}











