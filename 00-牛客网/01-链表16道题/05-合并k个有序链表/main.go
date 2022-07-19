package main

import "sort"

/*
合并 k 个升序的链表并将结果作为一个升序的链表返回其头节点。
数据范围：节点总数满足 0 \le n \le 10^50≤n≤10
5
 ，链表个数满足 1 \le k \le 10^5 \1≤k≤10
5
   ，每个链表的长度满足 1 \le len \le 200 \1≤len≤200  ，每个节点的值满足 |val| <= 1000∣val∣<=1000
要求：时间复杂度 O(nlogk)O(nlogk)
*/

func main() {

}
type ListNode struct {
	Val int
	Next *ListNode
}
func mergeKLists(lists []*ListNode) *ListNode {
	var ret []int
	for i := 0; i < len(lists); i++ {
		node := lists[i]
		for node != nil {
			ret = append(ret, node.Val)
			node = node.Next
		}
	}
	sort.Ints(ret)
	dummyNode := new(ListNode)
	cur := dummyNode
	for i := 0; i < len(ret); i++ {
		cur.Next = &ListNode{Val: ret[i]}
		cur = cur.Next
	}
	return dummyNode.Next
}











