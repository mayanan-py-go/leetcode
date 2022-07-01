package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val: 5, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}
	l2 := &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}}
	l3 := addTwoNumbers(l1, l2)
	fmt.Println(*l3, *l3.Next, *l3.Next.Next)
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	// 定义新的空链表指针，用于输出
	dummy := new(ListNode)  // dummy用于返回的链表
	curr := dummy  // curr用于游走的指针给dummy节点赋值
	carry := 0

	// 当l1 或l2 有一个不为nil时，或者carry>0
	for l1 != nil || l2 != nil || carry > 0 {
		curr.Next = new(ListNode)
		curr = curr.Next

		// 将l1和l2两个同位数相加
		if l1 != nil{
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil{
			carry += l2.Val
			l2 = l2.Next
		}

		// 若同位数相加 >= 10, 则curr.Val = 余数
		// 若同尾数相加 <10, 则curr.Val = carry
		curr.Val = carry % 10

		// 若同位数相加 >=10, 则carry=1
		// 若同位数相加 <10, 则carry=0
		carry /= 10
	}

	return dummy.Next
}



