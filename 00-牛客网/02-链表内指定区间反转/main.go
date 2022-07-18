package main

import "fmt"

/*
将一个节点数为 size 链表 m 位置到 n 位置之间的区间反转，要求时间复杂度 O(n)O(n)，空间复杂度 O(1)O(1)。
例如：
给出的链表为 1\to 2 \to 3 \to 4 \to 5 \to NULL1→2→3→4→5→NULL, m=2,n=4m=2,n=4,
返回 1\to 4\to 3\to 2\to 5\to NULL1→4→3→2→5→NULL.

数据范围： 链表长度 0 < size \le 10000<size≤1000，0 < m \le n \le size0<m≤n≤size，链表中每个节点的值满足 |val| \le 1000∣val∣≤1000
要求：时间复杂度 O(n)O(n) ，空间复杂度 O(n)O(n)
进阶：时间复杂度 O(n)O(n)，空间复杂度 O(1)O(1)
*/

func main() {
	node1 := &ListNode{Val: 11}
	node2 := &ListNode{Val: 12}
	node1.Next = node2

	fmt.Println(reverseBetween(node1, 1, 1))
}
type ListNode struct {
	Val int
	Next *ListNode
}
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// 创建虚拟头节点
	dummyNode := new(ListNode)
	dummyNode.Next = head

	// 定位第m个节点和m的前一个节点
	prev, cur := dummyNode, head
	for i := 0; i < m-1; i++ {
		prev = cur
		cur = cur.Next
	}

	// 从m反转到n
	for i := m; i < n; i++ {
		temp := cur.Next
		cur.Next = temp.Next
		temp.Next = prev.Next
		prev.Next = temp
	}

	return dummyNode.Next

}














