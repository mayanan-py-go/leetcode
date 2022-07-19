package main

/*
给一个长度为n链表，若其中包含环，请找出该链表的环的入口结点，否则，返回null。

数据范围： n\le10000n≤10000，1<=结点值<=100001<=结点值<=10000
要求：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)
*/

func main() {

}
type ListNode struct {
	Val int
	Next *ListNode
}
func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	if pHead == nil {
		return pHead
	}
	m1 := make(map[*ListNode]bool)
	for pHead != nil {
		if m1[pHead] {
			return pHead
		}else {
			m1[pHead] = true
		}
		pHead = pHead.Next
	}
	return nil
}










