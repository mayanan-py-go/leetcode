package main

import (
	"container/list"
	"fmt"
)

/*

*/

func main() {
	root := &TreeNode{Val: 1}
	n1 := &TreeNode{Val: 2}
	n2 := &TreeNode{Val: 3}
	n3 := &TreeNode{Val: 4}
	n4 := &TreeNode{Val: 5}
	root.Left, root.Right = n1, n2
	n1.Left, n2.Right = n3, n4
	fmt.Println(Print(root))
}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func Print(pRoot *TreeNode) [][]int {
	ret := make([][]int, 0)
	if pRoot == nil {
		return ret
	}
	k := 0
	queue := list.List{}
	queue.PushBack(pRoot)
	for queue.Len() > 0 {
		qLen := queue.Len()
		ret = append(ret, make([]int, qLen))
		for i := 0; i < qLen; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			ret[k][i] = node.Val
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		// 偶数不反转，奇数反转
		if k%2 == 1{
			reverse(ret[k])
		}
		k++
	}
	return ret
}
func reverse(li []int) {
	for i, j := 0, len(li)-1; i < j; i, j = i+1, j-1 {
		li[j], li[i] = li[i], li[j]
	}
}









