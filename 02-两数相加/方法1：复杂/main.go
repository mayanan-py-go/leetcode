package main

import "fmt"

/*

 */

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	fmt.Println(addTwoNumbers(l1, l2).Next)
}
type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var val int
	var ret = new(ListNode)
	var cur = ret
	for {
		if l1 != nil && l2 != nil {
			cur.Next = new(ListNode)
			cur = cur.Next
			tmp := l1.Val + l2.Val + val
			cur.Val = tmp%10
			val = tmp / 10
			l1 = l1.Next
			l2 = l2.Next
		}else if l1 != nil {
			cur.Next = new(ListNode)
			cur = cur.Next
			tmp := l1.Val + val
			cur.Val = tmp % 10
			val = tmp / 10
			l1 = l1.Next
		}else if l2 != nil {
			cur.Next = new(ListNode)
			cur = cur.Next
			tmp := l2.Val + val
			cur.Val = tmp % 10
			val = tmp / 10
			l2 = l2.Next
		} else if val > 0 {
			cur.Next = new(ListNode)
			cur = cur.Next
			cur.Val = val
			break
		}else {
			break
		}
	}
	return ret.Next
}











