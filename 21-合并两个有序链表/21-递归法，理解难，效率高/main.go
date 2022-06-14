package main

import "fmt"

func main() {
	l1 := ListNode{Val: 1, Next: &ListNode{Val: 5, Next: nil}}
	l2 := ListNode{Val: 3}
	l3 := mergeTwoLists(&l1, &l2)
	fmt.Println(l3.Val, l3.Next.Val, l3.Next.Next.Val)
}

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil{
		return list2
	}
	if list2 == nil{
		return list1
	}

	var res *ListNode
	if list1.Val < list2.Val{
		res = list1
		res.Next = mergeTwoLists(list1.Next, list2)
	}else{
		res = list2
		res.Next = mergeTwoLists(list1, list2.Next)
	}

	return res
}







