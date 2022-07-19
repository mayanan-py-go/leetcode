package main

/*
给定一个链表，请判断该链表是否为回文结构。
回文是指该字符串正序逆序完全一致。
数据范围： 链表节点数 0 \le n \le 10^50≤n≤10
5
 ，链表中每个节点的值满足 |val| \le 10^7∣val∣≤10
7
*/

func main() {

}
type ListNode struct {
	Val int
	Next *ListNode
}
func isPail(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	var arr = make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
}











