package main

import (
	"container/list"
)

/*

*/

func main() {

}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func preorderTraversal(root *TreeNode) []int {
	var ret = make([]int, 0)
	if root == nil {
		return ret
	}

	var stack = list.List{}
	stack.PushBack(root)
	for stack.Len() > 0 {  // 栈中有元素
		node := stack.Remove(stack.Back()).(*TreeNode)
		ret = append(ret, node.Val)
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return ret
}














