package main

import "fmt"

/*
输入两个链表，找出它们的第一个公共节点。
https://leetcode.cn/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/

如果两个链表没有交点，返回 null.
在返回结果后，两个链表仍须保持原有的结构。
可假定整个链表结构中没有循环。
程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。
 */

func main() {
	intersectionNode := &ListNode{55, nil}
	l1 := &ListNode{11, &ListNode{18, &ListNode{22, intersectionNode}}}
	l2 := &ListNode{10, &ListNode{17, intersectionNode}}
	fmt.Println(getIntersectionNode(l1, l2))
}
type ListNode struct {
	Val int
	Next *ListNode
}
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	/*
	时间复杂度：O(n+m)
	空间复杂度：O(1)
	 */
	h1, h2 := headA, headB
	for h1 != nil || h2 != nil {
		if h1 == h2 {
			return h1
		}
		if h1 == nil {
			h1 = headB
		}else {
			h1 = h1.Next
		}
		if h2 == nil {
			h2 = headA
		}else {
			h2 = h2.Next
		}
	}
	return nil
}















