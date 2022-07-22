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
func postorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	stack := list.List{}
	var pre *TreeNode
	for root != nil || stack.Len() > 0 {
		for root != nil {
			stack.PushBack(root)
			root = root.Left
		}
		node := stack.Remove(stack.Back()).(*TreeNode)
		// 如果该元素的右边没有或者已经访问过
		if node.Right == nil || node.Right == pre {
			// 访问中间的节点
			ret = append(ret, node.Val)
			// 且记录为访问过了
			pre = node
		}else{
			// 该节点入栈
			stack.PushBack(node)
			// 先访问右边
			root = node.Right
		}
	}
	return ret
}













