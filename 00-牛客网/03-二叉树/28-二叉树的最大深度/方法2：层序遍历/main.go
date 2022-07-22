package main

import (
	"container/list"
	"fmt"
)

func main() {
	root := &TreeNode{Val: 1}
	n1 := &TreeNode{Val: 2}
	n2 := &TreeNode{Val: 3}
	n3 := &TreeNode{Val: 4}
	n4 := &TreeNode{Val: 5}
	root.Left, root.Right = n1, n2
	n1.Left, n2.Right = n3, n4
	fmt.Println(maxDepth(root))
}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func maxDepth(root *TreeNode) int {
	max := 0
	if root == nil {
		return max
	}
	queue := list.List{}
	queue.PushBack(root)
	for queue.Len() > 0 {
		qLen := queue.Len()
		for i := 0; i < qLen; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		max++
	}
	return max
}











