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
	root.Left, root.Right = n1, n2

	fmt.Println(levelOrder(root))
}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func levelOrder(root *TreeNode) []int {
	ret := make([]int, 0)
	queue := list.List{}
	queue.PushBack(root)
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		ret = append(ret, node.Val)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return ret
}














