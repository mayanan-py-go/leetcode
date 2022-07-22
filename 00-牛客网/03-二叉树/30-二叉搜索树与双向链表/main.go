package main

import "fmt"

/*
输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表。如下图所示
 */

func main() {
	root := &TreeNode{Val: 1}
	n1 := &TreeNode{Val: 2}
	n2 := &TreeNode{Val: 3}
	n3 := &TreeNode{Val: 4}
	n4 := &TreeNode{Val: 5}
	root.Left, root.Right = n1, n2
	n1.Left, n1.Right = n3, n4
	fmt.Println(Convert(root))
}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func Convert(pRootOfTree *TreeNode) *TreeNode {
	var pre, head *TreeNode
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre == nil {
			pre, head = node, node
		}else{
			pre.Right = node
			node.Left = pre
			pre = node
		}
		dfs(node.Right)
	}
	dfs(pRootOfTree)
	return head
}















