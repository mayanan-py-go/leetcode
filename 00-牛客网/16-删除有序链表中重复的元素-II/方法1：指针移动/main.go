package main

/*

*/

func main() {
	//head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}
}
type ListNode struct {
	Val int
	Next *ListNode
}
func deleteDuplicates(head *ListNode) *ListNode {
	/*
	时间复杂度：O(n)
	空间复杂度：O(1)
	 */
	// 创建虚拟头节点
	if head == nil {
		return nil
	}
	dummyNode := new(ListNode)
	dummyNode.Next = head
	cur := dummyNode
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			tmp := cur.Next.Val
			// 跳过所有相同的节点
			for cur.Next != nil && cur.Next.Val == tmp {
				cur.Next = cur.Next.Next
			}
		}else{
			cur = cur.Next
		}
	}
	return dummyNode.Next
}
















