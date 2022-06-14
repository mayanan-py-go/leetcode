package main

import "fmt"

func main() {
	l1 := ListNode{Val: 1, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}
	l2 := ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 7}}}}
	l3 := mergeTwoLists(&l1, &l2)
	fmt.Println(l3.Val, l3.Next.Val, l3.Next.Next.Val, l3.Next.Next.Next.Val, l3.Next.Next.Next.Next.Val, l3.Next.Next.Next.Next.Next.Val, l3.Next.Next.Next.Next.Next.Next.Val)
}

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 判断边界条件1
	if list1 == nil{
		return list2
	}
	if list2 == nil{
		return list1
	}

	// 新建链表头
	dummy := &ListNode{0, nil}
	temp := dummy

	// 两个链表都不为nil时的处理
	for list1 != nil && list2 != nil{
		if list1.Val < list2.Val{
			temp.Next = list1
			list1 = list1.Next  // list1节点发生变化
		}else{
			temp.Next = list2
			list2 = list2.Next  // list2节点发生变化
		}
		temp = temp.Next  // 临时节点的变化
	}

	// 当list1链表比较长时，将list1剩下的接在总链表尾部
	if list1 != nil{
		temp.Next = list1
	}

	// 当list2链表比较长时，将list2剩下的接在总链表尾部
	if list2 != nil{
		temp.Next = list2
	}

	// 返回总链表
	return dummy.Next
}







