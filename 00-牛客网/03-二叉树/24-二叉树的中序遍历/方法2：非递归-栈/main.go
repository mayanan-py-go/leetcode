package main

import "container/list"

/*

*/

func main() {

}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func inorderTraversal(root *TreeNode) []int {
	var ret = make([]int, 0)
	if root == nil {
		return ret
	}
	var stack = list.List{}
	for root != nil || stack.Len() > 0 {
		for root != nil {
			stack.PushBack(root)
			root = root.Left
		}
		node := stack.Remove(stack.Back()).(*TreeNode)
		ret = append(ret, node.Val)
		root = node.Right
	}
	return ret
}












