package main

/*
输入两个无环的单向链表，找出它们的第一个公共结点，如果没有公共节点则返回空。
（注意因为传入数据是链表，所以错误测试数据的提示是用其他方式显示的，保证传入数据是正确的）
*/

func main() {

}
type ListNode struct {
	Val int
	Next *ListNode
}
func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	if pHead1 == nil || pHead2 == nil {
		return nil
	}
	n1, n2 := pHead1, pHead2
	for n1 != n2 {
		if n1 == nil {
			n1 = pHead2
		}else {
			n1 = n1.Next
		}
		if n2 == nil {
			n2 = pHead1
		}else {
			n2 = n2.Next
		}
	}
	return n1
}












