package main

import "fmt"

/*
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
https://leetcode.cn/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/

示例：
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
限制：0 <= 链表长度 <= 1000
 */

func main() {
	l1 := &ListNode{11, &ListNode{18, &ListNode{22, nil}}}
	l2 := &ListNode{10, &ListNode{17, &ListNode{23, nil}}}
	fmt.Println(mergeTwoLists(l1, l2).Next.Next.Next.Next.Next)
}
type ListNode struct {
	Val int
	Next *ListNode
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	/*
	时间复杂度：O(m+n)
	空间复杂度：O(m+n)
	 */
	dummy := new(ListNode)
	cur := dummy

	for l1 != nil && l2 != nil {
		cur.Next = new(ListNode)
		cur = cur.Next
		if l1.Val < l2.Val{
			cur.Val = l1.Val
			l1 = l1.Next
		}else {
			cur.Val = l2.Val
			l2 = l2.Next
		}
	}
	if l1 != nil {
		cur.Next = l1
	}else if l2 != nil{
		cur.Next = l2
	}
	return dummy.Next
}















